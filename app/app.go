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

	// dpv2
	notif := dip.NewNotification()
	dpv2 := dip.NewOrder(notif)
	err := dpv2.CreateOrder()
	if err != nil {
		panic(err)
	}

}
