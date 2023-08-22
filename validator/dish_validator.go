package validator

import (
	"ingredients-list/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IDishValidator interface {
	DishValidate(dish model.Dish) error
}

type dishValidator struct{}

func NewDishValidator() IDishValidator {
	return &dishValidator{}
}

func (dv *dishValidator) DishValidate(dish model.Dish) error {
	return validation.ValidateStruct(&dish,
		validation.Field(
			&dish.Dishname,
			validation.Required.Error("DishName is required."),
		),
	)
}
