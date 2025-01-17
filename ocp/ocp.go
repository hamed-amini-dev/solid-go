package ocp

import "fmt"

type Payment interface {
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

type PaymentProcessor interface {
	Pay(p Payment)
}

type Processor struct{}

func (P *Processor) Pay(p Payment) {
	p.ProcessPayment()
}
