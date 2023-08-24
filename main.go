package main

import (
	"ingredients-list/controller"
	"ingredients-list/db"
	"ingredients-list/repository"
	"ingredients-list/router"
	"ingredients-list/usecase"
	"ingredients-list/validator"
)

func main() {
	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	dishValidator := validator.NewDishValidator()
	ingredientValidator := validator.NewIngredientValidator()

	userRepository := repository.NewUserRepository(db)
	dishRepository := repository.NewDishRepository(db)
	ingredientRepository := repository.NewIngredientRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	dishUsecase := usecase.NewDishUsecase(dishRepository, dishValidator)
	ingredientUsecase := usecase.NewIngredientUsecase(ingredientRepository, ingredientValidator)

	userController := controller.NewUserController(userUsecase)
	dishController := controller.NewDishController(dishUsecase)
	ingredientController := controller.NewIngredientController(ingredientUsecase)

	e := router.NewRouter(userController, dishController, ingredientController)

	e.Logger.Fatal(e.Start(":8080"))
}
