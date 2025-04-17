package main

import "fmt"

func main() {
	x := 10
	y := &x
	*y = *y + 70
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*y)

}
