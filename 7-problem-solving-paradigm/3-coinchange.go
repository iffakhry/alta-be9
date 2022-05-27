package main

import (
	"fmt"
	"sort"
)

/*
[10,25,5,1]

42 = 1,1,1,1,1,1,1 ... (42) --> 42 koin
42 = 5,5,5,5,5,5 ...,1,1 (42)
42 = 10,10,10,10,1,1 -> (42) --> 6 koin
42 = 25,10,5,1,1 -> (42) --> 5 koin
57 = 25,25,5,1,1
57 = 25,10,10,10,1,1
*/

func coinChange(money int, coins []int) []int {
	result := []int{}

	sort.SliceStable(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	// fmt.Println(coins)
	for _, vcoin := range coins {
		if money >= vcoin {
			for money >= vcoin {
				result = append(result, vcoin)
				money = money - vcoin
			}
		} else {
			continue
		}
	}

	// for i := 0; i < count; i++ { //i=2
	// 	if condition {
	// 		continue
	// 	}
	// }

	return result
}

func main() {
	coinValue := []int{10, 25, 5, 1}
	fmt.Println(coinChange(42, coinValue)) //25,10,5,1,1
	fmt.Println(coinChange(57, coinValue)) //25,25,5,1,1
}
