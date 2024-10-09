package university

import (
	"slices"
	"time"
)

var allStudents []*student
var dateFormat string = "02/01/2006"

func GetAllStudents() []*student {
	return allStudents
}

type student struct {
	firstName               string
	lastName                string
	fullName                string
	dateOfBirth             string
	age                     int
	semesterCGPA            []float32
	finalCGPA               float32
	semesterGrades          []string
	finalGrade              string
	yearOfEnrollment        int
	yearOfPassing           int
	numberOfYearsToGraduate int
}

func NewStudent(firstName string,
	lastName string,
	dateOfBirth string,
	semesterCGPA []float32,
	yearOfEnrollment int,
	yearOfPassing int) *student {
	validateFirstName(firstName)
	validateLastName(lastName)
	validateDateOfBirth(dateOfBirth)
	validateYearOfPassingIsAfterYearOfEnrollment(yearOfEnrollment, yearOfPassing)
	validateAllCGPAs(semesterCGPA)

	var semesterGrades []string
	var finalGrade string
	var finalCGPA float32
	var fullName string = getFullNameFromFirstAndLastName(firstName, lastName)
	var age int = getAgeFromDOB(dateOfBirth)
	var numberOfYearsToGraduate int = getNumberOfYearsToGraduateFromPassingYear(yearOfPassing)
	calculateSemesterGradesFromSemesterCGPA(semesterCGPA, &semesterGrades)
	finalCGPA = calculateFinalCGPAFromSemesterCGPA(semesterCGPA)
	finalGrade = getGradeFromCGPA(finalCGPA)

	temporaryStudentObject := &student{

		firstName:               firstName,
		lastName:                lastName,
		fullName:                fullName,
		dateOfBirth:             dateOfBirth,
		age:                     age,
		semesterCGPA:            semesterCGPA,
		finalCGPA:               finalCGPA,
		semesterGrades:          semesterGrades,
		finalGrade:              finalGrade,
		yearOfEnrollment:        yearOfEnrollment,
		yearOfPassing:           yearOfPassing,
		numberOfYearsToGraduate: numberOfYearsToGraduate,
	}
	allStudents = append(allStudents, temporaryStudentObject)
	return temporaryStudentObject

}

func (s *student) UpdateStudent(parameter string, value interface{}) {
	switch parameter {
	case "firstName":
		s.updateFirstName(value)
	case "lastName":
		s.updateLastName(value)
	case "dateOfBirth":
		s.updateDateOfBirth(value)

	case "yearOfEnrollment":
		s.updateYearOfEnrollment(value)
	case "yearOfPassing":
		s.updateYearOfPassing(value)
	case "semesterCGPA":
		s.updateSemesterGrades(value)
	case "default":
		panic("Invalid Choice to Update")

	}
}

func DeleteStudent(s *student) {

	for index, value := range allStudents {
		if value == s {
			allStudents = slices.Delete(allStudents, index, index+1)
			return
		}
	}
	panic("Student Not Found!")

}

// ///////////////////////Package Functions/////////////////////////////////

func getFullNameFromFirstAndLastName(firstName string, lastName string) string {
	return firstName + " " + lastName
}
func getGradeFromCGPA(CGPA float32) string {
	if CGPA > 9.5 {
		return "A"
	} else if CGPA > 8.5 {
		return "B"
	}
	return "C"
}

func getAgeFromDOB(dateOfBirth string) int {

	dob, err := time.Parse(dateFormat, dateOfBirth)

	if err != nil {
		panic(err)
	}

	var age int

	age = time.Now().Year() - dob.Year()
	if time.Now().Month() < dob.Month() {
		age -= 1
	} else if time.Now().Month() == dob.Month() && time.Now().Day() < dob.Day() {
		age -= 1
	}
	if age < 0 {
		panic("Age Cannot be negative")
	}
	return age
}

func getNumberOfYearsToGraduateFromPassingYear(year int) int {
	numberOfYearsToGraduate := year - time.Now().Year()
	if numberOfYearsToGraduate < 0 {
		numberOfYearsToGraduate = 0
	}
	return numberOfYearsToGraduate
}

// ## Cannot use semester Grades as Parameter directly because whenever new Grades are appended, new Memory Location is allocated

func calculateSemesterGradesFromSemesterCGPA(semesterCGPA []float32, semesterGrades *[]string) {

	for _, CGPA := range semesterCGPA {
		*semesterGrades = append(*semesterGrades, getGradeFromCGPA(CGPA))
	}
}

func calculateFinalCGPAFromSemesterCGPA(semesterCGPA []float32) float32 {
	var finalCGPA float32 = 0
	for _, CGPA := range semesterCGPA {
		finalCGPA += CGPA
	}
	if len(semesterCGPA) != 0 {
		finalCGPA = finalCGPA / float32(len(semesterCGPA))
	}
	return finalCGPA

}

/////////////////////////Validation Functions//////////////////////////////

/*
*
Used to Validate Date Of Birth. Panic if Invalid Date or Date in Future
*/
func validateDateOfBirth(dateOfBirth string) {
	dateOfBirthInTimeFormate, err := time.Parse(dateFormat, dateOfBirth)
	if err != nil {
		panic("Invalid Date of Birth")
	}

	if dateOfBirthInTimeFormate.After(time.Now()) {
		panic("Date of Birth Cannot be in Future")
	}
}

func validateAllCGPAs(CGPA []float32) {
	for _, CGPA := range CGPA {
		validateCGPA(CGPA)
	}
}

func validateCGPA(CGPA float32) {
	if CGPA < 0 {
		panic("CGPA cannot be Negative")
	}
}

func validateFirstName(name string) {
	if len(name) == 0 {
		panic("Name Cannot be empty")
	}
}

func validateLastName(name string) {
	if len(name) == 0 {
		panic("Name Cannot be empty")
	}
}

func validateYearOfPassingIsAfterYearOfEnrollment(yearOfEnrollment int, yearOfPassing int) {
	if yearOfEnrollment > time.Now().Year() {
		panic("Year of Enrollment Cannot be in Future")
	}
	if yearOfPassing < yearOfEnrollment {
		panic("Year of Passing Cannot be before Year of Enrollment")
	}
}

/////update Functions/////////////////////////////////////////

func (s *student) updateFirstName(value interface{}) {
	tempFirstNameVariable, stringValidation := value.(string)
	if !stringValidation {
		panic("Please enter Correct First Name")
	}
	validateFirstName(tempFirstNameVariable)
	s.firstName = tempFirstNameVariable
	s.fullName = getFullNameFromFirstAndLastName(s.firstName, s.lastName)
}
func (s *student) updateLastName(value interface{}) {
	tempLastNameVariable, stringValidation := value.(string)
	if !stringValidation {
		panic("Please enter Correct Last Name")
	}
	validateLastName(tempLastNameVariable)
	s.lastName = tempLastNameVariable
	s.fullName = getFullNameFromFirstAndLastName(s.firstName, s.lastName)
}

func (s *student) updateDateOfBirth(value interface{}) {
	tempDateOfBirthVariable, stringValidation := value.(string)
	if !stringValidation {
		panic("Date of Birth should be a String")
	}
	validateDateOfBirth(tempDateOfBirthVariable)
	s.dateOfBirth = tempDateOfBirthVariable
	s.age = getAgeFromDOB(tempDateOfBirthVariable)
}

func (s *student) updateYearOfEnrollment(value interface{}) {
	tempYearOfEnrollmentVariable, intValidation := value.(int)
	if !intValidation {
		panic("Year of Enrollment should be Integer")
	}
	validateYearOfPassingIsAfterYearOfEnrollment(tempYearOfEnrollmentVariable, s.yearOfPassing)
	s.yearOfEnrollment = tempYearOfEnrollmentVariable
}

func (s *student) updateYearOfPassing(value interface{}) {
	tempYearOfPassingVariable, intValidation := value.(int)
	if !intValidation {
		panic("Year of Passing should be Integer")
	}
	validateYearOfPassingIsAfterYearOfEnrollment(s.yearOfEnrollment, tempYearOfPassingVariable)
	s.yearOfPassing = tempYearOfPassingVariable
	s.numberOfYearsToGraduate = getNumberOfYearsToGraduateFromPassingYear(tempYearOfPassingVariable)

}

func (s *student) updateSemesterGrades(value interface{}) {
	tempSemesterCPGAVariable, floatSliceValidation := value.([]float32)
	if !floatSliceValidation {
		panic("Invalid Array of CGPA")
	}
	var tempSemesterGradesVariable []string
	validateAllCGPAs(tempSemesterCPGAVariable)
	s.semesterCGPA = tempSemesterCPGAVariable
	calculateSemesterGradesFromSemesterCGPA(tempSemesterCPGAVariable, &tempSemesterGradesVariable)
	s.semesterGrades = tempSemesterGradesVariable

	s.finalCGPA = calculateFinalCGPAFromSemesterCGPA(tempSemesterCPGAVariable)
	s.finalGrade = getGradeFromCGPA(s.finalCGPA)
}
