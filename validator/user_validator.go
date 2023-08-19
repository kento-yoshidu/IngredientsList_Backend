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
			&user.UserName,
			validation.Required.Error("UserName is required."),
			is.Alphanumeric.Error("UserName is alpha or num only"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("Password is required."),
		),
	)
}
