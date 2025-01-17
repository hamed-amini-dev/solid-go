package ocp

import "fmt"

type PaymentProcessor interface {
	ProcessPayment()
}

type PayPal struct{}

func (p *PayPal) ProcessPayment() {
	fmt.Println("PayPal")
}

type Stripe struct{}

func (s *Stripe) ProcessPayment() {
	fmt.Println("Stripe")
}

func Pay(p PaymentProcessor) {
	p.ProcessPayment()
}
