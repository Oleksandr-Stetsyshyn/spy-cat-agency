package main

import (
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/controllers"
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/middleware"
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	// Middleware
	r.Use(middleware.Logger())

	// Database connection
	db, err := models.SetupDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Routes
	r.GET("/items", controllers.GetItems)
	r.POST("/items", controllers.CreateItem)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
