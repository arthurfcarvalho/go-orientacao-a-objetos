package accounts

import "strconv"

type CheckingAccount struct {
	Holder        string
	BranchNumber  int
	AccountNumber int
	Balance       float64
}

func (acc *CheckingAccount) Withdraw(amount float64) string {

	if amount <= 0 {
		return "You cannot withdraw an amount less than or equal to 0!"
	}

	canWithdraw := amount <= acc.Balance

	if canWithdraw {
		acc.Balance -= amount
		return "Withdrawal sucessful. Your new balance is: " + strconv.FormatFloat(acc.Balance, 'f', -1, 64) + "."
	}

	return "Insufficient balance to complete this withdraw."
}

func (acc *CheckingAccount) Deposit(amount float64) (string, float64) {

	if amount <= 0 {
		return "You cannot deposit an amount less than or equal to 0. You tried to deposit:", amount
	}

	acc.Balance += amount
	return "Deposit was sucessful. Your new balance is:", amount
}

func (acc *CheckingAccount) Transfer(amount float64, destination *CheckingAccount) string {

	if amount <= 0 {
		return "You cannot transfer an amount less than or equal to 0."
	}

	if amount > acc.Balance {
		return "You cannot transfer an amount higher than what you have in your balance."
	}

	acc.Balance -= amount
	destination.Balance += amount

	return "Transfer was sucessful! Your new balance is: " + strconv.FormatFloat(acc.Balance, 'f', -1, 64) + "."
}