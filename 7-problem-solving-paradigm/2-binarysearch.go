package main

import "fmt"

func BinarySeach(data []int, x int) int {
	var kiri = 0
	var kanan = len(data) - 1
	var count int
	for kiri <= kanan {
		count++
		var tengah = (kiri + kanan) / 2
		fmt.Println("tengah", tengah)
		fmt.Println("count ke:", count)
		if data[tengah] == x {
			return tengah
		} else if x < data[tengah] {
			kanan = tengah - 1
		} else if x > data[tengah] {
			kiri = tengah + 1
		}
	}
	return -1
}

func linearSearch(elements []int, x int) int {
	var count int
	for i := 0; i < len(elements); i++ {
		count++
		fmt.Println("count linear ke: ", count)
		if elements[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	var data = []int{1, 1, 3, 5, 5, 6, 7, 9, 10}
	fmt.Println(BinarySeach(data, 10))
	fmt.Println("---------")
	fmt.Println(linearSearch(data, 10))
}
