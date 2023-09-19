package repository

import (
	"fmt"
	"ingredients-list/model"

	"gorm.io/gorm"
)

type IIngredientRepository interface {
	GetIngredientsByDishId(ingredient *[]model.Ingredient, userId, dishId uint) error
	CreateIngredient(ingredient *model.Ingredient) error
	UpdateIngredient(ingredient *model.Ingredient, ingredientId uint) error
	DeleteIngredient(dishId, ingredientId uint) error
	GetShouldBuyIngredients(ingredient *[]model.Ingredient, userId uint) error
}

type ingredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(db *gorm.DB) IIngredientRepository {
	return &ingredientRepository{db}
}

func (ir *ingredientRepository) GetIngredientsByDishId(ingredient *[]model.Ingredient, userId, dishId uint) error {
	if err := ir.db.Order("created_at").Joins("Dish").Where("user_id=? AND dish_id=?", userId, dishId).Find(ingredient).Error; err != nil {
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
	result := ir.db.Model(ingredient).Select("ingredientname", "shouldbuy").Where("id=?", ingredientId).Updates(model.Ingredient{ID: ingredientId, Ingredientname: ingredient.Ingredientname, Shouldbuy: ingredient.Shouldbuy})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("更新するべきレコードがありませんでした")
	}

	return nil
}

func (ir *ingredientRepository) DeleteIngredient(dishId, ingredientId uint) error {
	result := ir.db.Debug().Joins("Dish").Where("id=? AND dish_id=?", ingredientId, dishId).Delete(&model.Ingredient{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("削除するべきレコードがありませんでした")
	}

	return nil
}

func (ir *ingredientRepository) GetShouldBuyIngredients(ingredient *[]model.Ingredient, userId uint) error {
	if err := ir.db.Table("ingredients").Joins("JOIN dishes on ingredients.dish_id = dishes.id").Joins("JOIN users on dishes.user_id = users.id").Where("user_id=? AND shouldbuy = true", userId).Find(ingredient).Error; err != nil {
		return err
	}

	return nil
}
