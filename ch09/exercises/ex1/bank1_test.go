package ex1

import (
	"fmt"
	"testing"
)

func TestBank1(t *testing.T) {
	done := make(chan struct{})

	// Alice 書き込み+読み出し
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob 書き込み
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBank2(t *testing.T) {
	done := make(chan struct{})

	// Alice 書き込み+読み出し
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob 書き込み
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	Withdraw(100)

	if got, want := Balance(), 500; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func ExampleBank() {
	Withdraw(600)
	// Output:
	// error
}
