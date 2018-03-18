package memo

import (
	"sync"
)

type entry struct {
	res   result
	ready chan struct{} // resが設定されたら排他制御
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// 値が未取得の場合、値を初期化するまでロック
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		// 以下でチャネルをクローズするまで、待たせながら時間がかかる処理を行う
		e.res.value, e.res.err = memo.f(key)

		close(e.ready)
	} else {
		// すでに初期化済（値は取得中かもしれない）の場合、キャッシュのロックはすぐに開放
		memo.mu.Unlock()

		<-e.ready // チャネルのクローズ＝書き込み終了を待つ
	}
	return e.res.value, e.res.err
}
