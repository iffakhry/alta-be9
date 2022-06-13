package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Books struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Author    []string `json:"author"`
	Publisher string   `json:"publisher"`
}

type article struct {
	ID      int
	Title   string
	Content string
}

var data = []article{
	article{1, "lorem", "lorem"},
	article{2, "ipsum", "ipsum"},
}

func main() {
	//inisialisasi echo
	e := echo.New()

	//routing endpoint
	v1 := e.Group("/v1")
	v1.GET("/articles", HelloJSONController)
	v1.POST("/articles", HelloController)

	e.GET("/books", GetBooksController)
	e.GET("/books/:id", GetBookByIdController)

	e.GET("/products", GetProductsController)

	e.POST("/users", CreateUserFormController)
	v1.POST("/users", CreateUserBindController)

	//start server di port 80
	e.Start(":80")
}

// handler - Simple handler to make sure environment is setup
func HelloController(c echo.Context) error {
	// return the string "Hello World" as the response body
	// with an http.StatusOK (200) status
	return c.String(http.StatusInternalServerError, "Hello World")
}

func HelloJSONController(c echo.Context) error {
	// return the string "Hello World" as the response body
	// with an http.StatusOK (200) status
	return c.JSON(http.StatusOK, data)
}

func GetBooksController(c echo.Context) error {
	books := []Books{
		{
			ID:        1,
			Title:     "One Piece",
			Author:    []string{"oda"},
			Publisher: "Gramedia",
		},
		{
			ID:        2,
			Title:     "Naruto",
			Author:    []string{"Masashi K"},
			Publisher: "Gramedia",
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get book data",
		"data":    books,
	})
}

//param
func GetBookByIdController(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("idnya", id)
	idBook, _ := strconv.Atoi(id)
	books := Books{
		ID:        idBook,
		Title:     "One Piece",
		Author:    []string{"oda"},
		Publisher: "Gramedia",
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get book data",
		"data":    books,
	})
}

//Query Param
func GetProductsController(c echo.Context) error {
	match := c.QueryParam("match")
	name := c.QueryParam("name")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get products",
		"match":   match,
		"name":    name,
		"data":    []string{"keyboard", "mouse", "monitor", match}, // hardcode data
	})
}

type User struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func CreateUserFormController(c echo.Context) error {
	// get data from value
	id := c.FormValue("id")
	name := c.FormValue("name")
	email := c.FormValue("email")

	iduser, _ := strconv.Atoi(id)
	var user User
	user.ID = iduser
	user.Name = name
	user.Email = email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user form",
		"user":     user,
	})
}

func CreateUserBindController(c echo.Context) error {
	// binding data
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("error", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "failed create user bind",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user bind",
		"user":     user,
	})
}
