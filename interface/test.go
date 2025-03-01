package main

import "fmt"

type paymenter interface {
	pay(amout float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amout float32) {
	p.gateway.pay(amout)
}

// by esewa
type esewa struct{}

func (e esewa) pay(amount float32) {
	fmt.Println("Payment by esewa,", amount)
}

// by bank
type bank struct{}

func (b bank) pay(amount float32) {
	fmt.Println("Payment by esewa,", amount)
}

// khalti
type khalti struct{}

func (k khalti) pay(amount float32) {
	fmt.Println("Payment by khalti", amount)
}

func main() {
	bankPayment := bank{}
	myPayment := payment{
		gateway: bankPayment,
	}

	myPayment.makePayment(100)
}
