package controllers

import (
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FetchTargetByID(c *gin.Context, id string) (*models.Target, bool) {
	var target models.Target
	if err := models.DB.First(&target, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Target not found"})
		return nil, false
	}
	return &target, true
}

func FetchMissionByID(c *gin.Context, id string) (*models.Mission, bool) {
	var mission models.Mission
	if err := models.DB.Preload("Targets").First(&mission, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
		return nil, false
	}
	return &mission, true
}

func HandleBadRequest(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
	}
	return false
}
