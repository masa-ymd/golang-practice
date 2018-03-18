package bank1

var deposits = make(chan int) // 入金額を送信するチャネル
var balances = make(chan int) // 残高を受領するチャネル

func Deposit(amount int) { deposits <- amount } // 入金を行う関数

func Balance() int { return <-balances } // 残高をチャネル経由で受領する

func teller() {
	var balance int // balanceはtellerゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits: // depositsチャネルに渡された数字を受領
			balance += amount
		case balances <- balance: // balancesに空きがある場合、このケースが実行される
		}
	}
}

func init() {
	go teller()
}
