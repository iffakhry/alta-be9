package main

import (
	_config "be9/rawsqlgo/config"
	_mahasiswaController "be9/rawsqlgo/controllers/mahasiswa"
	_entities "be9/rawsqlgo/entities"
	"database/sql"
	"fmt"
)

var DBConn *sql.DB

func init() {
	DBConn = _config.ConnectionDB()
}

func main() {

	// defer DBConn.Close()

	fmt.Println("Pilih Menu: (1: Lihat Data)/(2: Input Data)")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		results := _mahasiswaController.GetAllMahasiswa(DBConn)

		for _, v := range results {
			fmt.Println("ID:", v.ID, "Nama:", v.Nama, "Status:", v.Status.String)
		}

	case 2:
		newMahasiswa := _entities.Mahasiswa{}
		fmt.Println("Input ID:")
		fmt.Scanln(&newMahasiswa.ID)
		fmt.Println("Input Nama:")
		fmt.Scanln(&newMahasiswa.Nama)
		fmt.Println("Input Jenis Kelamin:")
		fmt.Scanln(&newMahasiswa.JenisKelamin)
		fmt.Println("Input Alamat:")
		fmt.Scanln(&newMahasiswa.Alamat)
		fmt.Println("Input Jurusan ID:")
		fmt.Scanln(&newMahasiswa.JurusanId)
		fmt.Println("Input Telepon:")
		fmt.Scanln(&newMahasiswa.Telp)
		fmt.Println("Input Status:")
		fmt.Scanln(&newMahasiswa.Status.String)

		row, err := _mahasiswaController.CreateMahasiswa(DBConn, newMahasiswa)

		if err != nil {
			fmt.Println("error insert", err.Error())
		} else {
			fmt.Println("Insert Success")
			fmt.Println("row affect", row)
		}
	}
}
