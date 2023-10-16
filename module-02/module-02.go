package main

import "fmt"

type Person struct {
	Name    string
	Age     uint
	Address string
}

func main() {
	person := Person{
		Name:    "Henrique Almeida",
		Age:     23,
		Address: "Random street",
	}

	// var array [2]string;
	array := [...]string{"First element", "Second element"}
	slice := make([]string, 3)
	slice = append(slice, "First element")

	fmt.Println(person, slice, array)
}
