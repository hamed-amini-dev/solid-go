package main

import (
	"github.com/hamed-amini-dev/solid-go/dip"
	"github.com/hamed-amini-dev/solid-go/isp"
	"github.com/hamed-amini-dev/solid-go/srp"
)

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

	// isp

	human := isp.NewHuman()
	robot := isp.NewRobot()

	err = human.Work()
	if err != nil {
		panic(err)
	}

	err = robot.Work()
	if err != nil {
		panic(err)
	}

	//  srp

	srpN := srp.NewEmailNotification()
	user := srp.NewUser(srpN)
	user.Notify()

	//

}
