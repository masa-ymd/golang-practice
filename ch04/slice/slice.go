package main

import (
    "fmt"
)

func main() {
    months := [...]string{
        1: "January",
        2: "February",
        3: "March",
        4: "April",
        5: "May",
        6: "June",
        7: "July",
        8: "August",
        9: "September",
        10: "October",
        11: "November",
        12: "December",
    }

    s1 := months[4:7]

    fmt.Printf("%v, %[1]T, %d, %d\n", s1, len(s1), cap(s1))

    s2 := s1[:5]

    fmt.Printf("%v, %[1]T, %d, %d\n", s2, len(s2), cap(s2))

    // スライスのゼロ値はnil, 長さが0でもnilでない場合もある
    var s3 []int
    fmt.Printf("%d, %t\n", len(s3), s3==nil)
    s4 := []int{}
    fmt.Printf("%d %t\n", len(s4), s4==nil)

    // makeで指定された型、長さ、容量のスライスを作成可能
    s5 := make([]int, 3, 5)
    fmt.Printf("%d %d\n", len(s5), cap(s5))
}
