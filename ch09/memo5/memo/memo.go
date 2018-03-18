package memo

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // resが読み出しOKの時closeされる
}

type request struct {
	key      string
	response chan<- result
}

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response} // チャネルにURLと値を格納するためのチャネルを送信
	res := <-response                       // 応答を待つ
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) } // 送信をやめる

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry) // モニターゴルーチンでマップ変数を閉じ込める
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// 未取得の場合
			e = &entry{ready: make(chan struct{})} // チャネルをオープンし他ゴルーチンを待たせる
			cache[req.key] = e
			go e.call(f, req.key) // 他処理を止めないため、新たなゴルーチンで処理を開始
		}
		go e.deliver(req.response) // 他処理を止めないため、新たなゴルーチンを生成
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key) // 値を取得
	close(e.ready)                  // 結果が取得できたらクローズし、ブロードキャストする
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready // 結果が用意できるのをまつ
	response <- e.res
}
