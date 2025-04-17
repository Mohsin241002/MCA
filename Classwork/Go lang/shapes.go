package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Circle: %+v\n", circ)

	fmt.Println("Area of rectangle:", rect.Area())
	fmt.Println("Area of circle:", circ.Area())

	shapes := []Shape{rect, circ}
	for _, shape := range shapes {
		fmt.Println("Area of shape:", shape.Area())
	}
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
