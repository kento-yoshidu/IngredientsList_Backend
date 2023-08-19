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

	userRepository := repository.NewUserRepository(db)
	dishRepository := repository.NewDishRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	dishUsecase := usecase.NewDishUsecase(dishRepository, dishValidator)

	userController := controller.NewUserController(userUsecase)
	dishController := controller.NewDishController(dishUsecase)

	e := router.NewRouter(userController, dishController)

	e.Logger.Fatal(e.Start(":8080"))
}
