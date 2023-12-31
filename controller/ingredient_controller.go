package controller

import (
	"ingredients-list/model"
	"ingredients-list/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IIngredientController interface {
	GetIngredientsByDishId(c echo.Context) error
	CreateIngredient(c echo.Context) error
	UpdateIngredient(c echo.Context) error
	DeleteIngredient(c echo.Context) error
	GetShouldBuyIngredients(c echo.Context) error
}

type ingredientController struct {
	iu usecase.IIngredientUsecase
}

func NewIngredientController(iu usecase.IIngredientUsecase) IIngredientController {
	return &ingredientController{iu}
}

func (ic *ingredientController) GetIngredientsByDishId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("dishId")

	dishId, _ := strconv.Atoi(id)
	ingredientsRes, err := ic.iu.GetIngredientsByDishId(uint(userId.(float64)), uint(dishId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ingredientsRes)
}

func (ic *ingredientController) CreateIngredient(c echo.Context) error {
	id := c.Param("dishId")
	dishId, _ := strconv.Atoi(id)

	ingredient := model.Ingredient{}
	if err := c.Bind(&ingredient); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ingredient.DishId = uint(dishId)

	ingredientsRes, err := ic.iu.CreateIngredient(ingredient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ingredientsRes)
}

func (ic *ingredientController) UpdateIngredient(c echo.Context) error {
	id := c.Param("ingredientId")
	ingredientId, _ := strconv.Atoi(id)

	ingredient := model.Ingredient{}
	if err := c.Bind(&ingredient); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ingredientRes, err := ic.iu.UpdateIngredient(ingredient, uint(ingredientId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ingredientRes)
}

func (ic *ingredientController) DeleteIngredient(c echo.Context) error {
	id := c.Param("dishId")
	dishId, _ := strconv.Atoi(id)

	ingreid := c.Param("ingredientId")
	ingredientId, _ := strconv.Atoi(ingreid)

	err := ic.iu.DeleteIngredient(uint(dishId), uint(ingredientId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (ic *ingredientController) GetShouldBuyIngredients(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	ingredientsRes, err := ic.iu.GetShouldBuyIngredients(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ingredientsRes)
}
