package main 

import  "fmt"

type OrderStatus string

const (
	DRAFT  OrderStatus = "draft"
	PENDING OrderStatus = "pending"
	SEND OrderStatus = "send"
	APPROVED OrderStatus = "approved"
)

func changeOrderStatus(status  OrderStatus) {
	fmt.Println("order status changed to ",status)
}


func main() {

	changeOrderStatus(PENDING)
}
