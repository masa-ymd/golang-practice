package main

import (
    "fmt"
    "strconv"
)

func main() {
    x := 123
    // 変換の方法の一つSprintf
    y := fmt.Sprintf("%d", x)
    // もうひとつの選択肢 strconv.Itoa
    fmt.Println(y, strconv.Itoa(x))

    // 異なる基数に変換
    fmt.Println(strconv.FormatInt(int64(x), 2))

    // 文字列を整数に変換する方法１
    fmt.Println(strconv.Atoi("123"))
    // 方法２
    fmt.Println(strconv.ParseInt("123", 10, 64)) // 基数は10, 64bitまで
}
