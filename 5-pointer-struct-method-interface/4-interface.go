package main

import "fmt"

type calculate interface {
	large() int
	keliling() int
	print(str string) string
}

type square struct {
	side int
}

func (s square) large() int {
	return s.side * s.side
}

func (s square) keliling() int {
	return 4 * s.side
}

func (s square) print(data string) string {
	return "hello" + data
}

//struct circle
type circle struct {
	jari int
}

func (c circle) large() int {
	return c.jari * c.jari
}

func (c circle) keliling() int {
	return 4 * c.jari
}

func (c circle) print(data string) string {
	return "hello" + data
}

func main() {
	var dimResult calculate
	sqr := square{10}
	dimResult = sqr
	fmt.Println("large square :", dimResult.large())
	fmt.Println(" square :", dimResult.print("alta"))
	fmt.Println("struct", sqr.large())

	fmt.Println("=======")
	var dimCircleResult calculate
	cir := circle{10}
	dimCircleResult = cir
	fmt.Println("large square :", dimCircleResult.large())
}
