package db

import "log"

func BuatTabel() {

	db, err := ConnDB()
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	buatTabelAdmins := `
    CREATE TABLE IF NOT EXISTS admins (
		id INT(11) AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(50),
		password VARCHAR(300)
		
    );`

	buatTabelCatat := `
    CREATE TABLE IF NOT EXISTS catat (
		nomer INT(255) AUTO_INCREMENT PRIMARY KEY,
        id VARCHAR(255),
		grup INT(100),
		nama VARCHAR(255),
		jenis VARCHAR(5),
		nik INT(255),
		alamat VARCHAR(255),
		pos VARCHAR(255),
		waktu TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	buatTabelGelang := `
    CREATE TABLE IF NOT EXISTS gelang (
        id VARCHAR(255) UNIQUE
	   
    );`
	buatTabelLokasi := `
    CREATE TABLE IF NOT EXISTS lokasi (
        id INT(50) AUTO_INCREMENT PRIMARY KEY,
        lat VARCHAR(200),
		longt VARCHAR(200),
		nama_pos VARCHAR(50),
		keterangan TEXT
	   
    );`

	buatTabelPendaki := `
    CREATE TABLE IF NOT EXISTS pendaki (
        nom INT(100) AUTO_INCREMENT PRIMARY KEY,
        id VARCHAR(255) UNIQUE,
		grup INT(50),
		nama VARCHAR(300),
		jenis VARCHAR(5),
		nik INT(255) UNIQUE,
	    alamat VARCHAR(300)
    );`

	buatTabelSelesai := `
    CREATE TABLE IF NOT EXISTS selesai (
        no INT(100) AUTO_INCREMENT PRIMARY KEY,
		grup INT(100),
        nama VARCHAR(300),
	    jenis VARCHAR(5),
		nik INT(255),
        alamat VARCHAR(300),
		tgl TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        catatan TEXT
    );`

	_, err = db.Exec(buatTabelAdmins)
	if err != nil {
		log.Fatalf("Error pembuatan tabel admins: %v", err)
	}
	_, err = db.Exec(buatTabelCatat)
	if err != nil {
		log.Fatalf("Error pembuatan tabel catat %v", err)
	}
	_, err = db.Exec(buatTabelGelang)
	if err != nil {
		log.Fatalf("Error pembuatan tabel gelang %v", err)
	}
	_, err = db.Exec(buatTabelLokasi)
	if err != nil {
		log.Fatalf("Error pembuatan tabel lokasi: %v", err)
	}

	_, err = db.Exec(buatTabelPendaki)
	if err != nil {
		log.Fatalf("Error pembuatan tabel pendaki: %v", err)
	}
	_, err = db.Exec(buatTabelSelesai)
	if err != nil {
		log.Fatalf("Error pembuatan tabel selesai: %v", err)
	}

}
