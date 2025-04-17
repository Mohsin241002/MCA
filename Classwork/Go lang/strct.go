package main

import "fmt"

type Student struct {
	Name      string
	Age       int
	StudentID int
	Course1   Course
}

type Course struct {
	Coursename   string
	CourseID     int
	CourseCredit int
}
type teacher struct {
	teachername string
	teacherID   int
	Course1     Course
}

func main() {
	// create a student
	student1 := Student{
		Name:      "John Doe",
		Age:       20,
		StudentID: 12345,
		Course1: Course{
			Coursename:   "Math",
			CourseID:     101,
			CourseCredit: 3,
		},
	}

	// print the details
	fmt.Println(student1)
	fmt.Println(student1.Name)
	fmt.Println(student1.Age)
	fmt.Println(student1.StudentID)
	fmt.Println(student1.Course1.Coursename)
	fmt.Println(student1.Course1.CourseID)
	fmt.Println(student1.Course1.CourseCredit)

}
