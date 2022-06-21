package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", HelloController)
	e.GET("/guest", HelloGuestController)
	e.Start(":80")
}

func HelloController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello world",
	})
}

func HelloGuestController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello GUEST",
	})
}
