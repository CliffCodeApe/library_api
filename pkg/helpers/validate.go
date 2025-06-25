package helpers

import (
	"library_api/pkg/errs"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateStruct(payload interface{}) error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(payload)
	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}
