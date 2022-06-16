package data

import (
	"be9/restclean/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) SelectData(data string) (response []users.Core, err error) {
	var dataUsers []User
	result := repo.db.Find(&dataUsers)
	if result.Error != nil {
		return []users.Core{}, result.Error
	}
	return toCoreList(dataUsers), nil
}
