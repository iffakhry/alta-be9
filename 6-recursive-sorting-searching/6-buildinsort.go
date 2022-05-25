package main

import (
	"fmt"
	"sort"
)

func main() {
	var data = []int{1, 4, 3, 5, 7, 8, 5, 10}
	fmt.Println("sebelum:", data)
	sort.Ints(data)
	fmt.Println("sesudah:", data)

	var datastring = []string{"cba", "cab", "abc"}
	fmt.Println("sebelum:", datastring)
	sort.Strings(datastring)
	fmt.Println("sesudah:", datastring)

	datastatus := sort.IntsAreSorted(data)
	fmt.Println(datastatus)

	var data2 = []string{"cba", "cab", "abc"}
	/*
		data2[i] < data2[j] --> untuk ascending
		data2[i] > data2[j] --> untuk descending
	*/
	sort.SliceStable(data2, func(i, j int) bool {
		return data2[i] > data2[j]
	})
	fmt.Println(data2)

}
