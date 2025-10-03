package main 

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	account_number  string
	holder_name string
	balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount 
}

func (b *BankAccount) Withdraw(amount float64) error {
	if b.balance >= amount {
		b.balance -= amount
		return nil
	} else {
		return errors.New("Недостаточно средств")
	}
	
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance 
}

func main () {
	var acc = BankAccount{"880055","Владимир Александрович Злотоусов", 7777.0}
	fmt.Println("Начальный баланс: ", acc.GetBalance())
	acc.Withdraw(5000)
	fmt.Println("Баланс после снятия: ", acc.GetBalance())
	acc.Deposit(1234)
	fmt.Println("Баланс после пополнения: ", acc.GetBalance())
	acc.Withdraw(1111)

}