package main

import "fmt"

//pass by value
func ubahData(data string) string {
	fmt.Println("data:", data)
	data = "doe"
	return data
}

//pass by reference
func ubahDataPointer(data *string) string {
	fmt.Println("address:", data)
	*data = "pak doe" // mengubah value dari alamat yang ada di var data
	return *data
}

func main() {
	var nama = "john"
	namaResult := ubahData(nama)
	fmt.Println("nama:", nama)             //john
	fmt.Println("namaResult:", namaResult) //doe
	fmt.Println("=============")
	namaResultPointer := ubahDataPointer(&nama)
	fmt.Println("nama:", nama)                    //
	fmt.Println("namaResult:", namaResultPointer) //pak doe
}
