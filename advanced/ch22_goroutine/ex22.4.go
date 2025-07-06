package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock() //한번 획득한 Mutex는 반드시 반환해야한다.
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should be positive : %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			DepositAndWithdraw(account)
		}()
		wg.Done()
	}
	wg.Wait()
}
