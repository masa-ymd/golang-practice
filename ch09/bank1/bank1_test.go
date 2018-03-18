package bank1_test

import (
	"fmt"
	"testing"

	"github.com/masa-ymd/golang-practice/ch09/bank1"
)

func TestBank1(t *testing.T) {
	done := make(chan struct{})

	// Alice 書き込み+読み出し
	go func() {
		bank1.Deposit(200)
		fmt.Println("=", bank1.Balance())
		done <- struct{}{}
	}()

	// Bob 書き込み
	go func() {
		bank1.Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := bank1.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
