package app

import (
	db "checkpoint-go/db/config"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type DataSelesai struct {
	Grup    int    `json:"grup"`
	Nama    string `json:"nama"`
	Jenis   string `json:"jenis"`
	Nik     int    `json:"nik"`
	Alamat  string `json:"alamat"`
	Tgl     string `json:"tgl"`
	Catatan string `json:"catatan"`
}

func FileExcel(c *gin.Context) {

	database, err := db.ConnDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	rows, err := database.Query("SELECT grup, nama, jenis, alamat, tgl, catatan FROM selesai")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Buat file Excel baru
	file := xlsx.NewFile()

	// Tambahkan sheet
	sheet, err := file.AddSheet("Data Pendaki")
	if err != nil {
		log.Fatal(err)
	}
	//nama heade
	headerRow := sheet.AddRow()
	headerRow.AddCell().SetValue("No")
	headerRow.AddCell().SetValue("Nama")
	headerRow.AddCell().SetValue("Jenis")
	headerRow.AddCell().SetValue("Alamat")
	headerRow.AddCell().SetValue("Grup")
	headerRow.AddCell().SetValue("Tanggal Selesai")
	headerRow.AddCell().SetValue("Catatan")

	// Tambahkan data dari database ke tabel
	no := 1
	for rows.Next() {
		var nama, jenis, alamat, tgl, catatan, grup string
		if err := rows.Scan(&grup, &nama, &jenis, &alamat, &tgl, &catatan); err != nil {
			log.Fatal(err)
		}
		row := sheet.AddRow()
		row.AddCell().SetValue(fmt.Sprintf("%d", no))
		no++
		row.AddCell().SetValue(nama)
		row.AddCell().SetValue(jenis)
		row.AddCell().SetValue(alamat)
		row.AddCell().SetValue(grup)
		row.AddCell().SetValue(tgl)
		row.AddCell().SetValue(catatan)

		// Mengukur lebar konten di setiap kolom
		// contentWidth := []float64{0, 0, 0, 0, 0, 0} // inisialisasi dengan nilai 0
		// contentWidth[0] = float64(len(nama))
		// contentWidth[1] = float64(len(jenis))
		// contentWidth[2] = float64(len(alamat))
		// contentWidth[3] = float64(len(grup))
		// contentWidth[4] = float64(len(tgl))
		// contentWidth[5] = float64(len(catatan))

		// Mengatur lebar kolom berdasarkan lebar konten
		// new := 0.0
		// for _, width := range contentWidth {

		// 	new = new + float64(width)
		// 	// sheet.SetColWidth(i, i, new) // Mengatur lebar kolom
		// }

		sheet.SetColWidth(5, 6, 20.0)
	}

	fileFormat := "2006_01_02" // Format tanggal: tahun-bulan-hari
	filename := "data_pendaki_selesai_" + time.Now().Format(fileFormat) + ".xlsx"
	err = file.Save(filename)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/selesai?status=error")
	} else {
		c.Redirect(http.StatusSeeOther, "/selesai?status=success")
	}

	// fmt.Println("Dokumen Excel berhasil disimpan.")
}

func GetDataSelesai(c *gin.Context) {

	database, err := db.ConnDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	defer database.Close()

	rows, err := database.Query("SELECT grup, nama, jenis,nik, alamat, tgl, catatan FROM selesai")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var hasilselesai []DataSelesai

	for rows.Next() {
		var data DataSelesai

		if err := rows.Scan(&data.Grup, &data.Nama, &data.Jenis, &data.Nik, &data.Alamat, &data.Tgl, &data.Catatan); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		hasilselesai = append(hasilselesai, data)
	}

	c.JSON(http.StatusOK, hasilselesai)

}

func HapusTabelSelesai(c *gin.Context) {

	database, _ := db.ConnDB()
	defer database.Close()

	_, err := database.Exec("DELETE FROM selesai ")
	if err != nil {
		c.String(http.StatusInternalServerError, "Gagal menghapus data tabel")
		return
	}

	// Mengalihkan (redirect) ke halaman setelah berhasil
	c.Redirect(http.StatusSeeOther, "/selesai")

}
