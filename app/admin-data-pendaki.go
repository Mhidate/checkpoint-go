package app

import (
	db "checkpoint-go/db/config"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pendaki struct {
	Nom     int    `json:"nom" form:"nom"`
	Id      string `json:"id" form:"id"`
	Grup    int    `json:"grup" form:"grup"`
	Nama    string `json:"nama" form:"nama"`
	Jenis   string `json:"jenis" form:"jenis"`
	Nik     int    `json:"nik" form:"nik"`
	Alamat  string `json:"alamat" form:"alamat"`
	Catatan string `form:"catatan"`
}

func DataPendaki(c *gin.Context) {

	database, _ := db.ConnDB()
	defer database.Close()

	rows, err := database.Query("SELECT * FROM pendaki")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var results []Pendaki

	for rows.Next() {
		var pendaki Pendaki
		if err := rows.Scan(&pendaki.Nom, &pendaki.Id, &pendaki.Grup, &pendaki.Nama, &pendaki.Jenis, &pendaki.Nik, &pendaki.Alamat); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, pendaki)
	}
	c.JSON(http.StatusOK, gin.H{"pendaki": results})
}

func ButtonSelesai(c *gin.Context) {
	id := c.Param("id") // Mendapatkan ID dari URL

	database, _ := db.ConnDB()
	defer database.Close()

	var pendaki Pendaki
	err := database.QueryRow("SELECT id, grup, nama, jenis, nik, alamat FROM pendaki WHERE id = ?", id).Scan(&pendaki.Id, &pendaki.Grup, &pendaki.Nama, &pendaki.Jenis, &pendaki.Nik, &pendaki.Alamat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// fmt.Println("Data Pendaki:", pendaki.Nama)

	// Render template edit dengan data pendaki
	c.HTML(http.StatusOK, "form-selesai.html", gin.H{
		"title":   "Form Pendaki Selesai",
		"pendaki": pendaki, // Tambahkan objek pendaki ke dalam data yang dikirimkan
	})
}

func ProsesSelesai(c *gin.Context) {
	var pendaki Pendaki

	database, _ := db.ConnDB()
	defer database.Close()

	if err := c.ShouldBind(&pendaki); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.Query("INSERT INTO selesai (grup, nama, jenis, nik, alamat, catatan) VALUES (?,?, ?, ?, ?, ?)", pendaki.Grup,
		pendaki.Nama, pendaki.Jenis, pendaki.Nik, pendaki.Alamat, pendaki.Catatan)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/selesai?status=error")
	} else {
		database.Exec("DELETE FROM pendaki WHERE id = ?", pendaki.Id)
		c.Redirect(http.StatusSeeOther, "/data-pendaki?status=success")
	}
	// fmt.Println("Update successful for ID:", pendaki.Id)
}

func EditPendaki(c *gin.Context) {
	id := c.Param("id") // Mendapatkan ID dari URL

	database, _ := db.ConnDB()
	defer database.Close()

	var pendaki Pendaki
	err := database.QueryRow("SELECT * FROM pendaki WHERE id = ?", id).Scan(&pendaki.Nom, &pendaki.Id, &pendaki.Grup, &pendaki.Nama, &pendaki.Jenis, &pendaki.Nik, &pendaki.Alamat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// fmt.Println("Data Pendaki:", pendaki.Nama)

	// Render template edit dengan data pendaki
	c.HTML(http.StatusOK, "form-edit.html", gin.H{
		"title":   "Edit Pendaki",
		"pendaki": pendaki, // Tambahkan objek pendaki ke dalam data yang dikirimkan
	})
}

// Handler untuk update data pendaki
func UpdatePendaki(c *gin.Context) {
	var pendaki Pendaki

	database, _ := db.ConnDB()
	defer database.Close()

	if err := c.ShouldBind(&pendaki); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.Exec("UPDATE pendaki SET grup=?, nama=?, jenis=?, nik=?, alamat=? WHERE id=?", pendaki.Grup, pendaki.Nama, pendaki.Jenis, pendaki.Nik, pendaki.Alamat, pendaki.Id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/data-pendaki?status=error")
	} else {
		c.Redirect(http.StatusSeeOther, "/data-pendaki?status=success")
	}

	// fmt.Println("Update successful for ID:", pendaki.Id)

}

// untuk form pencarian
func CariData(c *gin.Context) {
	// Mendapatkan nilai dari query parameter "cari"
	query := c.Query("cari")

	database, _ := db.ConnDB()
	defer database.Close()

	var results []Pendaki

	rows, err := database.Query("SELECT * FROM pendaki WHERE nama LIKE ?", "%"+query+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var data Pendaki
		if err := rows.Scan(&data.Nom, &data.Id, &data.Grup, &data.Nama, &data.Jenis, &data.Nik, &data.Alamat); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, data)
	}

	c.JSON(http.StatusOK, gin.H{"pendaki": results})
}

// untuk button hapus
func HapusPendaki(c *gin.Context) {

	// Dapatkan nilai grup dari form
	id := c.PostForm("id")

	database, _ := db.ConnDB()
	defer database.Close()

	_, err := database.Exec("DELETE FROM pendaki WHERE id = ?", id)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/data-pendaki?status=error")
	} else {
		c.Redirect(http.StatusSeeOther, "/data-pendaki?status=success")
	}

	// Mengirim respons JSON
	// c.JSON(http.StatusOK, gin.H{"success": true, "message": "Data berhasil dihapus"})

}
