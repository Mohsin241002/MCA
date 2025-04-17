package main

import "fmt"

func main() {
	// var num1, num2 int

	// fmt.Print("Enter the first number: ")
	// fmt.Scanln(&num1)
	// fmt.Print("Enter the second number: ")
	// fmt.Scanln(&num2)
	// if num1 > num2 {
	// 	fmt.Printf("The greater number is: %d\n", num1)
	// } else if num2 > num1 {
	// 	fmt.Printf("The greater number is: %d\n", num2)
	// } else {
	// 	fmt.Println("Both numbers are equal.")
	// }
	// fmt.Println(test(45, 85))
	a := 10
	b := 20
	fmt.Println("Before swapping: a1 =", a, ", b1 =", b)
	swap(&a, &b)
	fmt.Println("After swapping: a1 =", a, ", b1 =", b)

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array elements:", arr)

	slc := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice elements:", slc)

	arr[2] = 10
	fmt.Println("Modified array:", arr)

	slc[2] = 10
	fmt.Println("Modified slice:", slc)

	slc = append(slc, 6, 7, 8)
	fmt.Println("Slice after appending:", slc)

	slc = append(slc[:3], slc[4:]...)
	fmt.Println("Slice after deleting:", slc)
}

// func test(a1 int, b1 int) int {
// 	c1 := a1 + b1
// 	return c1
// }

func swap(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
