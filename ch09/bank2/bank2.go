package bank2

var (
	sema    = make(chan struct{}, 1) // balanceを保護するバイナリセマフォ
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // トークンを取得
	balance = balance + amount
	<-sema // トークンを開放
}

func Balance() int {
	sema <- struct{}{} // トークンを獲得
	b := balance
	<-sema
	return b
}
