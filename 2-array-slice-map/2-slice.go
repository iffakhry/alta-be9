package main

import (
	"fmt"
	"reflect"
)

func main() {

	var primes = [5]int{2, 3, 5, 7, 11}

	// Creating a slice from the array
	var part_primes []int = primes[1:4]
	var part_primes1 []int = primes[1:]
	fmt.Println(part_primes1)

	part_primes = append(part_primes, 10000)
	// menambah data ke slice akan menambah data ke array juga

	fmt.Println(reflect.ValueOf(part_primes).Kind())
	fmt.Println(part_primes)
	var datas = []int{1, 2, 3, 4}
	fmt.Println(datas)
	datas = append(datas, 18)
	fmt.Println(datas)

}
