package main

import (
	"fmt"
	"time"
)

type order struct {
	orderId      int
	customerName string
	amount       float64
	status       string
	createdAt    time.Time
}

func (o order) newOrder(orderId int, customerName string, amount float64, status string) *order {

	order := order{
		orderId:      orderId,
		customerName: customerName,
		amount:       amount,
		status:       status,
		createdAt:    time.Now(),
	}

	return &order
}

func (o *order) changeOrderStatus(status string) {
	o.status = status
}

func (o *order) changeCustomerName(customerName string) {
	o.customerName = customerName
}

func main() {

	order1 := order{}.newOrder(1, "Rakibul Islam", 150.00, "pending")
	order1.changeOrderStatus("completed")

	order2 := order{}.newOrder(2, "Sajib Islam", 168.00, "refunded")
	order2.changeOrderStatus("approved")

	fmt.Println(order1.customerName)
	fmt.Println(order2.status)
}
