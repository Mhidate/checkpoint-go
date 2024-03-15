package app

import (
	db "checkpoint-go/db/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
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

func GetData(c *gin.Context) {

	database, err := db.ConnDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	defer database.Close()

	rows, err := database.Query("SELECT * FROM catat")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var results []Data

	for rows.Next() {
		var data Data
		var waktu []uint8
		if err := rows.Scan(&data.Nomer, &data.ID, &data.Grup, &data.Nama, &data.Jenis, &data.Nik, &data.Alamat, &data.Pos, &waktu); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		data.Waktu = string(waktu)
		results = append(results, data)
	}

	c.JSON(http.StatusOK, results)
	// c.HTML(http.StatusOK, "index.html", gin.H{
	// 	"results": results,
	// })
}
