package main

func pipeline(n int) (in chan int, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < n; i++ {
		in = out
		out = make(chan int)
		// https://qiita.com/sudix/items/67d4cad08fe88dcb9a6d
		go func(in chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}
	return first, out
}
