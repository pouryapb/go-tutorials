package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["ic"] = "https://isee.com"
	fmt.Println(websites)

	fmt.Println(websites)

	list := make([]float64, 3, 5)
	list[0] = 3
	list[1] = 1
	list[2] = 6

	for i, val := range list {
		fmt.Printf("Index: %v, Value: %v\n", i, val)
	}

	for key, val := range websites {
		fmt.Printf("Key: %v\nValue: %v\n", key, val)
	}
}
