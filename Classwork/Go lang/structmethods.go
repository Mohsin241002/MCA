package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade string
}

func (s Student) PrintDetails() {
	fmt.Printf("Name: %s, Age: %d, Grade: %s\n", s.Name, s.Age, s.Grade)
}

func (s Student) GetAge() int {
	return s.Age
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func main() {
	student := Student{
		Name:  "John Doe",
		Age:   20,
		Grade: "A",
	}

	student.PrintDetails()
	fmt.Println("Age:", student.GetAge())
	student.SetAge(21)
	student.PrintDetails()
}
