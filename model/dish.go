package model

import "time"

type Dish struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Dishname  string    `json:"dishname" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
}

type DishResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Dishname  string    `json:"dishname" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
