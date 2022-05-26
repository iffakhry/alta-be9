package main

import "fmt"

/*
input: angka (1 to n)
output : Ganjil/Genap --> tergantung angka yang diinputkan.

1 -> Ganjil
2 -> Genap
3 -> Ganjil
4 -> Genap
5 -> Ganjil
11 -> Ganjil

Genap -> Habis dibagi 2
Ganjil -> tidak habis dibagi 2
*/

func GanjilGenap(angka int) string {
	fmt.Println(angka)
	var result string
	if angka%2 == 0 {
		result = "genap"
	} else {
		result = "ganjil"
	}
	return result
}

func CheckData() (data int, data2 string) {
	data = 20
	return
}

func main() {
	// var angka int = 10
	var status = GanjilGenap(100)
	fmt.Println(status)
}
