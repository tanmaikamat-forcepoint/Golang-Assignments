package validations

import (
	"errors"
	"strings"
)

type Number interface {
	int | int64 | float64
}

func ValidateIfNotEmpty(key string, object string) error {
	object = strings.TrimSpace(object)
	if len(object) == 0 {
		return errors.New(key + " is Empty. Please input Valid String")
	}
	return nil
}

func ValidateIfNotNegative(key string, object int) error {

	if object < 0 {
		return errors.New(key + " Cannot be negative")
	}
	return nil
}

func ValidateIfNotZero(key string, object int) error {

	if object <= 0 {
		return errors.New(key + " Cannot be Zero")
	}
	return nil
}
func ValidateIfNotNegativeF(key string, object float32) error {

	if object < 0 {
		return errors.New(key + " Cannot be Negative")
	}
	return nil
}
