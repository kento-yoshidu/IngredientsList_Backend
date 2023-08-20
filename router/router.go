package router

import (
	"ingredients-list/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, dc controller.IDishController) *echo.Echo {
	e := echo.New()

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	d := e.Group("/dishes")
	d.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	d.GET("", dc.GetAllDishes)
	d.GET("/:dishId", dc.GetDishById)
	d.POST("", dc.CreateDish)
	d.DELETE("/:dishId", dc.DeleteDish)

	return e
}
