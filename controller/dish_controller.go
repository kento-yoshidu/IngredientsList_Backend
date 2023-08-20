package controller

import (
	"ingredients-list/model"
	"ingredients-list/usecase"
	"net/http"
	"strconv"

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

func (dc *dishController) GetDishById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("dishId")

	dishId, _ := strconv.Atoi(id)
	dishRes, err := dc.du.GetDishById(uint(userId.(float64)), uint(dishId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dishRes)
}

func (dc *dishController) CreateDish(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	dish := model.Dish{}
	if err := c.Bind(&dish); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	dish.UserId = uint(userId.(float64))

	dishRes, err := dc.du.CreateDish(dish)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, dishRes)
}

func (dc *dishController) DeleteDish(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("dishId")
	dishId, _ := strconv.Atoi(id)

	err := dc.du.DeleteDish(uint(userId.(float64)), uint(dishId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
