package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnDB() (*sql.DB, error) {

	basisdata, err := sql.Open("mysql", "root:@tcp(localhost:3306)/pendaki_go")
	if err != nil {
		log.Fatal(err)
	}

	// Menguji koneksi ke database
	err = basisdata.Ping()
	if err != nil {
		log.Fatal("Koneksi ke database gagal:", err)
	}

	fmt.Println("Koneksi sukses!")

	return basisdata, nil
}
