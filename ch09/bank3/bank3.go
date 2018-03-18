package bank3

import (
	"sync"
)

var (
	mu      sync.RWMutex // balanceを保護する(Readは平行稼働可能にする)
	balance int
)

func Deposit(amount int) {
	mu.Lock()         // トークンを取得
	defer mu.Unlock() // トークンを開放
	deposit(amount)
}

func Balance() int {
	mu.RLock()         // トークンを獲得
	defer mu.RUnlock() // トークンを開放
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 { // 残高不足
		deposit(amount)
		return false
	}
	return true
}

// この関数はロックが獲得されていることを前提としている
func deposit(amount int) { balance += amount }
