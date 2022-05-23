package main

import "fmt"

var (
	data int
)

func ProsesNilai(nilai int) string {
	var result string
	// var data = 10
	if nilai == 100 {
		// result += " [SEMPURNA]"
		return "[SEMPURNA]"
	}

	if nilai >= 80 && nilai <= 100 {

		fmt.Println(data)
		result = "A"
		// return "A"
	} else if nilai >= 65 && nilai <= 79 {
		result = "B+"
		// return "B+"
	} else {
		result = "Invalid"
		// return "Invalid"
	}
	data = 100
	return result

}

func main() {
	data = 10
	fmt.Println(ProsesNilai(90))
}
