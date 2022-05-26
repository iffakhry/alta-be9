package main

import "fmt"

func main() {
	var kata string = "ALTERRA"
	for _, v := range kata {
		fmt.Println(v)
	}

	for i := 0; i < len(kata); i++ {
		fmt.Println(kata[i]) //kata[2]
	}
}
