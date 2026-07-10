package main

import "fmt"

type Customer struct {
	ID    int
	Name  string
	Email string
}

type Order struct {
	ID    int
	Price float64
	Customer
}

func main() {
	newOrder := Order{
		ID:    1,
		Price: 100.50,
		Customer: Customer{
			ID:    1,
			Name:  "John Doe",
			Email: "john@gmail.com",
		},
	}

	fmt.Println(newOrder.Customer.Name)
}
