package controllers

import (
	_entities "be9/restmvc/entities"
	_helper "be9/restmvc/helper"
	_repositories "be9/restmvc/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	authData := _entities.AuthRequestData{}
	c.Bind(&authData)

	token, e := _repositories.LoginUsers(authData)
	if e != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email or password incorrect"))
	}
	data := map[string]interface{}{
		"token": token,
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccessWithData("login success", data))
}
