package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1)
	hobbies := [3]string{"Guitar", "Games", "Coding"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4)
	mainHobbies = mainHobbies[1:cap(mainHobbies)]
	fmt.Println(mainHobbies)

	// 5)
	courseGoals := []string{"learning go?", "bulding cool stuff :D"}
	fmt.Println(courseGoals)

	// 6)
	courseGoals[1] = "create cool stuff :>"
	courseGoals = append(courseGoals, "understand other go codes?")
	fmt.Println(courseGoals)

	// 7)
	products := []Product{
		{
			id:    "id-1",
			title: "milk",
			price: 6.99,
		},
		{
			id:    "id-2",
			title: "meat",
			price: 99.99,
		},
	}
	fmt.Println(products)

	newProduct := Product{
		"id-3",
		"cups",
		8.5,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}
