package main

import (
	"fmt"
	"slices"
)

func main() {

	var num = make([]int, 0, 5)
	var num2 = make([]int, 0, 5)


	for i := 0; i < 20; i++ {
		num = append(num, i)
	}

	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))

	fmt.Println(num[5:10])
	fmt.Println(num[:10])

	copy(num2, num)

	fmt.Println(num2)

	var test1 = []int{1, 2, 3, 4, 5}
	var test2 = []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Equal(test1,test2))
}