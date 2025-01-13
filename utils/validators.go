package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// validates if the value matches the given regex
func ValidateRegex(fl validator.FieldLevel) bool {
	value := fl.Field().String() // value to validate
	regex := fl.Param()          // regex to match

	// match the value with the regex
	matched, err := regexp.MatchString(regex, value)
	if err != nil {
		return false // return false if there is an error
	}

	return matched // return the result of the match
}

// validates if the value has the given precision
func ValidatePrecision(fl validator.FieldLevel) bool {
	value := fl.Field().String() // value to validate
	precision := fl.Param()      // precision to match

	// match the value with the regex for the given precision
	matched, err := regexp.MatchString(`^\d+\.\d{`+precision+`}$`, value)
	if err != nil {
		return false // return false if there is an error
	}

	return matched // return the result of the match
}
