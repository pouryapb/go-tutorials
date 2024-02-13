package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	floatNumbers := []float64{1.3, 2.24, 3.10, 4.96}
	doubled := transformNumbers(&numbers, getTransformerFunction())
	fmt.Println(doubled)
	trippled := transformNumbers(&floatNumbers, createTransformer(3.0))
	fmt.Println(trippled)

	quadrouple := transformNumbers(&numbers, func(num int) int {
		return num * 4
	})

	test := func() {
		fmt.Println("this is anonymus")
	}
	test()

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
	return createTransformer(2)
}

func createTransformer[T int | float64](factor T) func(T) T {
	return func(number T) T {
		return number * factor
	}
}
