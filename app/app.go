package main

import "github.com/hamed-amini-dev/solid-go/dip"

func main() {
	// DIP
	paypal := &dip.Paypal{}
	stripe := &dip.Stripe{}

	// We can inject different payment processors without changing Payment
	p := new(dip.Payment)
	p.PayProcessors = paypal
	p.Paying()
	p.PayProcessors = stripe
	p.Paying()
}
