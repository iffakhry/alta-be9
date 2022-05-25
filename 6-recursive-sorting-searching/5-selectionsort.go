package main

import "fmt"

func selectionSort(elements []int) []int {
	n := len(elements)
	for k := 0; k < n; k++ {
		minimal := k
		for j := k + 1; j < n; j++ {
			if elements[j] < elements[minimal] {
				minimal = j
			}
		}
		//swap
		// a,b=b,a
		elements[k], elements[minimal] = elements[minimal], elements[k]
	}
	return elements
}

func main() {
	var data = []int{1, 4, 3, 5, 7, 8, 5, 10}
	fmt.Println("sebelum:", data)
	var datasort = selectionSort(data)
	fmt.Println("sesudah:", datasort)

}
