package main

import (
	"fmt"
	"strconv"
)

type CheckingAccount struct {
	holder        string
	branchNumber  int
	accountNumber int
	balance       float64
}

func (c *CheckingAccount) withdraw(amount float64) string {

	if amount <= 0 {
		return "You can't withdraw a value lower or equal to 0!"
	}

	canWithdraw := amount <= c.balance

	if canWithdraw {
		c.balance -= amount
		return "The withdrawal was sucessful." + "Your new balance is: " + strconv.FormatFloat(c.balance, 'f', -1, 64) + "."
	}

	return "You don't have enough balance in your account to withdraw this amount."
}

func main() {

	arthurAccount := CheckingAccount{holder: "Arthur", balance: 9000}

	fmt.Println(arthurAccount.withdraw(1500))

	// i'll delete the comments under this one after i finish this 2nd class.

	// arthurAccount := CheckingAccount{holder: "Arthur", balance: 56495.50}
	//arthurAccount := CheckingAccount{"Arthur", 15, 123456789, 56495.50}

	//secondArthurAccount := CheckingAccount{"Arthur", 40, 123456789, 56495.50}

	//fmt.Println(arthurAccount)
	//fmt.Println(arthurAccount == secondArthurAccount)

	/* var newAccount *CheckingAccount
	newAccount = new(CheckingAccount)
	newAccount.holder = "Random Person"
	newAccount.balance = 900

	var randomAccount *CheckingAccount
	randomAccount = new(CheckingAccount)
	randomAccount.holder = "Random Person"
	randomAccount.balance = 900

	fmt.Println(&newAccount)
	fmt.Println(&randomAccount)

	fmt.Println(newAccount == randomAccount)
	fmt.Println(*newAccount == *randomAccount) */
}
