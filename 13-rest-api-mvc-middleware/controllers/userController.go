package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	_entities "be9/restmvc/entities"
	_helper "be9/restmvc/helper"
	_middlewares "be9/restmvc/middlewares"
	_repositories "be9/restmvc/repositories"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {

	users, err := _repositories.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed get users",
		})
	}

	responseData := []_entities.UserResponseData{}
	for _, value := range users {
		var response _entities.UserResponseData
		response.ID = value.ID
		response.Name = value.Name
		response.Email = value.Email
		response.CreatedAt = value.CreatedAt

		responseData = append(responseData, response)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    responseData,
	})
}

func GetUserByIdController(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	id := c.Param("id")
	iduser, errid := strconv.Atoi(id)
	if errid != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("id not recognise"))
	}

	if idToken != iduser {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("unauthorized"))
	}

	users, err := _repositories.GetUserById(iduser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed get user data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {
	user := _entities.UserRequestData{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("error", err)

		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed create user bind"))
	}

	result, errCreate := _repositories.CreateUser(user)

	if errCreate != nil {

		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert data"))
	}
	if result == 0 {
		// return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 	"status":  "error",
		// 	"message": "failed to insert data",
		// })
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert data"))
	}
	// To change its value, you could use `Update`
	// DB.Model(&user).Update("created_at", time.Now().In(gmt))

	return c.JSON(http.StatusOK, _helper.ResponseSuccessNoData("success insert data"))

}
