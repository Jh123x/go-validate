package govalidate

import (
	"fmt"
)

// ValidateError is an error message used for validation.
type ValidateError struct {
	checkName string
	errMsg    string
}

var _ error = (*ValidateError)(nil)

// NewValidateError returns a new ValidateError with the given field name and error message.
func NewValidateError(fieldName, errMsg string) ValidateError {
	return ValidateError{fieldName, errMsg}
}

// Error returns the error message.
func (v ValidateError) Error() string {
	return fmt.Sprintf(ErrorFormat, v.checkName, v.errMsg)
}

// ValidateErrGenerator returns a ValidateError with the given field name and error message.
func ValidateErrGenerator(checkName, errMsg string) ValidateError {
	return NewValidateError(checkName, errMsg)
}
