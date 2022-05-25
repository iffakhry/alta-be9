package main

import "fmt"

func fibonacci(number int) int {
	//base case
	if number <= 1 {
		return number
	}

	//recurrent
	return fibonacci(number-1) + fibonacci(number-2)

}

func main() {
	fmt.Println(fibonacci(6))
}
