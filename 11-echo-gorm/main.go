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
	Products []Product
	Address  Address
}

type Jurusan struct {
	// gorm.Model
	ID        uint           `gorm:"primaryKey" json:"id" form:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime:false" json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `gorm:"autoCreateTime:false" json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"autoCreateTime:false" json:"deleted_at" form:"deleted_at"`
	Name      string         `json:"name" form:"name"`
}

type Product struct {
	gorm.Model
	UserID uint   `json:"user_id" form:"user_id"`
	Name   string `json:"name" form:"name"`
	Price  int    `json:"price" form:"price"`
}

type Address struct {
	gorm.Model
	UserID  uint   `json:"user_id" form:"user_id"`
	Address string `json:"address" form:"address"`
}

type ResponseUserFormat struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB *gorm.DB

// database connection
func InitDB() {

	// declare struct config & variable connectionString
	connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/db_be9_gorm?charset=utf8&parseTime=True&loc=UTC"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Jurusan{})
	DB.AutoMigrate(&Product{})
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
	e.POST("/jurusan", CreateJurusanController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":80"))
}

func ResponseFailed(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "error",
		"message": msg,
	}
}

func ResponseSuccessNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}

func ResponseSuccessWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}

func GetUsersController(c echo.Context) error {
	var users []User
	tx := DB.Find(&users)
	if tx.Error != nil {

		return c.JSON(http.StatusInternalServerError, ResponseFailed("failed to get data"))
	}
	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"status":  "success",
	// 	"message": "success get data",
	// 	"data":    users,
	// })

	responseData := []ResponseUserFormat{}
	for _, value := range users {
		var response ResponseUserFormat
		response.ID = value.ID
		response.Name = value.Name
		response.Email = value.Email

		responseData = append(responseData, response)
	}
	return c.JSON(http.StatusOK, ResponseSuccessWithData("success get data", responseData))
}

func CreateUserController(c echo.Context) error {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("error", err)

		return c.JSON(http.StatusBadRequest, ResponseFailed("failed create user bind"))
	}
	// gmt, _ := time.LoadLocation("UTC")
	// time.Local = gmt
	// user.CreatedAt = time.Now().In(gmt)
	// fmt.Println(user.CreatedAt)
	tx := DB.Create(&user)
	if tx.Error != nil {

		return c.JSON(http.StatusInternalServerError, ResponseFailed("failed to insert data"))
	}
	if tx.RowsAffected == 0 {
		// return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 	"status":  "error",
		// 	"message": "failed to insert data",
		// })
		return c.JSON(http.StatusInternalServerError, ResponseFailed("failed to insert data"))
	}
	// To change its value, you could use `Update`
	// DB.Model(&user).Update("created_at", time.Now().In(gmt))

	return c.JSON(http.StatusOK, ResponseSuccessNoData("success insert data"))

}

func CreateJurusanController(c echo.Context) error {

	type jurusanRequestFormat struct {
		Name string `json:"name" form:"name"`
	}
	jurusanReqBody := jurusanRequestFormat{}
	err := c.Bind(&jurusanReqBody)
	if err != nil {
		fmt.Println("error", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "failed create jurusan bind",
		})
	}
	// gmt, _ := time.LoadLocation("UTC")
	// time.Local = gmt
	// jurusan.CreatedAt = time.Now().In(gmt)
	// jurusan.UpdatedAt = time.Now().In(gmt)
	// fmt.Println(jurusan.CreatedAt)
	jurusan := Jurusan{
		Name: jurusanReqBody.Name,
	}
	tx := DB.Create(&jurusan)
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
	// DB.Model(&ju).Update("created_at", time.Now().In(gmt))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success insert data",
	})

}

// func UpdateUser(c echo.Context) error {
// 	user := User{}
// 	user.ID = 111
// 	db.Model(&user).Update("name", "hello")
// }
