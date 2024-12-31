package main

import (
	"alura/banco/accounts"
	"fmt"
)

func main() {

	arthurAccount := accounts.CheckingAccount{Holder: "Arthur", Balance: 9000}
	leticiaAccount := accounts.CheckingAccount{Holder: "Leticia", Balance: 5000}

	//fmt.Println(arthurAccount.Withdraw(1500))
	//fmt.Println(arthurAccount.Deposit(500))
	fmt.Println(arthurAccount.Transfer(5000, &leticiaAccount))
}
