package delspace

import "fmt"

func ExampleDelspace() {
	buf := []byte("   de")
	buf2 := delspace(buf)
	fmt.Printf("result: %s\n", string(buf2))
	// Output:
	// result:  de
}
