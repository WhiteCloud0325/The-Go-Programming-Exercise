//Package bank provide a concurrency-safe bank with one account
package bank

var deposits = make(chan int) //send amount to deposite
var balances = make(chan int) //receive balance

func Deposite(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int //balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
