package main

import "fmt"

func main() {
    var f float64 = 123456789

    // %e-指数 %f-指数なし %g-適切に表示
    fmt.Printf("%e, %[1]f %[1]g\n", f)
}
