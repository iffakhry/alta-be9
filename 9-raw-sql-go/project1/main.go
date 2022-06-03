package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Mahasiswa struct {
	ID           int
	Nama         string
	JenisKelamin string
	Alamat       string
	JurusanId    int
	Telp         string
	Status       string
}

func main() {
	// defer fmt.Println("satu")
	// defer fmt.Println("dua")
	// fmt.Println("empat")
	// fmt.Println("lima")
	// defer fmt.Println("tiga")

	// buat koneksi ke DB
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	db, err := sql.Open("mysql", "root:qwerty123@tcp(localhost:3306)/db_be9")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("success")
	}

	defer db.Close()

	fmt.Println("Pilih Menu: (1: Lihat Data)/(2: Input Data)")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		results, err := db.Query("SELECT id, nama, jenis_kelamin, alamat, jurusan_id, telp, status FROM mahasiswa")
		if err != nil {
			fmt.Println("error", err.Error())
		}

		var dataAll []Mahasiswa
		for results.Next() {
			var mahasiswa Mahasiswa
			err := results.Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.JenisKelamin, &mahasiswa.Alamat, &mahasiswa.JurusanId, &mahasiswa.Telp, &mahasiswa.Status)
			// mahasiswa.Nama --> Budi 5
			if err != nil {
				fmt.Println("error scan", err.Error())
			}
			dataAll = append(dataAll, mahasiswa)
		}
		// dataAll --> adi, Cindy, Budi 1, .... Budi 5

		for _, v := range dataAll {
			fmt.Println("ID:", v.ID, "Nama:", v.Nama, "Status:", v.Status)
		}
		// fmt.Println(dataAll[0].ID)

	case 2:
		newMahasiswa := Mahasiswa{}
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

		var query = fmt.Sprintf("insert into mahasiswa (id, nama, jenis_kelamin, alamat, jurusan_id, telp) values (%d, '%s', '%s', '%s', %d, '%s')", newMahasiswa.ID, newMahasiswa.Nama, newMahasiswa.JenisKelamin, newMahasiswa.Alamat, newMahasiswa.JurusanId, newMahasiswa.Telp)
		result, err := db.Exec(query)
		if err != nil {
			fmt.Println("error insert", err.Error())
		} else {
			fmt.Println("Insert Success")
			row, _ := result.RowsAffected()
			fmt.Println("row affect", row)
		}
	}

}
