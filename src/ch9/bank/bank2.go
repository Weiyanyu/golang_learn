package bank

import "sync"

import "fmt"

var (
	sema      = make(chan struct{}, 1)
	balances2 int
)

func deposit2(amount int) {
	sema <- struct{}{}
	balances2 += amount
	<-sema
}

func balance2() int {
	sema <- struct{}{}
	b := balances2
	<-sema
	return b
}

func Bank2() {
	var w sync.WaitGroup

	w.Add(1)
	go func() {
		defer w.Done()
		deposit2(100)
		fmt.Println(balance2())
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		deposit2(200)
	}()

	w.Wait()
}
