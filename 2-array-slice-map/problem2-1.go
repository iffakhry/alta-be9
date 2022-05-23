package main

import (
	"fmt"
)

func FaktorBilangan(n int) string {
	// your code here
	var result string
	for i := 1; i <= n; i++ { //1,2,3,4,5,6 ....
		if n%i == 0 {
			result += fmt.Sprintf("%d\n", i)
			// result = result + "\n" + strconv.Itoa(i)
		}
	}
	fmt.Printf("tipe: %T %s\n", n, result)
	return result
}

func main() {
	var number int
	fmt.Println("masukkan bilangan: ")
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilangan(number))

	/*
	   6 --> 1,2,3,6
	    --> 6,3,2,1
	*/

	// var data = "*"
	// dataInt, err := strconv.Atoi(data)
	// fmt.Println("data", dataInt)
	// fmt.Println("err", err)
	// if err != nil {
	// 	fmt.Println("ada error saat konversi")
	// }

	// res := strconv.Itoa(10)
	// fmt.Println("res", res)

}
