package model

type Shouldbuy struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Ingredientname string `json:"ingredientname" gorm:"not null"`
	Dishname       string `json:"dishname"`
}

type ShouldbuyResponse struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Ingredientname string `json:"ingredientname" gorm:"not null"`
	Dishname       string `json:"dishname"`
}
