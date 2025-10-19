package controllers

import (
	"net/http"
	"ngodinginaja-be/config"
	"ngodinginaja-be/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetLesson(c* gin.Context) {
	moduleID := c.Param("id")
	
	var lessons []models.Lesson
	if  err := config.DB.Preload("Submissions").Where("module_id = ?", moduleID).Find(&lessons).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
		
	}
	c.JSON(http.StatusOK, lessons)

}


func CreateLesson(c *gin.Context) {
	var input struct {
		ModuleID       uint   `json:"module_id" binding:"required"`
		Title          string `json:"title" binding:"required"`
		Description    string `json:"description" binding:"required"`
		Difficult      string `json:"difficult" binding:"required"`
		Input          *string `json:"input"`
		ExpectedOutput *string `json:"expected_output"`
	}


	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	var module models.Module
	if err := config.DB.First(&module, input.ModuleID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Module not found"})
		return
	}

	
	validDifficulties := map[string]bool{
		"easy": true, "normal": true, "hard": true, "extreme": true,
	}
	if !validDifficulties[strings.ToLower(input.Difficult)] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid difficult level"})
		return
	}

	lesson := models.Lesson{
		ModuleID:       input.ModuleID,
		Title:          input.Title,
		Description:    input.Description,
		Difficult:      models.Difficult(strings.ToLower(input.Difficult)),
		Input:          input.Input,
		ExpectedOutput: input.ExpectedOutput,
	}

	if err := config.DB.Create(&lesson).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lesson created successfully",
		"lesson":  lesson,
	})
}
