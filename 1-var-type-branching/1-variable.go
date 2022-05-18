package main

import "fmt"

var Name string

func main() {
	//long declaration
	var alamat string              //deklarasi variabel
	alamat = "Bali"                //memberikan nilai ke variabel
	fmt.Println("alamat1", alamat) //mencetak variabel alamat
	alamat = "Jakarta"             //memberikan nilai ke variabel
	fmt.Println("alamat2", alamat) //mencetak variabel alamat
	alamat = "Jayapura"            //memberikan nilai ke variabel
	fmt.Println("alamat2", alamat) //mencetak variabel alamat

	var umur uint8 = 10
	fmt.Println(umur)

	var angka int
	fmt.Println(angka) //0

	var firstName, lastName, nickName string = "Alterra", "Academy", "alta"
	fmt.Println(firstName) //Alterra
	fmt.Println(lastName)  //Academy
	fmt.Println(nickName)  //alta

	//short
	country := "Indonesia"
	fmt.Println(country)

	var check bool
	check = true
	fmt.Println(check)

	const pi float32 = 3.14
	// pi = 5.13 // error
	fmt.Println(pi)

	//operator
	bil1 := 10
	bil2 := 20
	hasil := bil1 * bil2
	fmt.Println(hasil)
	bil1 += 50 // bil1 = bil1 + 50
	fmt.Println(bil1)

	kata1 := "alterra"
	kata2 := "academy"
	katagabung := kata1 + "-" + kata2
	fmt.Println(katagabung)
}
