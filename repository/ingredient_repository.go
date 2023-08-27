package repository

import (
	"fmt"
	"ingredients-list/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IIngredientRepository interface {
	GetIngredientsByDishId(ingredient *[]model.Ingredient, userId, dishId uint) error
	CreateIngredient(ingredient *model.Ingredient) error
	UpdateIngredient(ingredient *model.Ingredient, ingredientId uint) error
	DeleteIngredient(ingredientId uint) error
}

type ingredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) IIngredientRepository {
	return &ingredientRepository{db}
}

func (ir *ingredientRepository) GetIngredientsByDishId(ingredient *[]model.Ingredient, userId, dishId uint) error {
	if err := ir.db.Joins("Dish").Where("user_id=? AND dish_id=?", userId, dishId).Find(ingredient).Error; err != nil {
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

func (ir *ingredientRepository) UpdateIngredient(ingredient *model.Ingredient, ingredientId uint) error {
	result := ir.db.Model(ingredient).Clauses(clause.Returning{}).Where("id=?", ingredientId).Update("ingredientname", ingredient.Ingredientname)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("更新するべきレコードがありませんでした")
	}

	return nil
}

func (ir *ingredientRepository) DeleteIngredient(ingredientId uint) error {
	fmt.Printf("id=%v", ingredientId)

	result := ir.db.Joins("User").Joins("Dish").Where("id=?", ingredientId).Delete(&model.Ingredient{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("削除するべきレコードがありませんでした")
	}

	return nil
}
