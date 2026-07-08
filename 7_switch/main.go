package main 

import "fmt"

func main() {

	var marks float64 = 75

	switch {
	case marks >= 80 :
		fmt.Println("Grade: A+")
	case marks >= 60 , marks < 80 :
		fmt.Println("Grade: A")

	case marks >= 50 , marks < 60 :
		fmt.Println("Grade: B")

	case marks >= 40 , marks < 50 :
		fmt.Println("Grade: D")	
	default:
		fmt.Println("Grade: F")
	}

}