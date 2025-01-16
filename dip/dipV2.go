package dip

import "fmt"

type NotificationService interface {
	SendNotification() error
}

type notification struct {
}

func NewNotification() NotificationService {
	return &notification{}
}

func (n *notification) SendNotification() error {
	fmt.Println("Sending notification")
	return nil
}

type OrderService interface {
	CreateOrder() error
}

type order struct {
	notificationService NotificationService
}

func NewOrder(notificationService NotificationService) OrderService {
	return &order{
		notificationService: notificationService,
	}
}

func (o *order) CreateOrder() error {
	err := o.notificationService.SendNotification()
	if err != nil {
		return err
	}
	return nil
}
