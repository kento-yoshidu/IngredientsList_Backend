package main

import (
	"fmt"
	"ingredients-list/db"
	"ingredients-list/model"
)

func main() {
	dbConnect := db.NewDB()
	defer fmt.Println("Successfully Migrate!")
	defer db.CloseDB(dbConnect)
	dbConnect.AutoMigrate(&model.User{}, &model.Dish{}, &model.Ingredient{})
}
