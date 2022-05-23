package main

import "fmt"

func ArrayUnique(arrayA, arrayB []int) []int {
	// your code here
	result := []int{}
	mapB := map[int]bool{}
	for _, valueB := range arrayB {
		mapB[valueB] = true
	}
	// fmt.Println(mapB)
	// value, ok := mapB[4]
	// fmt.Println("value:", value, "isExist", ok)

	for _, valueA := range arrayA {
		_, isExist := mapB[valueA]
		// fmt.Println("value:", value, "isExist", isExist)
		if !isExist { // jika key di mapB tidak ada, maka tambahkan valueA ke slice
			result = append(result, valueA)
		}
	}
	return result
}

func main() {
	fmt.Println(ArrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}))   // [2 4]
	fmt.Println(ArrayUnique([]int{10, 20, 30, 40}, []int{5, 10, 15, 59})) // [20 30 40]
	fmt.Println(ArrayUnique([]int{1, 3, 7}, []int{1, 3, 5}))              // [7]
	fmt.Println(ArrayUnique([]int{3, 8}, []int{2, 8}))                    // [3]
	fmt.Println(ArrayUnique([]int{1, 2, 3}, []int{3, 2, 1}))              // []
}
