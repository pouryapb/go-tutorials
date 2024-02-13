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

	delete(websites, "Google")
	delete(websites, "Amazon Web Services")
	fmt.Println(websites)
}
