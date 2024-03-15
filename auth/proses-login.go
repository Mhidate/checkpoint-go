package auth

import (
	db "checkpoint-go/db/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func BuatAkun(c *gin.Context) {
	var admin Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input
	if strings.TrimSpace(admin.Username) == "" || strings.TrimSpace(admin.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username dan password tidak boleh kosong"})
		return
	}

	// Enkripsi password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}

	database, err := db.ConnDB()
	defer database.Close()

	_, err = database.Exec("INSERT INTO admins (username, password) VALUES (?, ?)", admin.Username, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data ke database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Akun admin berhasil dibuat"})
}

func DoLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.HTML(http.StatusBadRequest, "page-login.html", gin.H{
			"Title":        "Login",
			"ErrorMessage": "Username dan password tidak boleh kosong",
		})
		return
	}

	database, err := db.ConnDB()
	defer database.Close()

	var hashedPassword string
	err = database.QueryRow("SELECT password FROM admins WHERE username = ?", username).Scan(&hashedPassword)
	if err != nil {
		c.HTML(http.StatusBadRequest, "page-login.html", gin.H{
			"Title":        "Login",
			"ErrorMessage": "Username atau password salah",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		c.HTML(http.StatusBadRequest, "page-login.html", gin.H{
			"Title":        "Login",
			"ErrorMessage": "Username atau password salah",
		})
		return
	}

	c.SetCookie("username", username, 3600, "/", "", false, true)

	c.Redirect(http.StatusSeeOther, "/admin")
}
