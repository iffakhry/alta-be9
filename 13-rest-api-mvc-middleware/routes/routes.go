package routes

import (
	_controllers "be9/restmvc/controllers"
	"be9/restmvc/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", _controllers.GetUsersController, middlewares.JWTMiddleware())
	e.GET("/users/:id", _controllers.GetUserByIdController)
	e.POST("/users", _controllers.CreateUserController)
	e.POST("/auth", _controllers.LoginUserController)

	return e
}
