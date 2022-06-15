package repositories

import (
	_config "be9/restmvc/config"
	_entities "be9/restmvc/entities"
	_middlewares "be9/restmvc/middlewares"
	_models "be9/restmvc/models"
	"fmt"
)

func LoginUsers(auth _entities.AuthRequestData) (interface{}, error) {
	var userData _models.User
	// var err error
	result := _config.DB.Where("email = ? AND password = ?", auth.Email, auth.Password).First(&userData)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println(result.RowsAffected)
	if result.RowsAffected != 1 {
		return nil, fmt.Errorf("error")
	}

	token, errToken := _middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return nil, errToken
	}

	return token, nil
}
