package main

import "fmt"

type Circle struct {
	jari int
}

type Square struct {
	sisi int
}

func (c Circle) HitungLuas() float64 {
	luas := 3.14 * float64(c.jari) * float64(c.jari)
	return luas
}

func (s Square) HitungLuas() float64 {
	luas := s.sisi * s.sisi
	return float64(luas)
}

func main() {
	circle := Circle{
		jari: 5,
	}
	fmt.Println(circle.HitungLuas())

	square := Square{
		sisi: 10,
	}
	fmt.Println(square.HitungLuas())

}
