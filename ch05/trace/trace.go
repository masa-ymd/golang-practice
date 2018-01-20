package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	// 以下のようにdeferは入った処理と出た処理を一対にするために使うこともできる
	defer trace("bigSlowOperation")() // 最後の()は必要
	// ...大量の処理をsleepで模倣
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

// 名前付き結果を定義し、defer実行時に結果を操作することも可能
func double(x int) (result int) {
	defer func() {
		result += x
		fmt.Printf("double(%d) = %d", x, result)
	}()
	return x + x
}

func main() {
	bigSlowOperation()
	_ = double(2)
}
