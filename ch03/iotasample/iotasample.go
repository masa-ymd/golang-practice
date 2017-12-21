package main

import (
	"fmt"
)

type Weekday int

const (
	Sunday Weekday = iota
	Mondey
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota // 2進数を1ずつ左にシフト
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

const (
    _ = 1 << (10 * iota) // 2進数を10ずつ左にシフト
    KiB
    MiB
    GiB
    TiB
    PiB
)

func main() {
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
	fmt.Printf("%b %b %b %b %b\n", FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
    fmt.Println(KiB, MiB, GiB, TiB, PiB)
    fmt.Printf("%b %b %b %b %b\n", KiB, MiB, GiB, TiB, PiB)
}
