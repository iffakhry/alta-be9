package main

import (
	"fmt"
	"math"
)

func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true

}

func PrimeX(number int) int {
	var counter int
	var bilangan int = 1
	for counter < number {
		bilangan++ // 2
		if checkPrime(bilangan) {
			counter++
		}
		// fmt.Println("count:", counter, "bil:", bilangan)
	}
	return bilangan
}

func main() {
	// fmt.Println(checkPrime(11))
	fmt.Println(PrimeX(1))  // 2
	fmt.Println(PrimeX(5))  // 11
	fmt.Println(PrimeX(8))  // 19
	fmt.Println(PrimeX(9))  // 23
	fmt.Println(PrimeX(10)) // 29
}
