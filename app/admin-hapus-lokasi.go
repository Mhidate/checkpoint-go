package app

import (
	db "checkpoint-go/db/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Location struct {
	ID         int    `json:"id"`
	Lat        string `json:"lat"`
	Longt      string `json:"longt"`
	NamaPos    string `json:"nama_pos"`
	Keterangan string `json:"keterangan"`
}

func MarkerMap(c *gin.Context) {
	database, _ := db.ConnDB()
	defer database.Close()

	rows, err := database.Query("SELECT * FROM lokasi")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data dari database"})
		return
	}
	defer rows.Close()

	var locations []Location

	for rows.Next() {
		var loc Location
		err := rows.Scan(&loc.ID, &loc.Lat, &loc.Longt, &loc.NamaPos, &loc.Keterangan)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data dari database"})
			return
		}
		locations = append(locations, loc)
	}

	c.JSON(http.StatusOK, gin.H{"locations": locations})
}

func HapusPos(c *gin.Context) {

	// nilai dari form
	id := c.PostForm("id")

	database, _ := db.ConnDB()
	defer database.Close()

	_, err := database.Exec("DELETE FROM lokasi WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal menghapus data pos dari tabel"})
		return
	}

	// Mengirim respons JSON
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Data pos berhasil dihapus"})

}
