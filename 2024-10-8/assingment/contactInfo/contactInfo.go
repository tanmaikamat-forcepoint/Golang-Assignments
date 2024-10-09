package contactInfo

import (
	"contactApp/helper"
	"errors"
	"net/mail"
	"regexp"
)

type ContactInfo struct {
	ContactInfoId    int
	IsActive         bool
	ContactInfoType  string
	ContactInfoValue string
}

func NewContactInfo(
	contactInfoId int,
	contactInfoType string,
	contactInfoValue string) (*ContactInfo, error) {
	err := helper.ValidateAll(
		validateContactInfoType(contactInfoType),
		validateContactInfoValue(contactInfoType, contactInfoValue))
	if err != nil {
		return nil, err
	}
	tempContactInfoObject := &ContactInfo{
		ContactInfoId:    contactInfoId,
		IsActive:         true,
		ContactInfoType:  contactInfoType,
		ContactInfoValue: contactInfoValue,
	}

	return tempContactInfoObject, nil
}

func (cInfo *ContactInfo) UpdateContactInfo(parameter string, value interface{}) error {
	err := validateIfActive(cInfo)
	if err != nil {
		return err
	}
	switch parameter {
	case "contactInfoValue":
		return cInfo.updateContactInfoValue(value)
	}
	return errors.New("n Valid Parameter Found")
}

func (cInfo *ContactInfo) DeleteContactInfo() error {
	err := validateIfActive(cInfo)
	if err != nil {
		return err
	}
	cInfo.IsActive = false
	return nil
}

// validations
func validateIfActive(cinfo *ContactInfo) error {
	if !cinfo.IsActive {
		return errors.New("invalid Contact Info")
	}
	return nil
}

func validateContactInfoType(contactInfoType string) error {
	if contactInfoType == "email" || contactInfoType == "number" {
		return nil
	}
	return errors.New("not a valid info type.Valid Info type are 'email' and 'number'")

}

func validateContactInfoValue(ContactInfoType string, contactInfoValue string) error {
	if ContactInfoType == "email" {
		_, err := mail.ParseAddress(contactInfoValue)
		if err != nil {
			return err
		}
		return nil
	}
	if ContactInfoType == "number" {
		match, err := regexp.MatchString(`(0|\\+91)[0-9]+$`, contactInfoValue)
		if err != nil {
			return err
		}
		if !match {
			return errors.New("not a Valid Phone Number")
		}
		return nil
	}
	//Validation for Phone Number
	return errors.New("invalid InfoType")
}

// update Functions
func (cInfo *ContactInfo) updateContactInfoValue(value interface{}) error {
	tempContactInfoValue, stringValidation := value.(string)
	if !stringValidation {
		return errors.New("invalid Value. Required String")
	}

	err := validateContactInfoValue(cInfo.ContactInfoType, tempContactInfoValue)
	if err != nil {
		return err
	}

	cInfo.ContactInfoValue = tempContactInfoValue
	return nil
}
