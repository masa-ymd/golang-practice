package memo

import (
	"fmt"
)

type Func func(key string, done <-chan struct{}) (interface{}, error)

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
	done     <-chan struct{}
	response chan<- result
}

type Memo struct {
	requests, cancels chan request
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request), cancels: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, done, response} // チャネルにURLと値を格納するためのチャネルを送信
	fmt.Println("get: waiting for response")
	res := <-response // 応答を待つ
	fmt.Println("get: checking if cancelled")
	select {
	case <-done:
		fmt.Println("get: queueing cancellation request")
		memo.cancels <- request{key, done, response}
	default:
	}
	fmt.Println("get: return")
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) } // 送信をやめる

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry) // モニターゴルーチンでマップ変数を閉じ込める
Loop:
	for {
	Cancel:
		for {
			select {
			case req := <-memo.cancels:
				fmt.Println("server: deleting cancelled entry (early)")
				delete(cache, req.key)
			default:
				break Cancel
			}

		}
		select {
		case req := <-memo.cancels:
			fmt.Println("server: deleting cancelled entry")
			delete(cache, req.key)
			continue Loop
		case req := <-memo.requests:
			fmt.Println("server: request")
			e := cache[req.key]
			if e == nil {
				// 未取得の場合
				e = &entry{ready: make(chan struct{})} // チャネルをオープンし他ゴルーチンを待たせる
				cache[req.key] = e
				go e.call(f, req.key, req.done) // 他処理を止めないため、新たなゴルーチンで処理を開始
			}
			go e.deliver(req.response) // 他処理を止めないため、新たなゴルーチンを生成
		}
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	e.res.value, e.res.err = f(key, done) // 値を取得
	fmt.Println("call: returned from f")
	close(e.ready) // 結果が取得できたらクローズし、ブロードキャストする
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready // 結果が用意できるのをまつ
	response <- e.res
}
