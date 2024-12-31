package main

import (
	"alura/banco/accounts"
	"alura/banco/customers"
	"fmt"
)

func main() {

	customerArthur := customers.Holder{
		Name:     "Arthur",
		Document: "123456",
		Job:      "Developer",
	}

	arthurAccount := accounts.CheckingAccount{
		Holder:        customerArthur,
		BranchNumber:  15,
		AccountNumber: 123,
	}

	arthurAccount.Deposit(9000)
	fmt.Println(arthurAccount.GetBalance())

	customerLeticia := customers.Holder{
		Name:     "Leticia",
		Document: "987654",
		Job:      "Developer",
	}

	leticiaAccount := accounts.CheckingAccount{
		Holder:        customerLeticia,
		BranchNumber:  30,
		AccountNumber: 987,
	}

	//fmt.Println(arthurAccount.Withdraw(1500))
	//fmt.Println(arthurAccount.Deposit(500))
	fmt.Println(arthurAccount.Transfer(5000, &leticiaAccount))
}
