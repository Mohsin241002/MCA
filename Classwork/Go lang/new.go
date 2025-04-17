// First Go program
package main

import "fmt"

// Main function
func main() {
	// Printing a simple message
	fmt.Println("!... Hello World ...!")

	// Declaring constants
	const name, age = "Kim", 22
	const age1, age2 = 20, 23
	const averageAge = (age1 + age2) / 2

	// Printing constants
	fmt.Println(name, "is", age, "years old.")
	fmt.Println("The average age of two people is", averageAge, "years.")

	// Using Printf and Sprintf
	fmt.Printf("The average age of two people is %v\n", averageAge)
	msg := fmt.Sprintf("%s is %d years old.\n", name, age)
	fmt.Println(msg)

	// Declaring variables with zero values
	var age3 float32
	var name1 string
	fmt.Printf("age3: %f\n", age3)
	fmt.Println("name1:", name1)

	// Working with complex numbers
	var complexNum complex128 = complex(2, 3)
	fmt.Println("The complex number is:", complexNum)

	// Type conversion
	var experience = 3.6
	fmt.Printf("Experience: %f\n", experience)
	var experience1 = int(experience)
	fmt.Printf("Experience1: %v\n", experience1)
}
