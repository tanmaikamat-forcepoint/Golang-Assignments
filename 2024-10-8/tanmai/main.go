package main

import (
	"fmt"
	"tanmai/university"
)

func main() {
	var testSemesterCGPA []float32

	testStudent := university.NewStudent(
		"Tanmai",
		"Kamat",
		"28/04/2022",
		testSemesterCGPA,
		2012,
		2024)
	fmt.Println(testStudent)
	fmt.Println("All Students :")
	for studentNumber, studentObject := range university.GetAllStudents() {
		fmt.Println(studentNumber, studentObject)
	}
	testSemesterCGPA = append(testSemesterCGPA, 9)
	testSemesterCGPA = append(testSemesterCGPA, 8)
	testStudent.UpdateStudent("yearOfPassing", 2025)
	// testStudent.UpdateStudent("firstName", 10)
	fmt.Println(testStudent)

}
