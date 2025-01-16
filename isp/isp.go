package isp

import "fmt"

type work interface {
	Work() error
}

type eat interface {
	Eat() error
}

type Human interface {
	work
	eat
}

type Robot interface {
	work
}

type iHuman struct {
	name   string
	family string
}

func NewHuman() Human {
	return &iHuman{}
}
func (h *iHuman) Work() error {
	fmt.Println("Human Work")
	return nil
}
func (h *iHuman) Eat() error {
	fmt.Println("Human eat")
	return nil
}

type iRobot struct {
	name string
	code string
}

func NewRobot() Robot {
	return &iRobot{}
}

func (r *iRobot) Work() error {
	fmt.Println("Robot Work")
	return nil
}
