package main 

import "fmt"

func main() {

	var age int = 20;

	if age >= 18 {
		fmt.Println("You are an adult.")
	}else if age >= 13 {
		fmt.Println("You are a teenager.")
	}else {
		fmt.Println("You are a child.")
	}

	var username string = "admin"
	var password string = "password123"

	if username == "admin" && password == "password123" {
		fmt.Println("Login successful!")
	}else {
		fmt.Println("Invalid username or password.")
	}

}