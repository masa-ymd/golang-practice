package main

import (
	"fmt"
)

const (
	thd = 1000
	KB  = thd
	MB  = KB * thd
	GB  = MB * thd
	TB  = GB * thd
	PB  = TB * thd
	EB  = PB * thd
	ZB  = EB * thd
	YB  = ZB * thd
)

func main() {
	fmt.Printf("%v\n", PB)
}
