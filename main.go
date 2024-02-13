package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	floatNumbers := []float64{1.3, 2.24, 3.10, 4.96}
	double := transformNumbers(&numbers, getTransformerFunction())
	fmt.Println(double)
	trip := transformNumbers(&floatNumbers, triple)
	fmt.Println(trip)

	quadrouple := transformNumbers(&numbers, func(num int) int {
		return num * 4
	})

	fmt.Println(quadrouple)
}

func transformNumbers[T int | float32 | float64](numbers *[]T, transfrom func(T) T) []T {
	dNumbers := []T{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transfrom(val))
	}
	return dNumbers
}

func getTransformerFunction() func(int) int {
	return double
}

func double(number int) int {
	return number * 2
}

func triple(number float64) float64 {
	return number * 3
}
