package main

import (
	"fmt"
	"time"
)

func main() {
	getUserRoleWithSwitchStatement("asddsasad")

	// Using a map is faster, at least for this example
	getUserRoleWithMap("asdsadsaddas")
}

func getUserRoleWithSwitchStatement(role string) {
	start := time.Now()
	switch role {
	case "admin":
		fmt.Println("This User is Admin!")
	case "client":
		fmt.Println("This User is Client!")
	case "salesman":
		fmt.Println("This User is Salesman!")
	default:
		fmt.Println("Ops, this guy doesn\"t have user profile")
	}

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("getUserRoleWithSwitchStatement execution Time: %s\n", elapsed)
}

func getUserRoleWithMap(role string) {
	start := time.Now()

	roles := map[string]string{
		"admin":    "This User is Admin!",
		"client":   "This User is Client!",
		"salesman": "This User is Salesman!",
	}

	// This is a go specific type of syntax that checks the existence of a key in a given map
	if message, ok := roles[role]; ok {
		fmt.Println(message)
	} else {
		fmt.Println("Ops, this guy doesn't have a user profile")
	}
	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("getUserRoleWithMap execution Time: %s\n", elapsed)

}
