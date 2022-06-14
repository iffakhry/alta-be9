package main

// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql
// go get -u github.com/labstack/echo/v4

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var DB *gorm.DB

// database connection
func InitDB() {

	// declare struct config & variable connectionString
	connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/db_be9_gorm?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

func init() {
	InitDB()           //create connection
	InitialMigration() //run auto migrate
}

func main() {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":80"))
}

func GetUsersController(c echo.Context) error {
	var users []User
	tx := DB.Find(&users)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "failed to get data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get data",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("error", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "failed create user bind",
		})
	}
	gmt, _ := time.LoadLocation("UTC")
	time.Local = gmt
	user.CreatedAt = time.Now().In(gmt)
	fmt.Println(user.CreatedAt)
	tx := DB.Create(&user)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "failed to insert data",
		})
	}
	if tx.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "failed to insert data",
		})
	}
	// To change its value, you could use `Update`
	DB.Model(&user).Update("created_at", time.Now().In(gmt))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success insert data",
	})

}
