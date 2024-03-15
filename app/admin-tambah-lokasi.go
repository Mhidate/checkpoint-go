package app

import (
	db "checkpoint-go/db/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Locmarker struct {
	ID         int    `json:"id"`
	Lat        string `json:"latitude"`
	Longt      string `json:"longitude"`
	NamaPos    string `json:"nama_pos"`
	Keterangan string `json:"keterangan"`
}

func AddMarker(c *gin.Context) {
	// database, err := db.ConnDB()

	database, _ := db.ConnDB()
	defer database.Close()

	var locMarker Locmarker
	if err := c.ShouldBindJSON(&locMarker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	ad, err := database.Query("INSERT INTO lokasi (lat, longt, nama_pos, keterangan) VALUES (?, ?, ?, ?)", locMarker.Lat, locMarker.Longt, locMarker.NamaPos, locMarker.Keterangan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare statement"})
		return
	}
	defer ad.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil menambah penanda pos", "location": locMarker})
}
