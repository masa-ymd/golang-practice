package rotate

import (
	"fmt"
)

func ExampleRoatate() {
	s := []int{1, 2, 3, 4}
	rotate(s, 1)
	fmt.Printf("%v\n", s)
	// Output:
	// [2 3 4 1]
}
