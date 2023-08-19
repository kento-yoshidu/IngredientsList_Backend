package controller

import (
	"ingredients-list/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IDishController interface {
	GetAllDishes(c echo.Context) error
	GetDishById(c echo.Context) error
	CreateDish(c echo.Context) error
	DeleteDish(c echo.Context) error
}

type dishController struct {
	du usecase.IDishUsecase
}

func NewDishController(du usecase.IDishUsecase) IDishController {
	return &dishController{du}
}

func (dc *dishController) GetAllDishes(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	dishesRes, err := dc.du.GetAllDishes(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dishesRes)
}
