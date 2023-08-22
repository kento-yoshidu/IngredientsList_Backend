package validator

import (
	"ingredients-list/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IIngredientValidator interface {
	IngredientValidate(dish model.Ingredient) error
}

type ingredientValidator struct{}

func NewIngredientValidator() IIngredientValidator {
	return &ingredientValidator{}
}

func (iv *ingredientValidator) IngredientValidate(ingredient model.Ingredient) error {
	return validation.ValidateStruct(&ingredient,
		validation.Field(
			&ingredient.Ingredientname,
			validation.Required.Error("Ingredientname is required."),
		),
	)
}
