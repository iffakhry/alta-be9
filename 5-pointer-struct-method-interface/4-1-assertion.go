package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	var data = map[string]interface{}{}
	data["nama"] = "John doe"
	data["age"] = 30
	data["isAktif"] = true
	fmt.Println(data)
}
