package controllers

import (
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func AssignCatToMission(c *gin.Context) {
	missionID := c.Param("mission_id")
	var cat models.Cat
	if HandleBadRequest(c, c.ShouldBindJSON(&cat)) {
		return
	}

	var responseMessage string
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE").Error; err != nil {
			return err
		}
		var mission models.Mission
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&mission, missionID).Error; err != nil {
			return err
		}
		if mission.CatID != 0 {
			responseMessage = "Mission already has a cat assigned"
			return nil
		}

		var existingMission models.Mission
		if err := tx.Where("cat_id = ? AND complete = ?", cat.ID, false).First(&existingMission).Error; err == nil {
			responseMessage = "Cat is already assigned to another incomplete mission"
			return nil
		}

		mission.CatID = cat.ID
		if err := tx.Save(&mission).Error; err != nil {
			return err
		}
		responseMessage = "Cat assigned to mission"
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if responseMessage == "Cat assigned to mission" {
		c.JSON(http.StatusOK, gin.H{"message": responseMessage})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": responseMessage})
	}
}

func CreateMission(c *gin.Context) {
	var mission models.Mission
	if HandleBadRequest(c, c.ShouldBindJSON(&mission)) {
		return
	}

	if mission.CatID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot set cat_id during mission creation"})
		return
	}

	for i := range mission.Targets {
		mission.Targets[i].MissionID = mission.ID
	}

	if err := models.DB.Create(&mission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mission)
}

func GetMissions(c *gin.Context) {
	var missions []models.Mission
	models.DB.Preload("Targets").Find(&missions)
	c.JSON(http.StatusOK, missions)
}

func GetMission(c *gin.Context) {
	id := c.Param("id")
	mission, found := FetchMissionByID(c, id)
	if !found {
		return
	}
	c.JSON(http.StatusOK, mission)
}

func UpdateMission(c *gin.Context) {
	id := c.Param("id")
	mission, found := FetchMissionByID(c, id)
	if !found {
		return
	}

	var updatedMission models.Mission
	if HandleBadRequest(c, c.ShouldBindJSON(&updatedMission)) {
		return
	}

	if mission.CatID != 0 && updatedMission.CatID != 0 && mission.CatID != updatedMission.CatID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot update cat_id once it is set"})
		return
	}

	if updatedMission.Complete {
		mission.Complete = true
	}

	models.DB.Save(&mission)
	c.JSON(http.StatusOK, mission)
}

func DeleteMission(c *gin.Context) {
	id := c.Param("id")
	var mission models.Mission
	if err := models.DB.First(&mission, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
		return
	}
	if mission.CatID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Mission is assigned to a cat"})
		return
	}

	if err := models.DB.Where("mission_id = ?", id).Delete(&models.Target{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Delete(&mission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
