package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	temp := make([]int, len(jumps)+1)

	temp[0] = 0
	temp[1] = int(math.Abs(float64(jumps[0] - jumps[1])))
	for i := 2; i < len(jumps); i++ {
		jump1 := int(math.Abs(float64(jumps[i]-jumps[i-1]))) + temp[i-1]
		fmt.Println("jump1", jump1, "jump1", jumps[i], "jump1", jumps[i-1])
		jump2 := int(math.Abs(float64(jumps[i]-jumps[i-2]))) + temp[i-2]
		fmt.Println("jump2", jump2, "jump2", jumps[i], "jump2", jumps[i-2])
		temp[i] = int(math.Min(float64(jump1), float64(jump2)))
		fmt.Println(temp[i])
	}
	return temp[len(jumps)-1]

}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30
	// fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40

}
