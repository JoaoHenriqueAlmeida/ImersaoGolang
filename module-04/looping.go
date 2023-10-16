package main

import "time"

func main() {
	arr := [2]string{"Henrique", "Almeida"}

	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	name := "Henrique Almeida"

	for i := 0; i < len(name); i++ {
		println(name[i], string(name[i]))
	}

	// while
	i := 0
	for i < 10 {
		println((i))
		i++
	}

	// for each
	numberSlice := []uint8{1, 23, 45, 100, 27, 12}
	for _, v := range numberSlice {
		println(v)
	}

	// iterating a map
	personMap := map[string]string{
		"name":             "Henrique Almeida",
		"age":              "24",
		"fiscalIdentifier": "000.000.000-00",
	}

	for k, v := range personMap {
		println(k, v)
	}

	// endless loop for specific tasks
	for {
		println(name)

		time.Sleep(time.Second * 5)
	}

}
