package main

import "fmt"

func main() {
	var bilangan [5]int //0,1,2,3,4,5,6,7
	bilangan[0] = 10
	bilangan[1] = 15
	bilangan[4] = 25
	bilangan[2] = 35
	bilangan[len(bilangan)-1] = 10

	fmt.Println(bilangan)
	fmt.Println(len(bilangan))

	var bilangans = [5]int{0: 1, 3: 6, 4: 2}
	fmt.Println(bilangans)

	var data = [...]int{1, 3, 5}
	fmt.Println(len(data))

	var countries1 = [2]string{"indonesia", "malaysia"}
	var countries2 = [2]string{"Amerika", "Inggris"}
	var countries3 []string
	countries3 = append(countries1[:], countries2[:]...)
	// fmt.Println(countries1, countries2)
	fmt.Println(countries3)
}
