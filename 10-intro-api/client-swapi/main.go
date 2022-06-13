package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A Response struct to map the Entire Response
type StarWarsPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Films     []string `json:"films"`
}

func main() {
	response, _ := http.Get("https://swapi.dev/api/people/1")

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var people1 StarWarsPeople
	json.Unmarshal(responseData, &people1)

	fmt.Println("---- Print Field ----")
	fmt.Println(people1.Name)
	fmt.Println(people1.Height)
	fmt.Println(people1.HairColor)
	// fmt.Println(people1.Films[0])

	for _, v := range people1.Films {
		fmt.Println(v)
	}

}
