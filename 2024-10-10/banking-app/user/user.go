package user

import (
	"bankingApp/helper"
	"errors"
)

var allUsers []*User
var userIdCounter = 0

type User struct {
	userId    int
	isAdmin   bool
	firstName string
	lastName  string
	customer  *Customer
}

//Factory

func NewAdminUser(firstName, lastName string) (*User, error) {
	firstName = helper.RemoveAllLeadingAndTrailingSpaces(firstName)
	lastName = helper.RemoveAllLeadingAndTrailingSpaces(lastName)
	err := helper.ValidateAll(
		validateFirstName(firstName),
		validateLastName(lastName))
	if err != nil {
		return nil, err
	}
	tempUserObject := &User{
		userId:    userIdCounter,
		isAdmin:   true,
		firstName: firstName,
		lastName:  lastName,
		customer:  nil,
	}
	userIdCounter++
	allUsers = append(allUsers, tempUserObject)
	return tempUserObject, nil
}

func NewCustomerUser(firstName, lastName string, customerParameters map[string]interface{}) (*User, error) {
	firstName = helper.RemoveAllLeadingAndTrailingSpaces(firstName)
	lastName = helper.RemoveAllLeadingAndTrailingSpaces(lastName)
	err := helper.ValidateAll(
		validateFirstName(firstName),
		validateLastName(lastName))
	if err != nil {
		return nil, err
	}
	tempUserObject := &User{
		userId:    userIdCounter,
		isAdmin:   true,
		firstName: firstName,
		lastName:  lastName,
		customer:  newCustomer(customerParameters),
	}
	userIdCounter++
	return tempUserObject, nil
}

// validations

func validateInputCustomerId(userId int) error {
	if userId < 0 || userIdCounter >= userId {
		return errors.New("Invalid User Id")
	}
	return nil
}
func validateFirstName(firstName string) error {
	if len(firstName) < 2 {
		return errors.New("length of First Name should atleast be 2")
	}
	return nil
}
func validateLastName(lastName string) error {
	if len(lastName) == 0 {
		return errors.New("last Name Cannot be Empty")
	}
	return nil
}

func validateIfAdmin(user *User) error {

	if !user.IsUserAdmin() {
		return errors.New("unauthoried Access")
	}
	return nil
}
func validateIfCustomer(user *User) error {

	if user.IsUserAdmin() {
		return errors.New("Only Customers can access this")
	}
	return nil
}

// getters
func (user *User) GetUserId() int {
	return user.userId
}
func (user *User) IsUserAdmin() bool {
	return user.isAdmin
}
func (user *User) GetFullName() string {
	return user.firstName + " " + user.lastName
}

func (user *User) GetUserById(userId int) (*User, error) {
	err := helper.ValidateAll(
		validateIfAdmin(user),
		validateInputCustomerId(userId))
	if err != nil {
		return nil, err
	}

	return allUsers[userId], nil

}
