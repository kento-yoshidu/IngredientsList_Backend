package repository

import (
	"fmt"
	"ingredients-list/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IIngredientRepository interface {
	GetAllIngredients(ingredients *[]model.Ingredient, dishId uint) error
	GetIngredientByDishId(ingredient *model.Ingredient, dishId, ingredientId uint) error
	CreateIngredient(ingredient *model.Ingredient) error
	UpdateIngredient(ingredient *model.Ingredient, dishId, ingredientId uint) error
	DeleteIngredient(dishId, ingredientId uint) error
}

type ingredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) IIngredientRepository {
	return &ingredientRepository{db}
}

func (ir *ingredientRepository) GetAllIngredients(ingredients *[]model.Ingredient, dishId uint) error {
	if err := ir.db.Joins("Dish").Where("dish_id=?", dishId).Order("created_at").Find(ingredients).Error; err != nil {
		return err
	}

	return nil
}

func (ir *ingredientRepository) GetIngredientByDishId(ingredient *model.Ingredient, dishId, ingredientId uint) error {
	if err := ir.db.Joins("Dish").Where("dish_id=?", dishId).First(ingredient, ingredientId).Error; err != nil {
		return err
	}

	return nil
}

func (ir *ingredientRepository) CreateIngredient(ingredient *model.Ingredient) error {
	if err := ir.db.Create(ingredient).Error; err != nil {
		return err
	}

	return nil
}

func (ir *ingredientRepository) UpdateIngredient(ingredient *model.Ingredient, dishId, ingredientId uint) error {
	result := ir.db.Model(ingredient).Clauses(clause.Returning{}).Where("id=? AND dish_id=?", ingredientId, dishId).Update("ingredientname", ingredient.Ingredientname)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("更新するべきレコードがありませんでした")
	}

	return nil
}

func (ir *ingredientRepository) DeleteIngredient(dishId, ingredientId uint) error {
	result := ir.db.Where("id=? AND dish_id=?", ingredientId, dishId).Delete(&model.Ingredient{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("削除するべきレコードがありませんでした")
	}

	return nil
}