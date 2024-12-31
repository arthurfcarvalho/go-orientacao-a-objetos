package main

import "fmt"

type CheckingAccount struct {
	holder        string
	branchNumber  int
	accountNumber int
	balance       float64
}

func main() {
	arthurAccount := CheckingAccount{"Arthur", 15, 123456789, 56495.50}

	fmt.Println(arthurAccount)
}
