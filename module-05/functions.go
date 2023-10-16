package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
}

func main() {
	p := Person{
		Name: "Henrique",
	}

	// passing reference instead of value
	showPerson(&p)

	total, err := convertAndSum("1", "2", "5", "9", "B")

	fmt.Println(total, err)
}

// receiving reference instead of value
func showPerson(p *Person) {
	fmt.Println(p)
}

// naming returns, pretty cool stuff
func convertAndSum(s ...string) (total int, err error) {
	var n int
	for _, v := range s {
		n, err = strconv.Atoi(v)
		if err != nil {
			total = 0
			break
		}
		total += n
	}
	return
}
