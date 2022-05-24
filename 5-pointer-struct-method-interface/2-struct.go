package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       uint
	Weight    uint
	Telp      []string
}

type Mahasiswa struct {
	NIM      string
	Nama     string
	Jurusan  string
	Semester uint
	Alamat   string
}

type User struct {
	Firstname string
	Lastname  string
	Addresses Address
}

type Address struct {
	Street string
	City   string
}

func main() {
	// var person1firstname = "john"
	// var person1lastname = "doe"

	// var person2firstname = "mike"

	var person1 Person
	person1.Firstname = "John"
	person1.Lastname = "doe"
	person1.Age = 10
	person1.Weight = 30
	person1.Telp = []string{"12345", "34567", "23345"}
	fmt.Println(person1)
	fmt.Println(person1.Telp[len(person1.Telp)-1])

	// var person2 Person
	// person2.Firstname = "Mike"
	// person2.Lastname = "joe"
	// person2.Age = 20
	// person2.Weight = 50
	// fmt.Println(person2)

	var budi Mahasiswa
	budi.NIM = "11223344556677"
	budi.Nama = "Budi sudarsono"
	budi.Semester = 1
	budi.Jurusan = "Ekonomi"
	budi.Alamat = "Jakarta"

	var ahmad Mahasiswa
	ahmad.NIM = "22334455667788"
	ahmad.Nama = "Ahmad sudarsono"
	ahmad.Semester = 1
	ahmad.Jurusan = "Ekonomi"
	ahmad.Alamat = "Jakarta"

	var user1 User
	user1.Firstname = "Ahmad"
	user1.Lastname = "abu"
	// var address1 = Address{"Jl Jawa", "jakarta"}
	// user1.Addresses = address1
	user1.Addresses = Address{
		Street: "Jl.Lurus",
		City:   "Jakarta",
	}
	fmt.Println(user1.Addresses.City)

	type Mobil struct {
		Type  string
		CC    uint
		Color string
	}

	var data = map[string]Mobil{}
	data["123"] = Mobil{
		Type:  "APV",
		CC:    1200,
		Color: "Red",
	}

	data["234"] = Mobil{
		Type:  "Sedan",
		CC:    2000,
		Color: "Black",
	}

	for k, v := range data {
		fmt.Println(k, v.Type, v.CC)
	}

}
