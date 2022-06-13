package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StarWars struct {
	Count    int              `json:"count"`
	Next     string           `json:"next"`
	Previous string           `json:"previous"`
	Results  []StarWarsPeople `json:"results"`
}

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
	response, _ := http.Get("https://swapi.dev/api/people")

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var people StarWars
	json.Unmarshal(responseData, &people)

	fmt.Println("---- Print Field ----")
	fmt.Println("Hello", people.Results[1].Name)
	// fmt.Println(people1.Height)
	// fmt.Println(people1.HairColor)
	// fmt.Println(people1.Films[0])

	for _, v := range people.Results {
		fmt.Println(v.Name)
	}

}
