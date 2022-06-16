package routes

import (
	"be9/restclean/factory"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	presenter := factory.InitFactory()
	e := echo.New()

	e.GET("/users", presenter.UserPresenter.GetAll)
	// e.GET("/products", presenter.ProductPresenter.GetAll)
	// e.GET("/users/:id", _controllers.GetUserByIdController, middlewares.JWTMiddleware())
	// e.POST("/users", _controllers.CreateUserController)
	// e.POST("/auth", _controllers.LoginUserController)

	return e
}
