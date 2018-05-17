package popcount

//import "fmt"

// 0～255の数字（インデックス）に対応する1がたっているbit数をいれる器
var pc [256]byte

// 器(pc)に0～255に対応する1がたっているbit数をいれる初期化関数
func init() {
	for i := range pc {
		// & ビット演算 byte(i&1)は末尾のbitが1のとき、1なので、0,1繰り返し
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Printf("%b, %b\n", byte(uint64(255)), byte(uint64(255)>>8))
	}
}

func Popcount(x uint64) int {
	// fmt.Printf("%v\n", pc)
	return int(
		// 8 bit -> 2^8 -> 256
		//　8bitずつ右にずらし（下8bitを切り捨て）下0～255までの数字について1がたっているbitがいくつあるかを
		// pcを利用して算出する
		// それを8回（2^8^8桁）繰り返す
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopcountFor(x uint64) int {
	var res int
	var i uint
	for i = 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}

	return res
}

func PopcountShift(x uint64) int {
	var res int
	var i uint
	for i = 0; i < 64; i++ {
		res += int(byte((x >> i) & 1))
	}
	return res
}

func PopcountClear(x uint64) int {
	var res int
	for x != 0 {
		//fmt.Printf("%b, %b\n", byte(x), byte(x-1))
		x &= x - 1
		res++

	}
	return res
}
