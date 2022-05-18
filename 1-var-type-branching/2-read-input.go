package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	var data = "alterra"
	var angka = 10
	fmt.Println("halo")
	fmt.Print("guys")
	fmt.Printf("halo %s %d \n", data, angka)

	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName) //inputan dari user
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
