package main 

import "fmt"

func main() {

	var numbers[5]int = [5]int{}

	numbers[3] = 99

	fmt.Println(numbers)


	var fruites[4]string = [4]string{"Apple", "Banana", "Mango", "Orange"}

	fmt.Println(fruites)


	var marks[5]float64 = [5]float64{75.5, 80.0, 90.0, 85.0, 95.0}

	fmt.Println(marks)

	var matrix[2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(matrix)

	var mixedArray[3]interface{} = [3]interface{}{1, "Hello", 3.14}

	fmt.Println(mixedArray)

}


