package rev

import (
	"fmt"
)

func ExampleRev() {
	s := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&s)
	fmt.Printf("%v\n", s)
	// Output:
	// [5 4 3 2 1 0]
}
