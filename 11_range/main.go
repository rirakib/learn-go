package main 

import "fmt"

func main() {

	data := map[string]string{
		"name":    "John Doe",
		"email":   "john@gmail.com",
		"phone":   "123-456-7890",
	}

	for key, value := range data {
		fmt.Println(key, ":", value)
	}

	for index, code := range "abcdefghijklmnopqrstuvwxyz" {
		fmt.Println(index, ":", string(code), ":", code)
	}
}