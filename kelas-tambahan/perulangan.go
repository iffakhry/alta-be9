package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// 	//0, -1, -2,-3,-4 ....
	// }

	// for i := 10; i > 0; i-- {
	// 	fmt.Println(i)

	// }

	// j := 0
	// for j <= 10 {
	// 	fmt.Println(j)
	// 	j++
	// }

	/*
		baris 1 --> kolom 1-4
		baris 2 --> kolom 1-4
		baris 3 --> kolom 1-4
	*/
	var counter int
	for baris := 1; baris <= 3; baris++ { //i=2
		for kolom := 1; kolom <= baris; kolom++ { //j=3
			// fmt.Println("i:", baris, "j:", kolom)
			counter++
			fmt.Print(kolom, " ")
		}
		fmt.Println()
	}

	Ulang(100, 0)

	Ulang(20, 5)
}

func Ulang(max int, min int) {
	for i := max; i >= min; i-- {
		fmt.Println(i)
	}
}
