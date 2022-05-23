package main

import "fmt"

func main() {
	var data = [4]int{1, 2, 3, 4}
	var temp1 = data[:2]
	var temp2 = data[2:]
	fmt.Println(temp1)
	fmt.Println(temp2)
	var temp3 = append(temp1, 10)
	fmt.Println(temp3)
	temp3 = append(temp3, temp2...)
	fmt.Println(temp3)

}
