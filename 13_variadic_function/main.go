package main 

import "fmt"

func sum(numbers ...int) int {

	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5))

	marks := []int{90, 85, 92, 88, 95}

	fmt.Println(sum(marks...))
}