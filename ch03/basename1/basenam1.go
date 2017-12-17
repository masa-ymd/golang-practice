package main

import (
    "fmt"
)

func main() {
    s1 := "/a/b/c.go"
    s2 := "c.d.go"
    s3 := "abc"

    fmt.Printf("%s %s %s\n", basename(s1), basename(s2), basename(s3))
}

func basename(s string) string {
    // 最後の/とその前の全ての文字列を破棄する
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '/' {
            s = s[i+1:]
            break
        }
    }
    // 最後の.より前の全てを保持する
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '.' {
            s = s[:i]
            break
        }
    }

    return s
}
