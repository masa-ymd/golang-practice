package dupdel

import "fmt"

func ExampleDupdel() {
	s := []string{"a", "b", "c", "c", "d", "d"}
	r := dupdel(s)
	fmt.Println(r)
	// Output:
	// [a b c d]
}
