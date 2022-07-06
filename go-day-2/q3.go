package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	accountBalance int
}

func (bankAc *BankAccount) depositMoney(amount int, wg *sync.WaitGroup, mu *sync.Mutex) {

	mu.Lock()
	fmt.Println("Before Depositing Money, accountBalance", bankAc.accountBalance)
	bankAc.accountBalance += amount
	fmt.Println("After Depositing", amount, " Money accountBalance", bankAc.accountBalance)
	mu.Unlock()
	wg.Done()
}
func (bankAc *BankAccount) withdrawMoney(withdrawAmount int, wg *sync.WaitGroup, mu *sync.Mutex) {

	mu.Lock()
	fmt.Println("Before withdrawing Money, accountBalance", bankAc.accountBalance)
	if v := bankAc.accountBalance - withdrawAmount; v < 0 {
		fmt.Println("Cannot Withdraw the amount")
	} else {

		bankAc.accountBalance -= withdrawAmount
		fmt.Println("After withdrawing", withdrawAmount, " Money, accountBalance", bankAc.accountBalance)
	}
	mu.Unlock()
	wg.Done()
}
func main() {

	bankAc := BankAccount{accountBalance: 5000}

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(4)

	go bankAc.depositMoney(200, wg, mu)
	go bankAc.withdrawMoney(1000, wg, mu)
	go bankAc.withdrawMoney(5000, wg, mu)
	go bankAc.depositMoney(3000, wg, mu)
	wg.Wait()
	fmt.Println("Final Account Balance :", bankAc.accountBalance)
}
