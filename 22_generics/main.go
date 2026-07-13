package main

import "fmt"


func printSlices[T string | int](items []T) {

	for _,item := range items {
		fmt.Println(item)
	}

}


func main() {

	nums := []int{1,2,3}
	names := []string{"php","go","javascript","python"}

	printSlices(nums)
	printSlices(names)

	
}