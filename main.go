package main

import (
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/models"
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/router"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Set Gin mode based on environment
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	// Database connection
	db, err := models.SetupDatabase()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Setup router
	r := router.SetupRouter()

	// Set trusted proxies
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal(err)
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
