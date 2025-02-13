package controllers

import (
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateTarget(c *gin.Context) {
	id := c.Param("id")
	target, found := FetchTargetByID(c, id)
	if !found {
		return
	}

	missionIDStr := strconv.Itoa(int(target.MissionID))
	mission, found := FetchMissionByID(c, missionIDStr)
	if !found {
		return
	}
	if mission.Complete || target.Complete {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot update completed mission or target"})
		return
	}

	var updatedTarget models.Target
	if HandleBadRequest(c, c.ShouldBindJSON(&updatedTarget)) {
		return
	}

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if updatedTarget.Complete {
			target.Complete = true
		}
		if updatedTarget.Notes != "" {
			target.Notes = updatedTarget.Notes
		}
		if err := tx.Model(&target).Where("complete = ?", false).Updates(target).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, target)
}

func DeleteTarget(c *gin.Context) {
	id := c.Param("id")
	target, found := FetchTargetByID(c, id)
	if !found {
		return
	}
	if target.Complete {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete completed target"})
		return
	}
	models.DB.Delete(target)
	c.JSON(http.StatusNoContent, nil)
}
func AddTarget(c *gin.Context) {
	missionID := c.Param("mission_id")
	mission, found := FetchMissionByID(c, missionID)
	if !found {
		return
	}
	if mission.Complete {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot add target to completed mission"})
		return
	}

	var targetCount int64
	models.DB.Model(&models.Target{}).Where("mission_id = ?", mission.ID).Count(&targetCount)
	if targetCount >= 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot add more than 3 targets to a mission"})
		return
	}

	var target models.Target
	if HandleBadRequest(c, c.ShouldBindJSON(&target)) {
		return
	}
	target.MissionID = mission.ID
	models.DB.Create(&target)
	c.JSON(http.StatusCreated, target)
}

func MarkTargetAsComplete(c *gin.Context) {
	id := c.Param("id")
	var target models.Target
	if err := models.DB.First(&target, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Target not found"})
		return
	}

	if target.Complete {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Target is already completed"})
		return
	}

	target.Complete = true
	if err := models.DB.Save(&target).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, target)
}
