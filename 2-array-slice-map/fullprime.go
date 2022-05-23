package main

import (
	"fmt"
)

func main() {
	var angka = 23
	angkastr := fmt.Sprintf("%d", angka)
	for _, v := range angkastr {
		fmt.Println(string(v))
	}
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a)

	var a = 10
	fmt.Println(&a)

}
