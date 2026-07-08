package main 

import "fmt"

func changeNumber(num int){

	num = 10

	fmt.Println("Change Number ",num)

}

func referenceChangeNumber(num *int){

	*num = 15
    fmt.Println("Reference Change Number ",*num)
}

func main() {

	num := 1

	fmt.Println("Before Change Number ",num)

	changeNumber(num)

	fmt.Println("After Change Number ",num)

	referenceChangeNumber(&num)

	fmt.Println("After Reference Change Number ",num)

}