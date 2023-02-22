package shared

import (
	"fmt"

	"encore.dev/beta/errs"
	"github.com/go-playground/validator/v10"
)

type apiValidator struct {
	Validator *validator.Validate
}

func NewAPIValidator(v *validator.Validate) *apiValidator {
	return &apiValidator{
		Validator: v,
	}
}

func (v *apiValidator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

func (v *apiValidator) ParseValidatorError(err error) error {
	message := ""
	space := ""
	for index, err := range err.(validator.ValidationErrors) {
		if index > 0 {
			space = ", "
		}
		message = message + fmt.Sprintf("%v%v is %v", space, err.Field(), err.Tag())
	}
	return &errs.Error{
		Code:    errs.InvalidArgument,
		Message: message,
	}
}
