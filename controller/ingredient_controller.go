package controller

import (
	"fmt"
	"ingredients-list/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IIngredientController interface {
	GetAllIngredients(c echo.Context) error
	// CreateIngredient(c echo.Context) error
}

type ingredientController struct {
	iu usecase.IIngredientUsecase
}

func NewIngredientController(iu usecase.IIngredientUsecase) IIngredientController {
	return &ingredientController{iu}
}

func (ic *ingredientController) GetAllIngredients(c echo.Context) error {

	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)

	fmt.Printf("path is = %T\n", c.Param("dishId"))

	dishId, _ := strconv.Atoi(c.Param("dishId"))

	ingredientsRes, err := ic.iu.GetAllIngredients(uint(dishId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ingredientsRes)
}

/*
func (ic *ingredientController) CreateIngredient(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	dishId := claims["dish_id"]

	ingredientsRes, err := ic.iu.GetAllIngredients(uint(dishId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ingredientsRes)

}
*/
