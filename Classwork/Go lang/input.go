package main

import (
	"errors"
	"fmt"
)
,
func getNumber() (int, error) {
	var number int
	fmt.Println("Enter a number:")
	_, err := fmt.Scan(&number)
	if err != nil {
		return 0, errors.New("Invalid input")
	}
	if number < 0 {
		return 0, errors.New("Negative number")
	}
	return number + 10, nil
}

func main() {
	number, err := getNumber()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("You entered:", number)
}
