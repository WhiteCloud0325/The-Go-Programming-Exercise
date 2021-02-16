//Package bank provide a concurrency-safe bank with one account
package bank

var deposits = make(chan int) //send amount to deposite
var balances = make(chan int) //receive balance
var withdraws = struct {
	amount chan int
	res    chan bool
}{make(chan int), make(chan bool)}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	withdraws.amount <- amount
	return <-withdraws.res
}

func teller() {
	var balance int //balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws.amount:
			if balance >= amount {
				balance -= amount
				withdraws.res <- true
			} else {
				withdraws.res <- false
			}
		}
	}
}

func init() {
	go teller()
}
