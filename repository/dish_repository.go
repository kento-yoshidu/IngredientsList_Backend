package repository

import (
	"fmt"
	"ingredients-list/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IDishRepository interface {
	GetAllDishes(dishes *[]model.Dish, userId uint) error
	GetDishById(dish *model.Dish, userId, dishId uint) error
	CreateDish(dish *model.Dish) error
	UpdateDish(dish *model.Dish, userId, dishId uint) error
	DeleteDish(userId, dishId uint) error
}

type dishRepository struct {
	db *gorm.DB
}

func NewDishRepository(db *gorm.DB) IDishRepository {
	return &dishRepository{db}
}

func (dr *dishRepository) GetAllDishes(dishes *[]model.Dish, userId uint) error {
	if err := dr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(dishes).Error; err != nil {
		return err
	}

	return nil
}

func (dr *dishRepository) GetDishById(dish *model.Dish, userId, dishId uint) error {
	if err := dr.db.Joins("User").Where("user_id=?", userId).First(dish, dishId).Error; err != nil {
		return err
	}

	return nil
}

func (dr *dishRepository) CreateDish(dish *model.Dish) error {
	if err := dr.db.Create(dish).Error; err != nil {
		return err
	}

	return nil
}

func (dr *dishRepository) UpdateDish(dish *model.Dish, userId, dishId uint) error {
	result := dr.db.Model(dish).Clauses(clause.Returning{}).Where("id=? AND user_id=?", dishId, userId).Update("dishname", dish.Dishname)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}

func (dr *dishRepository) DeleteDish(userId, dishId uint) error {
	result := dr.db.Where("id=? AND user_id=?", dishId, userId).Delete(&model.Dish{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}
