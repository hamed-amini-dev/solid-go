package lsp

import "fmt"

type Shape interface {
	Area() float64
}

type IShape struct {
}

func (IShape) CalculateArea(s Shape) {
	fmt.Println(fmt.Sprintf("%f", s.Area()))
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}
