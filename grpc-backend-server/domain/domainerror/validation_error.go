package domainerror

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type ValidationError struct {
	message string
	fields  map[string][]string
}

type ValidationErrorBuilder struct {
	fields map[string][]string
}

func NewValidationErrorBuilder() *ValidationErrorBuilder {
	return &ValidationErrorBuilder{fields: make(map[string][]string)}
}

func (v *ValidationErrorBuilder) AddMessage(field, message string) *ValidationErrorBuilder {
	v.fields[field] = append(v.fields[field], message)
	return v
}

func (v *ValidationErrorBuilder) Build() error {
	errorFields := make(map[string][]string)
	errorMessages := make([]string, 0)
	for key, value := range v.fields {
		errorFields[key] = append(errorFields[key], value...)
		errorMessages = append(errorMessages, fmt.Sprintf("[key=%s, details=%s]", key, strings.Join(value, ", ")))
	}
	return errors.WithStack(&ValidationError{fields: errorFields, message: fmt.Sprintf("validation error: %s", strings.Join(errorMessages, ", "))})
}

func (v *ValidationError) Error() string {
	return v.message
}

func (v *ValidationError) Unwrap() error {
	return v
}

func (v *ValidationError) IsNotFound() bool {
	return false
}

func (v *ValidationError) IsValidationError() bool {
	return true
}

func (v *ValidationError) IsInternalServerError() bool {
	return false
}

func (v ValidationError) GetFields() map[string][]string {
	return v.fields
}
