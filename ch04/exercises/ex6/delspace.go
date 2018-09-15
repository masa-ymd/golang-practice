package delspace

import (
	"unicode"
)

func delspace(buf []byte) []byte {
	var d int
	var bfSpace bool
	for _, b := range buf {
		if unicode.IsSpace(rune(b)) {
			if bfSpace {
				d++
			}
			bfSpace = true
		} else {
			bfSpace = false
		}
	}
	return buf[d:]
}
