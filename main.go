package main

import (
	"checkpoint-go/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	//parameter pertama untuk path url.parameter kedua untuk path lokal
	router.Static("/css", "assets/css")
	router.Static("/script", "assets/script")
	router.Static("/images", "assets/images")

	handlers.SetupRoutes(router)

	// Start server
	router.Run(":8080")
}
