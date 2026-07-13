package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("task complete number", id)
}


func main() {

	for i := 1; i <= 10; i++ {
		// go task(i)

		go func(orderNumber int){
			fmt.Println("Task completed number ",orderNumber)
		}(i)
	}

	time.Sleep(time.Second * 2)
}

