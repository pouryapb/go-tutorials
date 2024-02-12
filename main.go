package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/pouryapb/go-tutorials/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Address: ", randomdata.Address())
	fmt.Println()

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Deposit value: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdraw value: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawlAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdrawl more than you have.")
				continue
			}

			accountBalance -= withdrawlAmount
			fmt.Println("Balance updated! New balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 4:
			fmt.Println("Goodbye! 👋")
			fmt.Println("Thanks for using our bank. 😊")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}

}
