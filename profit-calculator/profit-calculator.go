package main

import "fmt"

func main() {
	revenue := getInputFloat64("Revenue: ")
	expenses := getInputFloat64("Expenses: ")
	taxRate := getInputFloat64("Tax Rate: ")

	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

}

func getInputFloat64(message string) float64 {
	var result float64
	fmt.Print(message)
	fmt.Scan(&result)
	return result
}

func calculateValues(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
