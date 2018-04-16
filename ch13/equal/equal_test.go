package equal

import (
	//"bytes"
	"fmt"
	//"testing"
)

func ExampleEqualCycle() {
	//!+cycle
	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Equal(a, a)) // "true"
	fmt.Println(Equal(b, b)) // "true"
	fmt.Println(Equal(c, c)) // "true"
	fmt.Println(Equal(a, b)) // "false"
	fmt.Println(Equal(a, c)) // "false"
	//!-cycle

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleEqual() {

	type aaa struct {
		n     int
		value string
		num   int
	}
	type bbb struct {
		n   int
		val string
		no  int
	}

	a := aaa{n: 5, value: "test", num: 1}
	b := bbb{n: 5, val: "test", no: 1}

	fmt.Println(Equal(a, b))

	// Output:
	// true
}
