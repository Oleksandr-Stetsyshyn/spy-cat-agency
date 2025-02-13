package controllers

import (
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/models"
	"github.com/Oleksandr-Stetsyshyn/spy-cat-agency/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCat(c *gin.Context) {
	var cat models.Cat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !services.IsValidBreed(cat.Breed) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breed"})
		return
	}
	models.DB.Create(&cat)
	c.JSON(http.StatusCreated, cat)
}

func UpdateCat(c *gin.Context) {
	id := c.Param("id")
	var cat models.Cat
	if err := models.DB.First(&cat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cat not found"})
		return
	}
	var updatedCat models.Cat
	if err := c.ShouldBindJSON(&updatedCat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !services.IsValidBreed(updatedCat.Breed) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid breed"})
		return
	}
	cat.Name = updatedCat.Name
	cat.YearsOfExperience = updatedCat.YearsOfExperience
	cat.Breed = updatedCat.Breed
	cat.Salary = updatedCat.Salary
	models.DB.Save(&cat)
	c.JSON(http.StatusOK, cat)
}

func GetCats(c *gin.Context) {
	var cats []models.Cat
	if err := models.DB.Find(&cats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}

func GetCat(c *gin.Context) {
	id := c.Param("id")
	var cat models.Cat
	if err := models.DB.First(&cat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cat not found"})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func DeleteCat(c *gin.Context) {
	id := c.Param("id")
	var missions []models.Mission
	if err := models.DB.Where("cat_id = ?", id).Find(&missions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(missions) > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Cat has active missions"})
		return
	}
	if err := models.DB.Delete(&models.Cat{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cat deleted"})
}
