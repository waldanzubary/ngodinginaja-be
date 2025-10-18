package controllers

import (
	"net/http"
	"ngodinginaja-be/config"
	"ngodinginaja-be/models"

	"github.com/gin-gonic/gin"
)

func GetModule(c *gin.Context) {
	courseID := c.Param("id")

	var modules []models.Module
	if err := config.DB.Where("course_id = ?", courseID).Find(&modules).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
		
	}
	c.JSON(http.StatusOK, modules)

}

func CreateModule(c *gin.Context) {
	var input struct {
		CourseID    uint   `json:"course_id" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Attachment  string `json:"attachment"`
		IsLocked    bool   `json:"is_locked"`
		Order       int    `json:"order"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	var course models.Course
	if err := config.DB.First(&course, input.CourseID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course not found"})
		return
	}

	module := models.Module{
		CourseID:    input.CourseID,
		Title:       input.Title,
		Description: input.Description,
		Attachment:  input.Attachment,
		IsLocked:    input.IsLocked,
		Order:       input.Order,
	}

	if err := config.DB.Create(&module).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Module created successfully", "module": module})
}
