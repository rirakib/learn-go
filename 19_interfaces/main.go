package main

import "fmt"


// Contract
type Paymenter interface {
	Pay(amount float32)
}

// Registry
var gateways = make(map[string]Paymenter)
// Register a gateway
func Register(name string, gateway Paymenter) {
	gateways[name] = gateway
}

// Factory
func GetGateway(name string) (Paymenter, error) {

	gateway, ok := gateways[name]

	if !ok {
		return nil, fmt.Errorf("%s gateway not found", name)
	}

	return gateway, nil
}

// Service
type Payment struct{}
func (Payment) MakePayment(amount float32, gateway string) {

	paymentGateway, err := GetGateway(gateway)

	if err != nil {
		fmt.Println(err)
		return
	}

	paymentGateway.Pay(amount)
}



//bkash
type Bkash struct{}
func (Bkash) Pay(amount float32) {
	fmt.Println("Bkash Payment :", amount)
}

// automatically register
func init() {
	Register("bkash", Bkash{})
}


//paypal 
type Paypal struct{}
func (Paypal) Pay(amount float32) {
	fmt.Println("Paypal Payment :", amount)
}

func init() {
	Register("paypal", Paypal{})
}


//stripe
type Stripe struct{}
func (Stripe) Pay(amount float32) {
	fmt.Println("Stripe Payment :", amount)
}

// automatically register
func init() {
	Register("stripe", Stripe{})
}


//binance
type Binance struct{}
func (Binance) Pay(amount float32) {
	fmt.Println("Binance Payment :", amount)
}

// automatically register
func init() {
	Register("binance", Binance{})
}



func main() {

	payment := Payment{}

	payment.MakePayment(100, "stripe")
	payment.MakePayment(250, "paypal")
	payment.MakePayment(500, "bkash")
	payment.MakePayment(499, "binance")

	// Invalid gateway
	payment.MakePayment(1000, "sslcommerz")
}