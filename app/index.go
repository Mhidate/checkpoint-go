package app

import (
	"html"
	"html/template"
	"net/http"

	db "checkpoint-go/db/config"

	"github.com/gin-gonic/gin"
)

type Indec struct {
	Nomer  int    `json:"nomer"`
	ID     string `json:"id"`
	Grup   int    `json:"grup"`
	Nama   string `json:"nama"`
	Jenis  string `json:"jenis"`
	Nik    int    `json:"nik"`
	Alamat string `json:"alamat"`
	Pos    string `json:"pos"`
	Waktu  string `json:"waktu"`
}

var templates map[string]*template.Template

func init() {
	// Inisialisasi template
	templates = make(map[string]*template.Template)
	templates["index.html"] = template.Must(template.ParseFiles("templates/index.html", "templates/base-index.html"))

}

func IndecData(c *gin.Context) {

	database, err := db.ConnDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer database.Close()

	rows, err := database.Query("SELECT nomer, id, grup, nama, jenis, nik, alamat, pos, waktu FROM catat")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Membuat slice untuk menyimpan hasil query
	var results []Indec

	// Iterasi melalui setiap baris hasil query
	for rows.Next() {
		// Membuat variabel untuk menampung data dari setiap baris
		var data Indec
		var waktu []uint8

		// Membaca data dari setiap kolom pada baris saat ini
		if err := rows.Scan(&data.Nomer, &data.ID, &data.Grup, &data.Nama, &data.Jenis, &data.Nik, &data.Alamat, &data.Pos, &waktu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Mengkonversi data waktu menjadi string
		data.Waktu = string(waktu)

		// Menambahkan data ke dalam slice results
		results = append(results, data)

	}

	// Mengembalikan data dalam bentuk HTML menggunakan template
	// c.HTML(http.StatusOK, "index.html", gin.H{
	// 	"Title":"Beranda",
	// 	"results": results,
	// })

	renderIndex(c, "index.html", gin.H{
		"Title":   "Beranda",
		"results": results,
	})
}

func IndecSearch(c *gin.Context) {

	query := c.Query("query")

	// contoh input sanitization,untuk menghindari sql injection juga
	safeQuery := html.EscapeString(query)

	database, err := db.ConnDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer database.Close()

	// contoh parameterized query,salah satu untuk menghindari sql injection
	rows, err := database.Query("SELECT nomer, id, grup, nama, jenis, nik, alamat, pos, waktu FROM catat WHERE nama LIKE ?", "%"+safeQuery+"%")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Membuat slice untuk menyimpan hasil query
	var results []Indec

	// Iterasi melalui setiap baris hasil query
	for rows.Next() {
		// Membuat variabel untuk menampung data dari setiap baris
		var data Indec
		var waktu []uint8

		// Membaca data dari setiap kolom pada baris saat ini
		if err := rows.Scan(&data.Nomer, &data.ID, &data.Grup, &data.Nama, &data.Jenis, &data.Nik, &data.Alamat, &data.Pos, &waktu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Mengkonversi data waktu menjadi string
		data.Waktu = string(waktu)

		// Menambahkan data ke dalam slice results
		results = append(results, data)
	}

	// Mengembalikan data dalam bentuk HTML menggunakan template
	renderIndex(c, "index.html", gin.H{
		"Title":   "Beranda",
		"results": results,
	})

}

func renderIndex(c *gin.Context, name string, data gin.H) {
	tmpl, ok := templates[name]
	if !ok {
		c.String(http.StatusNotFound, "Template not found")
		return
	}

	err := tmpl.ExecuteTemplate(c.Writer, "base-index.html", data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
