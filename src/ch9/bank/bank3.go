package bank

import "sync"

import "fmt"

var (
	mutex     sync.RWMutex
	balances3 int
)

func deposit3(amount int) {
	mutex.Lock()
	defer mutex.Unlock()
	balances3 += amount

}

func balance3() int {
	mutex.RLock()
	defer mutex.RUnlock()
	return balances3
}

func withdraw3(amount int) bool {
	deposit3(-amount)
	if balance3() < 0 {
		deposit3(amount)
		return false
	}
	return true
}

func Bank3() {
	var w sync.WaitGroup

	w.Add(1)
	go func() {
		defer w.Done()
		deposit3(100)
		fmt.Println("save 100")
		fmt.Println(balance3())
		if ok := withdraw3(100); !ok {
			fmt.Println("100 error")
		}
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		deposit3(200)
		fmt.Println("save 200")
		if ok := withdraw3(300); !ok {
			fmt.Println("300 err")
		}

	}()

	w.Wait()
	fmt.Println(balance3())
}
