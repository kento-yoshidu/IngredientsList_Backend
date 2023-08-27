package router

import (
	"fmt"
	"ingredients-list/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

func NewRouter(uc controller.IUserController, dc controller.IDishController, ic controller.IIngredientController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.BodyDump(bodyDumpHandler))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		//CookieMaxAge:   60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	dish := e.Group("/dish")
	dish.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	dish.POST("", dc.CreateDish)
	dish.GET("/:dishId", dc.GetDishById)
	dish.PUT("/:dishId", dc.UpdateDish)
	dish.DELETE("/:dishId", dc.DeleteDish)

	dish.GET("/:dishId/ingredients", ic.GetIngredientsByDishId)
	dish.POST("/:dishId/ingredients", ic.CreateIngredient)
	dish.DELETE("/:dishId/ingredient/:ingredientId", ic.DeleteIngredient)

	dishes := e.Group("/dishes")
	dishes.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	dishes.GET("", dc.GetAllDishes)

	i := e.Group("/ingredient")
	i.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	i.PUT("/:ingredientId", ic.UpdateIngredient)

	return e
}
