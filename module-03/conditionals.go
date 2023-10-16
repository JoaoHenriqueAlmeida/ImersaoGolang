package main

import "fmt"

func main() {
	// Cursed code
	// animal := "dog"
	// if animal == "cat" {
	// 	fmt.Println("Meow")
	// } else if animal == "dog" {
	// 	fmt.Println("Woof woof")
	// } else if animal == "cow" {
	// 	fmt.Println("Moo")
	// } else {
	// 	fmt.Println("Unknown animal")
	// }

	animal := "dog"

	switch animal {
	case "cat":
		fmt.Println("Meow")
	case "dog":
		fmt.Println("Woof woof")
	case "cow":
		fmt.Println("Moo")
	default:
		fmt.Println("Unknown animal")
	}
}
