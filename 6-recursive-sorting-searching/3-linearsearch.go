package main

import "fmt"

// fungsi untuk mencari sebuah bilangan berada di index berapa.
// jika tidak ditemukan, maka akan mengembalikan -1
func linearSearch(elements []int, x int) int {
	for i := 0; i < len(elements); i++ {
		if elements[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	var data = []int{1, 4, 3, 5, 7, 8, 5, 10}
	result := linearSearch(data, 5)
	if result != -1 {
		fmt.Println("data ditemukan di index", result)
	} else {
		fmt.Println("data tidak ditemukan")
	}
}
