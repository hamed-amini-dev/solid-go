package srp

type EmailNotification interface {
	Notify()
}

type emailNotification struct {
}

func NewEmailNotification() EmailNotification {
	return &emailNotification{}
}

func (e *emailNotification) Notify() {
	println("email notification")
}

type User struct {
	name  string
	notif EmailNotification
}

func NewUser(notification EmailNotification) *User {
	return &User{
		notif: notification,
	}
}

func (u *User) Notify() {
	u.notif.Notify()
}
