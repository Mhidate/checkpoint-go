package app

import (
	db "checkpoint-go/db/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HapusLog(c *gin.Context) {

	grup := c.PostForm("grup")

	database, _ := db.ConnDB()
	defer database.Close()

	_, err := database.Exec("DELETE FROM catat WHERE grup = ?", grup)
	if err != nil {
		c.String(http.StatusInternalServerError, "Gagal menghapus data dari tabel")
		return
	}

	// Mengalihkan (redirect) ke halaman setelah berhasil
	c.Redirect(http.StatusSeeOther, "/log-pendaki")

}
