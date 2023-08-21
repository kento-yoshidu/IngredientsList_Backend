package validator

import (
	"ingredients-list/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Username,
			validation.Required.Error("Username is required."),
			is.Alphanumeric.Error("Username is alpha or num only"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("Password is required."),
		),
	)
}
