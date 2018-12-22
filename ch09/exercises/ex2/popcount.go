package popcount

import (
	"sync"
)

var pc [256]byte
var loadTableOnce sync.Once

// init関数は複数宣言可能
// プログラムが開始した時点で、
// 宣言された順番に実行される
func Table() {
	loadTableOnce.Do(func() {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
	})
}

func Popcount(x uint64) int {
	Table()
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}
