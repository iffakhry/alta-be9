package main

import (
	_hitung "be9/unittest/calculator"
	"fmt"
)

/*
jalankan perintah go mod dibawah, maka akan men-generate file go.mod:

go mod init namaproject
*/

func main() {
	hasil := _hitung.Tambah(10, 15)
	fmt.Println(hasil)
	hasil2 := _hitung.Kurang(30, 20)
	fmt.Println(hasil2)
}
