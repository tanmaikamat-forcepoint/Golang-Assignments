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

	validateDateOfBirth(dateOfBirth)
	validateYearOfPassingIsAfterYearOfEnrollment(yearOfEnrollment, yearOfPassing)

	var semesterGrades []string
	var finalGrade string
	var finalCGPA float32
	semesterGrades = calculateSemesterGradesFromSemesterCGPA(semesterCGPA)
	finalCGPA = calculateFinalCGPAFromSemesterCGPA(semesterCGPA)
	finalGrade = getGradeFromCGPA(finalCGPA)

	temporaryStudentObject := &student{

		firstName:               firstName,
		lastName:                lastName,
		fullName:                getFullNameFromFirstAndLastName(firstName, lastName),
		dateOfBirth:             dateOfBirth,
		age:                     getAgeFromDOB(dateOfBirth),
		semesterCGPA:            semesterCGPA,
		finalCGPA:               finalCGPA,
		semesterGrades:          semesterGrades,
		finalGrade:              finalGrade,
		yearOfEnrollment:        yearOfEnrollment,
		yearOfPassing:           yearOfPassing,
		numberOfYearsToGraduate: getNumberOfYearsToGraduateFromPassingYear(yearOfPassing),
	}
	allStudents = append(allStudents, temporaryStudentObject)
	return temporaryStudentObject

}

func (s *student) UpdateStudent(parameter string, value interface{}) {
	switch parameter {
	case "firstName":
		tempFirstNameVariable, stringValidation := value.(string)
		if !stringValidation {
			panic("Please enter Correct First Name")
		}
		s.firstName = tempFirstNameVariable
		s.fullName = getFullNameFromFirstAndLastName(s.firstName, s.lastName)
	case "lastName":
		tempLastNameVariable, stringValidation := value.(string)
		if !stringValidation {
			panic("Please enter Correct Last Name")
		}
		s.lastName = tempLastNameVariable
		s.fullName = getFullNameFromFirstAndLastName(s.firstName, s.lastName)
	case "dateOfBirth":
		tempDateOfBirthVariable, stringValidation := value.(string)
		if !stringValidation {
			panic("Date of Birth should be a String")
		}
		validateDateOfBirth(tempDateOfBirthVariable)
		s.dateOfBirth = tempDateOfBirthVariable
		s.age = getAgeFromDOB(tempDateOfBirthVariable)
	case "yearOfEnrollment":
		tempYearOfEnrollmentVariable, intValidation := value.(int)
		if !intValidation {
			panic("Year of Enrollment should be Integer")
		}
		validateYearOfPassingIsAfterYearOfEnrollment(tempYearOfEnrollmentVariable, s.yearOfPassing)
		s.yearOfEnrollment = tempYearOfEnrollmentVariable
	case "yearOfPassing":
		tempYearOfPassingVariable, intValidation := value.(int)
		if !intValidation {
			panic("Year of Passing should be Integer")
		}
		validateYearOfPassingIsAfterYearOfEnrollment(s.yearOfEnrollment, tempYearOfPassingVariable)
		s.yearOfPassing = tempYearOfPassingVariable
		s.numberOfYearsToGraduate = getNumberOfYearsToGraduateFromPassingYear(tempYearOfPassingVariable)
	case "semesterCGPA":
		tempSemesterCPGAVariable, floatSliceValidation := value.([]float32)
		if !floatSliceValidation {
			panic("Invalid Array of CGPA")
		}
		validateAllCGPAs(tempSemesterCPGAVariable)
		s.semesterCGPA = tempSemesterCPGAVariable
		s.semesterGrades = calculateSemesterGradesFromSemesterCGPA(tempSemesterCPGAVariable)

		s.finalCGPA = calculateFinalCGPAFromSemesterCGPA(tempSemesterCPGAVariable)
		s.finalGrade = getGradeFromCGPA(s.finalCGPA)
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

func calculateSemesterGradesFromSemesterCGPA(semesterCGPA []float32) []string {
	var semesterGrades []string
	for _, CGPA := range semesterCGPA {
		semesterGrades = append(semesterGrades, getGradeFromCGPA(CGPA))
	}
	return semesterGrades
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

func validateYearOfPassingIsAfterYearOfEnrollment(yearOfEnrollment int, yearOfPassing int) {
	if yearOfPassing < yearOfEnrollment {
		panic("Year of Passing Cannot be before Year of Enrollment")
	}
}
