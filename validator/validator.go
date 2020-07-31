package validator

import (
	"fmt"
	"strings"
)

// Validator : struct field validations
type Validator struct {
	Errors []error
}

// Validate :
func (v *Validator) Validate(cond bool, msg string, args ...interface{}) {
	if cond {
		v.Errors = append(v.Errors, fmt.Errorf(msg, args...))
	}
}

// ValidatePresenceOf : validates presence of struct string field
func (v *Validator) ValidatePresenceOf(fieldName string, fieldValue string) {
	cond := len(strings.TrimSpace(fieldValue)) == 0
	v.Validate(cond, "%s cannot be blank", fieldName)
}

// ValidateMaxLengthOf : validates maximum character length of struct string field
func (v *Validator) ValidateMaxLengthOf(fieldName string, fieldValue string, max int) {
	cond := len(fieldValue) > max
	v.Validate(cond, "%s cannot be greater than %d characters", fieldName, max)
}

// ValidateMinLengthOf : validates minimum character length of struct string field
func (v *Validator) ValidateMinLengthOf(fieldName string, fieldValue string, min int) {
	cond := len(fieldValue) < min
	v.Validate(cond, "%s must be at least %d characters", fieldName, min)
}
