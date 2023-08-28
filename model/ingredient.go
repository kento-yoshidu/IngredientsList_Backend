package model

import "time"

type Ingredient struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Ingredientname string    `json:"ingredientname" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Shouldbuy      bool      `json:"shouldbuy" gorm:"not null"`
	DishId         uint      `json:"dish_id" gorm:"not null"`
	Dish           Dish      `json:"dish" gorm:"foreignKey:DishId; constraint:OnDelete:CASCADE"`
}

type IngredientResponse struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Ingredientname string `json:"ingredientname" gorm:"not null"`
	Shouldbuy      bool   `json:"shouldbuy" gorm:"not null" gorm:"default:true"`
	Dishname       string `json:"dishname"`
}
