package validator

import (
	domainerror "git.paylabo.com/c002/harp/backend-purchase/domain/domainerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateStruct(ptr interface{}, rules ...*validation.FieldRules) error {
	err := validation.ValidateStruct(ptr, rules...)
	if err == nil {
		return nil
	}
	builder := domainerror.NewValidationErrorBuilder()
	if validationErrors, ok := err.(validation.Errors); ok {
		for key, value := range validationErrors {
			builder.AddMessage(key, value.Error())
		}
	} else {
		return domainerror.NewInternalServerError("unknown validation error", err)
	}
	return builder.Build()
}
