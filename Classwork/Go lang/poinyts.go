package main

import "fmt"

type Point struct {
	X float32
	Y float32
}

type circle struct {
	radius float32
	center Point
}

// func taransform to change the center of circle to a new point from the origin to +100, +100
func (c *circle) transform() {
	c.center.X += 100
	c.center.Y += 100
}
func main() {
	p := Point{X: 10, Y: 20}
	fmt.Printf("Point X: %f, Y: %f\n", p.X, p.Y)

	// create a circle
	c := circle{
		radius: 5,
		center: Point{},
	}

	fmt.Printf("Circle radius: %f\n", c.radius)
	fmt.Printf("Circle center: X: %f, Y: %f\n", c.center.X, c.center.Y)
	c.transform()
	fmt.Printf("Transformed Circle center: X: %f, Y: %f\n", c.center.X, c.center.Y)
}
