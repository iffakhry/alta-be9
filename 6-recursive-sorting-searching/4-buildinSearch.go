package main

import (
	"fmt"
	"sort"
)

func main() {
	var data = []int{1, 3, 5, 7, 10, 20, 25}
	value := 22
	findIndex := sort.SearchInts(data, value)
	fmt.Println(findIndex)
	if value == data[findIndex] {
		fmt.Println("data ditemukan")
	} else {
		fmt.Println("data tidak ditemukan")
	}

}
