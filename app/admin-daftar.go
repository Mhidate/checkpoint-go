package app

import (
	db "checkpoint-go/db/config"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gelang struct {
	Id string `json:"id"`
}

type Pendaftar struct {
	Id     string `form:"id"`
	Nama   string `form:"nama"`
	Grup   int    `form:"grup"`
	Jenis  string `form:"jenis"`
	Nik    int    `form:"nik"`
	Alamat string `form:"alamat"`
}

var templat map[string]*template.Template

func init() {
	// Inisialisasi map templat
	templat = make(map[string]*template.Template)
	templat["admin-daftar.html"] = template.Must(template.ParseFiles("templates/admin-daftar.html", "templates/base-admin.html"))
	templat["admin-dasboard.html"] = template.Must(template.ParseFiles("templates/admin-dasboard.html", "templates/base-admin.html"))

}

func SimpanGelang(c *gin.Context) {
	idgelang := c.Param("idgelang")

	database, _ := db.ConnDB()
	defer database.Close()

	database.Query("DELETE FROM gelang")

	_, err := database.Query("INSERT INTO gelang (id) VALUES (?)", idgelang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Memberikan respons sukses jika penyimpanan berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil disimpan"})
}

func TampilkanForm(c *gin.Context) {
	// Mengambil nilai dari database
	database, _ := db.ConnDB()
	defer database.Close()

	var idgelang string
	err := database.QueryRow("SELECT * FROM gelang ").Scan(&idgelang)
	if err != nil {
		idgelang = ""
	}

	renderDaftar(c, "admin-daftar.html", gin.H{
		"Title":    "Halaman Daftar",
		"idgelang": idgelang,
	})
}

func SimpanPendaki(c *gin.Context) {
	// Membuat instance struct Pendaki untuk menyimpan data dari form
	var pendaki Pendaftar

	// Mem-parsing data dari form ke struct Pendaki
	if err := c.ShouldBind(&pendaki); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database, _ := db.ConnDB()
	defer database.Close()

	_, err := database.Query("INSERT INTO pendaki (id, nama, grup, jenis, nik, alamat) VALUES (?,?, ?, ?, ?, ?)", pendaki.Id,
		pendaki.Nama, pendaki.Grup, pendaki.Jenis, pendaki.Nik, pendaki.Alamat)
	if err != nil {
		database.Query("DELETE FROM gelang")
		c.Redirect(http.StatusSeeOther, "/daftar?status=error")
	} else {
		database.Query("DELETE FROM gelang")
		c.Redirect(http.StatusSeeOther, "/daftar?status=success")
	}
	// c.JSON(http.StatusInternalServerError, gin.H{"success": true, "message": "Berhasil mendaftar"})

	// // Memberikan respons sukses jika penyimpanan berhasil
	// c.Redirect(http.StatusSeeOther, "/daftar")
}

func renderDaftar(c *gin.Context, name string, data gin.H) {
	tmpl, ok := templat[name]
	if !ok {
		c.String(http.StatusNotFound, "Template tidak ada")
		return
	}

	err := tmpl.ExecuteTemplate(c.Writer, "base-admin.html", data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
