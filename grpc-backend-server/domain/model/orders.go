package model

import (
	"time"

	"git.paylabo.com/c002/harp/backend-purchase/domain/validator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Order struct {
	ID          int64 `gorm:"primary_key"`
	name        string
	RequestDate time.Time `gorm:"-"`
}

type OrderValidationRules struct {
	ID   []validation.Rule
	name []validation.Rule
}

func GetOrderValidationRules() OrderValidationRules {
	return OrderValidationRules{
		ID: []validation.Rule{
			validation.Required.Error(validator.RequiredMessage),
			is.Int.Error(validator.IntMessage)},
	}
}
