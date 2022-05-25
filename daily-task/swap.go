package main

import "fmt"

func swap1(a, b *int) (int, int) {
	temp := *a
	*a = *b
	*b = temp
	return *a, *b
}

func swap2(a, b *int) (int, int) {
	*a, *b = *b, *a
	return *a, *b
}

func main() {
	var a = 50
	var b = 90
	swap2(&a, &b)
	fmt.Println(a, b)
}
