package controllers

import (
	"net/http"
	"ngodinginaja-be/config"
	"ngodinginaja-be/models"
	"ngodinginaja-be/utils"

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
	var input models.Module

	input.Title = c.PostForm("title")
	input.Description = c.PostForm("description")

	courseID := c.PostForm("course_id")
	if courseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course_id is required"})
		return
	}

	
	var course models.Course
	if err := config.DB.First(&course, courseID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course_id, course not found"})
		return
	}
	input.CourseID = course.ID

	
	file, fileHeader, err := c.Request.FormFile("attachment")
	if err == nil {
		defer file.Close()

		contentType := fileHeader.Header.Get("Content-Type")
		if contentType != "image/jpeg" &&
			contentType != "image/png" &&
			contentType != "application/pdf" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Only JPG, PNG, or PDF allowed"})
			return
		}

		url, uploadErr := utils.UploadToCloudinary(file, fileHeader)
		if uploadErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": uploadErr.Error()})
			return
		}

		input.Attachment = url
	}

	
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Module created successfully",
		"data":    input,
	})
}
