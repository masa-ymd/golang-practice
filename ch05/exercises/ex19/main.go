package main

import (
	"fmt"
)

func rtn() (ret string) {
	defer func() {
		p := recover()
		if p != nil {
			ret = "recover"
		}
	}()
	panic("panic!")
}

func main() {
	fmt.Println(rtn())
}
