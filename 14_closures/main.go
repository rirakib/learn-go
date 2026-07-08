package main 


import "fmt"

func increment() func() int {

	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {

	increment := increment()
	
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}