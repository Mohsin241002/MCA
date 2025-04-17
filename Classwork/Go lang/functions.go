package main

import "fmt"

func hello(x, y int) (int, int) {
	defer fmt.Println("Hello from defer")
	fmt.Println("Hello from hello")
	z1 := x + y
	z2 := x - y
	return z1, z2
}

func world(x int) int {
	fmt.Println("Hello from world")
	return x * x
}

func main() {
	fmt.Println(hello(1, 2))
	y := world
	fmt.Println(y(5))

	test := func(x int) int {
		return x * x
	}
	println(test(7))
}
