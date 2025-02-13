package controllers

import (
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetItems(c *gin.Context) {
	var items []models.Item
	models.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&item)
	c.JSON(http.StatusCreated, item)
}
