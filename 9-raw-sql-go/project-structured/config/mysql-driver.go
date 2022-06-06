package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() *sql.DB {
	// buat koneksi ke DB
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	dbConnection := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("success")
	}
	return db
}
