package mahasiswa

import (
	_entities "be9/rawsqlgo/entities"
	"database/sql"
	"fmt"
)

func GetAllMahasiswa(db *sql.DB) []_entities.Mahasiswa {
	results, err := db.Query("SELECT id, nama, jenis_kelamin, alamat, jurusan_id, telp, status FROM mahasiswa")
	if err != nil {
		fmt.Println("error", err.Error())
	}
	defer db.Close() //close db connection when process done

	var dataAll []_entities.Mahasiswa
	for results.Next() {
		var mahasiswa _entities.Mahasiswa
		err := results.Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.JenisKelamin, &mahasiswa.Alamat, &mahasiswa.JurusanId, &mahasiswa.Telp, &mahasiswa.Status)
		// mahasiswa.Nama --> Budi 5
		if err != nil {
			fmt.Println("error scan", err.Error())
		}
		dataAll = append(dataAll, mahasiswa)
	}

	return dataAll
}

/*
Dalam fungsi Create ini saya menggunakan pendekatan prepare statement. untuk meningkatkan performa aplikasi saat akses ke DB.
untuk detail penjelasan terkait prepare statement bisa dibaca di link berikut:
https://medium.com/pujanggateknologi/prepared-statement-di-go-927b1a8863ec
*/

func CreateMahasiswa(db *sql.DB, newMahasiswa _entities.Mahasiswa) (int, error) {
	// var query = fmt.Sprintf("INSERT INTO mahasiswa (id, nama, jenis_kelamin, alamat, jurusan_id, telp, status) values (%d, '%s', '%s', '%s', %d, '%s','%s')", newMahasiswa.ID, newMahasiswa.Nama, newMahasiswa.JenisKelamin, newMahasiswa.Alamat, newMahasiswa.JurusanId, newMahasiswa.Telp, newMahasiswa.Status.String)
	var query = (`INSERT INTO mahasiswa (id, nama, jenis_kelamin, alamat, jurusan_id, telp, status) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	result, err := statement.Exec(newMahasiswa.ID, newMahasiswa.Nama, newMahasiswa.JenisKelamin, newMahasiswa.Alamat, newMahasiswa.JurusanId, newMahasiswa.Telp, newMahasiswa.Status.String)

	defer db.Close() //close db connection when process done

	if err != nil {
		// fmt.Println("error insert", err.Error())
		return 0, err
	} else {
		// fmt.Println("Insert Success")
		row, _ := result.RowsAffected()
		return int(row), nil
	}
}

// please complete this function with your own code
func UpdateMahasiswa() {

}

// please complete this function with your own code
func DeleteMahasiswa() {

}
