package usecase

import (
	"fmt"
	"ingredients-list/model"
	"ingredients-list/repository"
	"ingredients-list/validator"
)

type IIngredientUsecase interface {
	GetIngredientsByDishId(userId, dishId uint) ([]model.IngredientResponse, error)
	CreateIngredient(ingredient model.Ingredient) (model.IngredientResponse, error)
	UpdateIngredient(ingredient model.Ingredient, ingredientId uint) (model.IngredientResponse, error)
	DeleteIngredient(dishId, ingredientId uint) error
}

type ingredientUsecase struct {
	ir repository.IIngredientRepository
	iv validator.IIngredientValidator
}

func NewIngredientUsecase(ir repository.IIngredientRepository, iv validator.IIngredientValidator) IIngredientUsecase {
	return &ingredientUsecase{ir, iv}
}

func (iu *ingredientUsecase) GetIngredientsByDishId(userId, dishId uint) ([]model.IngredientResponse, error) {
	ingredients := []model.Ingredient{}
	if err := iu.ir.GetIngredientsByDishId(&ingredients, userId, dishId); err != nil {
		return nil, err
	}

	resIngredients := []model.IngredientResponse{}
	for _, v := range ingredients {
		i := model.IngredientResponse{
			ID:             v.ID,
			Ingredientname: v.Ingredientname,
			Shouldbuy:      v.Shouldbuy,
			Dishname:       v.Dish.Dishname,
		}
		resIngredients = append(resIngredients, i)
	}

	return resIngredients, nil

}

func (iu *ingredientUsecase) CreateIngredient(ingredient model.Ingredient) (model.IngredientResponse, error) {
	if err := iu.iv.IngredientValidate(ingredient); err != nil {
		return model.IngredientResponse{}, err
	}

	if err := iu.ir.CreateIngredient(&ingredient); err != nil {
		return model.IngredientResponse{}, err
	}

	resIngredient := model.IngredientResponse{
		ID:             ingredient.ID,
		Ingredientname: ingredient.Ingredientname,
		Shouldbuy:      ingredient.Shouldbuy,
	}

	return resIngredient, nil
}

func (iu *ingredientUsecase) UpdateIngredient(ingredient model.Ingredient, ingredientId uint) (model.IngredientResponse, error) {
	fmt.Printf("%v", ingredient)
	if err := iu.ir.UpdateIngredient(&ingredient, ingredientId); err != nil {
		return model.IngredientResponse{}, err
	}

	resIngredient := model.IngredientResponse{
		ID:             ingredient.ID,
		Ingredientname: ingredient.Ingredientname,
		Shouldbuy:      ingredient.Shouldbuy,
	}

	return resIngredient, nil
}

func (iu *ingredientUsecase) DeleteIngredient(dishId, ingredientId uint) error {
	if err := iu.ir.DeleteIngredient(dishId, ingredientId); err != nil {
		return err
	}

	return nil
}
