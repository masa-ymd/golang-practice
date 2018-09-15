package revbyte

import "fmt"

func ExampleRevbyte() {
	buf := []byte("abcde")
	revbyte(buf)
	fmt.Println(string(buf))
	// Output:
	// edcba
}
