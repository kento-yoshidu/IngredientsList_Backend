package usecase

import (
	"ingredients-list/model"
	"ingredients-list/repository"
)

type IIngredientUsecase interface {
	GetIngredientsByDishId(userId, dishId uint) ([]model.IngredientResponse, error)
	CreateIngredient(ingredient model.Ingredient) (model.IngredientResponse, error)
	/*
		GetAllIngredients(dishId uint) ([]model.IngredientResponse, error)
		UpdateIngredient(dishId, ingredientId uint) (model.IngredientResponse, error)
		DeleteIngredient(dishId, ingredientId uint) error
	*/
}

type ingredientUsecase struct {
	ir repository.IIngredientRepository
}

func NewIngredientUsecase(ir repository.IIngredientRepository) IIngredientUsecase {
	return &ingredientUsecase{ir}
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
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
		}
		resIngredients = append(resIngredients, i)
	}

	return resIngredients, nil

}

func (iu *ingredientUsecase) CreateIngredient(ingredient model.Ingredient) (model.IngredientResponse, error) {
	if err := iu.ir.CreateIngredient(&ingredient); err != nil {
		return model.IngredientResponse{}, err
	}

	resIngredient := model.IngredientResponse{
		ID:             ingredient.ID,
		Ingredientname: ingredient.Ingredientname,
		CreatedAt:      ingredient.CreatedAt,
		UpdatedAt:      ingredient.UpdatedAt,
	}

	return resIngredient, nil
}

/*
func (iu *ingredientUsecase) GetAllIngredients(dishId uint) ([]model.IngredientResponse, error) {
	ingredients := []model.Ingredient{}
	if err := iu.ir.GetAllIngredients(&ingredients, dishId); err != nil {
		return nil, err
	}

	resIngredients := []model.IngredientResponse{}
	for _, v := range ingredients {
		ingredient := model.IngredientResponse{
			ID:             v.ID,
			Ingredientname: v.Ingredientname,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
		}
		resIngredients = append(resIngredients, ingredient)
	}

	return resIngredients, nil
}
*/
