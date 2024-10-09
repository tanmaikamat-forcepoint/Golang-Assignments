package contact

import (
	"contactApp/contactInfo"
	"contactApp/helper"
	"errors"
)

type Contact struct {
	ContactId      int
	FirstName      string
	LastName       string
	IsActive       bool
	ContactDetails []*contactInfo.ContactInfo
}

//CRUD Contact

func NewContact(
	contactId int,
	firstName string,
	lastName string) (*Contact, error) {
	err := helper.ValidateAll(
		validateFirstName(firstName),
		validateLastName(lastName))
	if err != nil {
		return nil, err
	}
	var tempContactInfoList []*contactInfo.ContactInfo

	tempContactObject := &Contact{
		ContactId:      contactId,
		FirstName:      firstName,
		LastName:       lastName,
		IsActive:       true,
		ContactDetails: tempContactInfoList,
	}

	return tempContactObject, nil
}

func (c *Contact) UpdateContact(paramenter string, value interface{}) error {
	err := c.validateIfActive()
	if err != nil {
		return err
	}

	switch paramenter {
	case "firstName":
		return c.updateFirstName(value)
	case "lastName":
		return c.updateLastName(value)
	}

	return errors.New("no parameter found")

}

func (c *Contact) DeleteContact() error {
	err := c.validateIfActive()
	if err != nil {
		return err
	}
	c.IsActive = false
	return nil
}

//CRUD for ContactInfo

func (c *Contact) NewContactInfo(contactInfoType string, contactInfoValue string) (*contactInfo.ContactInfo, error) {
	err := c.validateIfActive()
	if err != nil {
		return nil, err
	}
	newContactInfoId := c.getNextIdFromPrevContactInfo()
	tempContactInfo, err1 := contactInfo.NewContactInfo(newContactInfoId, contactInfoType, contactInfoValue)
	if err1 != nil {
		return nil, err1
	}
	c.ContactDetails = append(c.ContactDetails, tempContactInfo)
	return tempContactInfo, nil
}

func (c *Contact) GetAllContactInfo(contactDetails *[]*contactInfo.ContactInfo) error {
	err := c.validateIfActive()
	if err != nil {
		return err
	}
	for _, detail := range c.ContactDetails {
		if !detail.IsActive {
			continue
		}
		*contactDetails = append(*contactDetails, detail)
	}
	return nil
}

func (c *Contact) GetContactInfoById(contactInfoId int) (*contactInfo.ContactInfo, error) {
	err := c.validateIfActive()
	if err != nil {
		return nil, err
	}
	for _, detail := range c.ContactDetails {
		if !detail.IsActive {
			continue
		}
		if detail.ContactInfoId == contactInfoId {
			return detail, nil
		}
	}

	return nil, errors.New("Contact Info Not found")
}

func (c *Contact) UpdateContactInfo(contactInfoId int, parameter string, value interface{}) error {
	err := c.validateIfActive()
	if err != nil {
		return err
	}
	tempContactInfo, err1 := c.GetContactInfoById(contactInfoId)
	if err1 != nil {
		return err1
	}
	return tempContactInfo.UpdateContactInfo(parameter, value)
}

func (c *Contact) DeleteContactInfo(contactInfoId int) error {
	err := c.validateIfActive()
	if err != nil {
		return err
	}
	tempContactInfo, err1 := c.GetContactInfoById(contactInfoId)
	if err1 != nil {
		return err1
	}
	return tempContactInfo.DeleteContactInfo()
}

// helpers
func (c *Contact) getNextIdFromPrevContactInfo() int {
	if len(c.ContactDetails) == 0 {
		return 1
	}
	return c.ContactDetails[len(c.ContactDetails)-1].ContactInfoId + 1
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

func (c *Contact) validateIfActive() error {
	if !c.IsActive {
		return errors.New("contact do not exist")
	}
	return nil
}

//

///update Functions

func (c *Contact) updateFirstName(value interface{}) error {
	tempFirstNameVariable, stringValidation := value.(string)
	if !stringValidation {
		return errors.New("first name should be string")
	}
	err := validateFirstName(tempFirstNameVariable)
	if err != nil {
		return err
	}
	c.FirstName = tempFirstNameVariable

	return nil
}

func (c *Contact) updateLastName(value interface{}) error {
	tempLastNameVariable, stringValidation := value.(string)
	if !stringValidation {
		return errors.New("last name should be string")
	}
	err := validateFirstName(tempLastNameVariable)
	if err != nil {
		return err
	}
	c.LastName = tempLastNameVariable

	return nil
}
