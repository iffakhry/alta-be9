package entities

import "database/sql"

type Mahasiswa struct {
	ID           int
	Nama         string
	JenisKelamin string
	Alamat       string
	JurusanId    int
	Telp         string
	Status       sql.NullString
}
