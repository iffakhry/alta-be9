package main

import "fmt"

func findMaxMin(data []int) (max int, min int) {
	max = data[0]
	min = data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
		if data[i] < min {
			min = data[i]
		}
	}
	return
}

func main() {
	data := []int{2, 6, 10, 9, 8, 1, 3}
	max, min := findMaxMin(data)
	fmt.Println(max, min)
}
