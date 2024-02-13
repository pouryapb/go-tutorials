package main

import "fmt"

func main() {
	var prices []float64 = []float64{1.2, 2, 3.98}
	prices = append(prices, 3, 3, 4, 5, 9)

	fmt.Println(prices[1:3])
	fmt.Println(prices[1:])
	fmt.Println(prices[:3])
	fmt.Println(len(prices))
	fmt.Println(prices)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	slice := prices[:3]
	slice[1] = 69
	fmt.Println(prices)
}
