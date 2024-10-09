package user

import (
	"contactApp/contact"
	"contactApp/contactInfo"
	"contactApp/helper"
	"errors"
)

type User struct {
	UserId    int
	FirstName string
	LastName  string
	IsActive  bool
	IsAdmin   bool
	Contacts  []*contact.Contact
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
		Contacts:  nil,
	}
	userId++
	allAdmins = append(allAdmins, tempNewAdminObject)
	return tempNewAdminObject, nil

}

func (u *User) NewStaff(firstName string, lastName string) (*User, error) {

	err := helper.ValidateAll(
		validateAdminRights(u),
		validateFirstName(firstName),
		validateLastName(lastName))
	if err != nil {
		return nil, err
	}
	var tempNewContactArray []*contact.Contact
	tempNewStaffObject := &User{
		UserId:    userId,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
		IsAdmin:   false,
		Contacts:  tempNewContactArray,
	}
	userId++
	allStaff = append(allAdmins, tempNewStaffObject)
	return tempNewStaffObject, nil

}

func (u *User) GetAllStaff() ([]*User, error) {
	err := validateAdminRights(u)
	if err != nil {
		return nil, err
	}
	var tempStaff []*User
	for _, userObject := range allStaff {
		if !userObject.IsActive {
			continue
		}
		tempStaff = append(tempStaff, userObject)
	}
	return tempStaff, nil
}

func (u *User) GetStaffByID(userId int) (*User, error) {
	err := validateAdminRights(u)
	if err != nil {
		return nil, err
	}
	for _, userObject := range allStaff {
		if !userObject.IsActive {
			continue
		}
		if userObject.UserId == userId {
			return userObject, nil
		}
	}

	return nil, errors.New("User Not Found")
}
func (u *User) UpdateStaffByID(userId int, parameter string, value interface{}) error {
	err := validateAdminRights(u)
	if err != nil {
		return err
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
	err := validateAdminRights(u)
	if err != nil {
		return err
	}
	for _, userObject := range allStaff {
		if userObject.UserId == userId {
			userObject.IsActive = false
			return nil
		}
	}

	return errors.New("User Not Found")

}

//CRUD for Contacts

func (u *User) NewContact(
	firstName string,
	lastName string) (*contact.Contact, error) {
	err := validateStaffRights(u)
	if err != nil {
		return nil, err
	}
	nextContactId := u.getNextIdFromPrevContactList()
	tempContactObject, err := contact.NewContact(nextContactId, firstName, lastName)
	if err != nil {
		return nil, err
	}
	u.Contacts = append(u.Contacts, tempContactObject)
	return tempContactObject, nil
}

func (u *User) GetAllContacts() ([]*contact.Contact, error) {

	err := validateStaffRights(u)
	if err != nil {
		return nil, err
	}
	var tempContactList []*contact.Contact
	for _, contactValue := range u.Contacts {
		if !contactValue.IsActive {
			continue
		}
		tempContactList = append(tempContactList, contactValue)
	}
	return tempContactList, nil
}

func (u *User) GetContactById(contactId int) (*contact.Contact, error) {

	err := validateStaffRights(u)
	if err != nil {
		return nil, err
	}
	for _, contactValue := range u.Contacts {
		if !contactValue.IsActive {
			continue
		}
		if contactValue.ContactId == contactId {
			return contactValue, nil
		}
	}
	return nil, errors.New("no contact found")
}

func (u *User) UpdateContact(contactId int, parameter string, value interface{}) error {
	err := validateStaffRights(u)
	if err != nil {
		return err
	}
	tempContactObject, err1 := u.GetContactById(contactId)
	if err1 != nil {
		return err1
	}
	err2 := tempContactObject.UpdateContact(parameter, value)
	if err2 != nil {
		return err2
	}
	return nil
}

func (u *User) DeleteContact(contactId int) error {
	err := validateStaffRights(u)
	if err != nil {
		return err
	}
	tempContactObject, err1 := u.GetContactById(contactId)
	if err1 != nil {
		return err1
	}
	err2 := tempContactObject.DeleteContact()
	if err2 != nil {
		return err2
	}
	return nil
}

// /CRUD for ContactInfo
func (u *User) NewContactInfo(
	contactId int,
	contactInfoType string,
	contactInfoValue string) (*contactInfo.ContactInfo, error) {
	err := validateStaffRights(u)
	if err != nil {
		return nil, err
	}
	tempContactObject, err := u.GetContactById(contactId)
	if err != nil {
		return nil, err
	}
	return tempContactObject.NewContactInfo(contactInfoType, contactInfoValue)
}

func (u *User) GetAllContactInfo(
	contactId int,
	contactDetails *[]*contactInfo.ContactInfo) error {
	err := validateStaffRights(u)
	if err != nil {
		return err
	}
	tempContactObject, err := u.GetContactById(contactId)
	if err != nil {
		return err
	}
	tempContactObject.GetAllContactInfo(contactDetails)
	return nil
}

func (u *User) UpdateContactInfo(
	contactId int,
	contactInfoId int,
	parameter string,
	value interface{}) error {
	err := validateStaffRights(u)
	if err != nil {
		return err
	}
	tempContactObject, err := u.GetContactById(contactId)
	if err != nil {
		return err
	}
	return tempContactObject.UpdateContactInfo(contactInfoId, parameter, value)

}

func (u *User) DeleteContactInfo(
	contactId int, contactInfoId int) error {
	err := validateStaffRights(u)
	if err != nil {
		return err
	}
	tempContactObject, err := u.GetContactById(contactId)
	if err != nil {
		return err
	}
	return tempContactObject.DeleteContactInfo(contactInfoId)

}

// helpers
func (u *User) getNextIdFromPrevContactList() int {
	if u.Contacts == nil {
		panic("Error! Contact List Not Found")
	}
	if len(u.Contacts) == 0 {
		return 1
	}
	return u.Contacts[len(u.Contacts)-1].ContactId + 1

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

func validateAdminRights(user *User) error {
	if !user.IsActive {
		return errors.New("user do not exists")
	}

	if !user.IsAdmin {
		return errors.New("unauthorized Access")
	}

	return nil
}

func validateStaffRights(user *User) error {
	if !user.IsActive {
		return errors.New("user do not exists")
	}

	if user.IsAdmin {
		return errors.New("Admin Cannot Access Contacts")
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
