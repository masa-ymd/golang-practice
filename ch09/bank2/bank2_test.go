package bank2_test

import (
	"sync"
	"testing"

	"github.com/masa-ymd/golang-practice/ch09/bank2"
)

func TestBank(t *testing.T) {
	// Deposit を1000個同時起動
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank2.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank2.Balance(), (1000+1)*1000/2; got != want {
		t.Error("Balance = %d, want %d", got, want)
	}
}
