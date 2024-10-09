package user

import (
	"contactApp/helper"
	"errors"
	"slices"
)

type User struct {
	UserId    int
	FirstName string
	LastName  string
	IsActive  bool
	IsAdmin   bool
}

// /Stores ////////////
var allAdmins []*User
var allStaff []*User
var userId int = 0

// CRUD
func NewAdmin(firstName string, lastName string) (*User, error) {
	err := helper.ValidateAll(validateFirstName(firstName),
		validateLastName(lastName))
	if err != nil {
		return nil, err
	}
	tempNewAdminObject := &User{
		UserId:    userId,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
		IsAdmin:   true,
	}
	userId++
	allAdmins = append(allAdmins, tempNewAdminObject)
	return tempNewAdminObject, nil

}

func (u *User) NewStaff(firstName string, lastName string) (*User, error) {
	if !u.IsAdmin {
		return nil, errors.New("new staff can only be created by Admins")
	}
	err := helper.ValidateAll(
		validateFirstName(firstName),
		validateLastName(lastName))
	if err != nil {
		return nil, err
	}
	tempNewStaffObject := &User{
		UserId:    userId,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
		IsAdmin:   false,
	}
	userId++
	allStaff = append(allAdmins, tempNewStaffObject)
	return tempNewStaffObject, nil

}

func (u *User) GetAllStaff() ([]*User, error) {
	if !u.IsAdmin {
		return nil, errors.New("unauthorized Access")
	}

	return allStaff, nil
}

func (u *User) GetStaffByID(userId int) (*User, error) {
	if !u.IsAdmin {
		return nil, errors.New("unauthorized Access")
	}
	for _, userObject := range allStaff {
		if userObject.UserId == userId {
			return userObject, nil
		}
	}

	return nil, errors.New("User Not Found")
}
func (u *User) UpdateStaffByID(userId int, parameter string, value interface{}) error {
	if !u.IsAdmin {
		return errors.New("unauthorized Access")
	}
	tempStaffObject, err := u.GetStaffByID(userId)
	if err != nil {
		return err
	}
	switch parameter {
	case "firstName":
		return tempStaffObject.updateFirstName(value)
	case "lastName":
		return tempStaffObject.updateLastName(value)

	}
	return errors.New("no such parameter found")

}

func (u *User) DeleteStaffById(userId int) error {
	if !u.IsAdmin {
		return errors.New("unauthorized Access")
	}

	for index, userObject := range allStaff {
		if userObject.UserId == userId {
			allStaff = slices.Delete(allStaff, index, index+1)
			return nil
		}
	}

	return errors.New("User Not Found")

}

// Validations
func validateFirstName(firstName string) error {
	if len(firstName) < 3 {
		return errors.New("firstName should be atleast 3 characters Long")
	}
	return nil
}

func validateLastName(lastName string) error {
	if len(lastName) == 0 {
		return errors.New("lastName Cannot be Empty")
	}
	return nil
}

///update Functions

func (u *User) updateFirstName(value interface{}) error {
	tempFirstNameVariable, stringValidation := value.(string)
	if !stringValidation {
		return errors.New("first name should be string")
	}
	err := validateFirstName(tempFirstNameVariable)
	if err != nil {
		return err
	}
	u.FirstName = tempFirstNameVariable

	return nil
}

func (u *User) updateLastName(value interface{}) error {
	tempLastNameVariable, stringValidation := value.(string)
	if !stringValidation {
		return errors.New("last name should be string")
	}
	err := validateFirstName(tempLastNameVariable)
	if err != nil {
		return err
	}
	u.LastName = tempLastNameVariable

	return nil
}
