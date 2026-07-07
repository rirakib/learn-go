package main 

import "fmt"

func main() {

	// i := 1

	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i = range 5 {
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
		
	// 	if i == 5 {
	// 		break
	// 	}
		
	// 	fmt.Println(i)
	// }


	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}