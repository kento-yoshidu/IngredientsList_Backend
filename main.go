package main

import (
	"ingredients-list/controller"
	"ingredients-list/db"
	"ingredients-list/repository"
	"ingredients-list/router"
	"ingredients-list/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)

	e.Logger.Fatal(e.Start(":8080"))
}
