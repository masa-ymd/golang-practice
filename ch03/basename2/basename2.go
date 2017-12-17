package main

import (
    "fmt"
    "strings"
)

func main() {
    s1 := "/a/b/c.go"
    s2 := "c.d.go"
    s3 := "abc"

    fmt.Printf("%s %s %s\n", basename(s1), basename(s2), basename(s3))
}

func basename(s string) string {
    slash := strings.LastIndex(s, "/") // '/'が見つからないと-1
    s = s[slash+1:]
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        s = s[:dot]
    }
    return s
}
