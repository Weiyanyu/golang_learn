package bank

import "fmt"

import "sync"

type withdrawInfo struct {
	amount int
	ch     chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan withdrawInfo)

func balance() int {
	return <-balances
}

func deposit(amount int) {
	deposits <- amount
}

func withdraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- withdrawInfo{amount, ch}

	return <-ch
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			fmt.Printf("save %d\n", amount)
			balance += amount
		case balances <- balance:
		case wf := <-withdraws:
			if balance < wf.amount {
				wf.ch <- false
			} else {
				balance -= wf.amount
				wf.ch <- true
			}
		}
	}
}

func Bank1() {
	go teller()

	var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer w.Done()
		deposit(100)
		fmt.Println(balance())
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		deposit(200)
		if ok := withdraw(300); !ok {
			fmt.Println("error")
		}
	}()

	w.Wait()

}
