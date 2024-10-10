package dip

import "log"

// Abstraction (interface) that both high-level and low-level modules depend on
type PaymentProcessors interface {
	Pay(amount int) error
}

// Low-level module 1 (concrete implementation)
type Paypal struct {
}

func (p *Paypal) Pay(amount int) error {
	log.Println("Paypal.Pay")
	return nil
}

// Low-level module 2 (another implementation)
type Stripe struct {
}

func (p *Stripe) Pay(amount int) error {
	log.Println("Stripe.Pay")
	return nil
}

// High-level module (business logic)
type Payment struct {
	PayProcessors PaymentProcessors
}

// High-level module now depends on an abstraction (interface)
func (p *Payment) Paying() {
	err := p.PayProcessors.Pay(14)
	if err != nil {
		panic(err)
	}
}
