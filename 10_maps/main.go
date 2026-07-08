package main

import (
	"fmt"
	"maps"
)

func main() {

	data := make(map[string]string)

	data["name"] = "John Doe"
	data["email"] = "john@gmail.com"
	data["phone"] = "123-456-7890"

	delete(data, "phone")
	fmt.Println("Name:", data["name"], "Email:", data["email"], "Phone:", data["phone"], "Address ", data["address"])

	clear(data)
	fmt.Println(data)

	mark := map[string]float64{
		"Math":    90.5,
		"Science": 85.0,
		"English": 92.0,
	}

	fmt.Println("Math mark:", mark["Math"])

	value, ok := mark["Math"]

	fmt.Println(value)

	if ok {
		fmt.Println("Subject mark is present.")
	} else {
		fmt.Println("Subject mark is not present.")
	}


	test1 := map[string]int{
		"one":   1,"two":   2,
		"three": 3,"four":  4,
		"five":  5,"six":   6,
	}

	test2 := map[string]int{
		"one":   1,"two":   2,
		"three": 3,"four":  4,
		"five":  5,"six":   6,
	}

	fmt.Println(maps.Equal(test1,test2))

}
