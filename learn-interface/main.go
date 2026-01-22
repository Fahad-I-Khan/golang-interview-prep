package main

import "fmt"

type Paypal struct {
	Name  string
	Email string
}

type Payment interface {
	Pay(amount float64, email string) string
	CreateOrder(amount float64, email string) string
}

func (p Paypal) ProcessPayment(amount float64, email string) string {
	return p.Pay(amount, email)
}

func CreatePayment (p Payment, amount float64, email string) string {
	return p.Pay(amount, email)
}

func (p Paypal) Pay(amount float64, email string) string {
	return "Paid " + formatAmount(amount, email) + " using Paypal"
}

func (p Paypal) CreateOrder(amount float64, email string) string {
	return "Order created for " + formatAmount(amount, email) + " using Paypal"
}

func formatAmount(amount float64, email string) string {
	return "$" + fmt.Sprintf("%.2f", amount) + " for " + email
}

func main() {
	paypal := Paypal{Name: "John Doe", Email: "abc@gmail.com"}
	fmt.Println(paypal.Pay(100.0, paypal.Email))
	fmt.Println(paypal.CreateOrder(100.0, paypal.Email))
	paypal.ProcessPayment(150.0, paypal.Email)
	fmt.Println(CreatePayment(paypal, 200.0, paypal.Email))
	
}