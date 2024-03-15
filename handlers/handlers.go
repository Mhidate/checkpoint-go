package handlers

import (
	"checkpoint-go/app"
	"checkpoint-go/auth"

	// "fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var templates map[string]*template.Template

func init() {
	// Inisialisasi template
	templates = make(map[string]*template.Template)
	// templates["index.html"] = template.Must(template.ParseFiles("templates/index.html", "templates/base-index.html"))
	templates["page-map.html"] = template.Must(template.ParseFiles("templates/page-map.html", "templates/base-index.html"))
	templates["page-about.html"] = template.Must(template.ParseFiles("templates/page-about.html", "templates/base-index.html"))
	templates["admin-dasboard.html"] = template.Must(template.ParseFiles("templates/admin-dasboard.html", "templates/base-admin.html"))
	templates["admin-daftar.html"] = template.Must(template.ParseFiles("templates/admin-daftar.html", "templates/base-admin.html"))
	templates["admin-data-pendaki.html"] = template.Must(template.ParseFiles("templates/admin-data-pendaki.html", "templates/base-admin.html"))
	templates["admin-hapus-lokasi.html"] = template.Must(template.ParseFiles("templates/admin-hapus-lokasi.html", "templates/base-admin.html"))
	templates["admin-log-pendaki.html"] = template.Must(template.ParseFiles("templates/admin-log-pendaki.html", "templates/base-admin.html"))
	templates["admin-selesai.html"] = template.Must(template.ParseFiles("templates/admin-selesai.html", "templates/base-admin.html"))
	templates["admin-tambah-lokasi.html"] = template.Must(template.ParseFiles("templates/admin-tambah-lokasi.html"))
	templates["page-login.html"] = template.Must(template.ParseFiles("templates/page-login.html"))
	templates["form-edit.html"] = template.Must(template.ParseFiles("templates/form-edit.html"))
	templates["form-selesai.html"] = template.Must(template.ParseFiles("templates/form-selesai.html"))

	// templates["indec.html"] = template.Must(template.ParseFiles("templates/indec.html"))

}

func aboutHandler(c *gin.Context) {
	renderIndex(c, "page-about.html", gin.H{
		"Title": "Tentang",
	})
}

func mapHandler(c *gin.Context) {
	renderIndex(c, "page-map.html", gin.H{
		"Title": "Peta",
	})
}
func loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "page-login.html", gin.H{
		"Title":        "Login",
		"ErrorMessage": "",
	})
}

// func daftarHandler(c *gin.Context) {
// 	renderAdmin(c, "admin-daftar.html", gin.H{
// 		"Title": "Halaman Daftar",
// 	})
// }

func dataPendakiHandler(c *gin.Context) {
	renderAdmin(c, "admin-data-pendaki.html", gin.H{
		"Title": "Halaman Data Pendaki",
	})
}
func logPendakiHandler(c *gin.Context) {
	renderAdmin(c, "admin-log-pendaki.html", gin.H{
		"Title": "Halaman Log Pendaki",
	})
}
func selesaiHandler(c *gin.Context) {
	renderAdmin(c, "admin-selesai.html", gin.H{
		"Title": "Halaman Selesai",
	})
}

func hapusLokasiHandler(c *gin.Context) {
	renderAdmin(c, "admin-hapus-lokasi.html", gin.H{
		"Title": "Halaman Hapus Lokasi",
	})
}
func tambahLokasiHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-tambah-lokasi.html", gin.H{
		"Title": "Halaman Tambah Lokasi",
	})
}

// AdminPage handles the admin page
func adminPage(c *gin.Context) {
	// Retrieve username from cookie
	username, err := c.Cookie("username")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// Check if the user is authenticated
	if username == "mingo" {
		// Render the admin page
		renderAdmin(c, "admin-dasboard.html", gin.H{
			"Title": "Admin",
		})
		return
	}

	// If not authenticated, redirect to login page
	c.Redirect(http.StatusSeeOther, "/login")
}

func logoutHandler(c *gin.Context) {
	// Delete the cookie or session that stores the user information
	c.SetCookie("username", "", -1, "/", "", false, true)

	// Redirect to the login page after logout
	c.Redirect(http.StatusSeeOther, "/login")
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

func renderAdmin(c *gin.Context, name string, data gin.H) {
	tmpl, ok := templates[name]
	if !ok {
		c.String(http.StatusNotFound, "Template not found")
		return
	}

	err := tmpl.ExecuteTemplate(c.Writer, "base-admin.html", data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func SetupRoutes(router *gin.Engine) {
	// router.GET("/daftar", daftarHandler)
	// router.GET("/", homeHandler)
	//halaman
	router.GET("/tentang", aboutHandler)
	router.GET("/peta", mapHandler)
	router.GET("/login", loginHandler)
	router.GET("/admin", adminPage)
	router.GET("/daftar", app.TampilkanForm)
	router.GET("/data-pendaki", dataPendakiHandler)
	router.GET("/log-pendaki", logPendakiHandler)
	router.GET("/selesai", selesaiHandler)
	router.GET("/hapus-lokasi", hapusLokasiHandler)
	router.GET("/tambah-lokasi", tambahLokasiHandler)
	router.GET("/", app.IndecData)

	// endpoint macam-macam
	router.POST("/crx", auth.BuatAkun)
	router.POST("/proses", auth.DoLogin)
	router.GET("/logout", logoutHandler)

	router.GET("/tampil-marker", app.MarkerMap)
	router.POST("/add-marker", app.AddMarker)
	router.POST("/hapus-log", app.HapusLog)
	router.POST("/hapus-pos", app.HapusPos)
	router.GET("/jsun-pendaki", app.DataPendaki)
	router.GET("/cari-pendaki", app.CariData)
	router.POST("/hapus-pendaki", app.HapusPendaki)
	router.GET("/form-selesai/:id", app.ButtonSelesai)
	router.POST("/proses-selesai", app.ProsesSelesai)
	router.GET("/edit/:id", app.EditPendaki)
	router.POST("/update-pendaki", app.UpdatePendaki)
	router.GET("/Search", app.IndecSearch)
	router.GET("/form/:idgelang", app.SimpanGelang)
	router.POST("/simpan-pendaki", app.SimpanPendaki)
	router.GET("/ms-excel", app.FileExcel)
	router.GET("/hapus-selesai", app.HapusTabelSelesai)
	router.GET("/data-selesai", app.GetDataSelesai)

	//biodata dan log pendaki ,rencana endpoint untuk app mobile
	router.GET("/data", app.GetData)

}
