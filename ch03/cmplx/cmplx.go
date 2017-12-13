package main

import "fmt"

func main() {
    var x complex128 = complex(1, 2)

    fmt.Println(x, real(x), imag(x))
}
