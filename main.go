package main

import "fmt"

func main() {
	accountBalance := 1000.00

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		fmt.Print("Deposit value: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Balance updated! New balance:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdraw value: ")
		var withdrawlAmount float64
		fmt.Scan(&withdrawlAmount)

		if withdrawlAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		if withdrawlAmount > accountBalance {
			fmt.Println("Invalid amount. You can't withdrawl more than you have.")
			return
		}

		accountBalance -= withdrawlAmount
		fmt.Println("Balance updated! New balance:", accountBalance)
	} else {
		fmt.Println("Goodbye! 👋")
	}
}
