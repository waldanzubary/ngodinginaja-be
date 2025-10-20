package controllers

import (
	"net/http"
	"ngodinginaja-be/config"
	"ngodinginaja-be/models"
	"ngodinginaja-be/utils"

	"github.com/gin-gonic/gin"
)

	func GetCourse(c *gin.Context) {
		var courses []models.Course
		if err := config.DB.Preload("Modules").Find(&courses).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
			
		}
		c.JSON(http.StatusOK, courses)
	}

func CreateCourse(c *gin.Context) {
	var input models.Course

	input.Title = c.PostForm("title")
	input.Description = c.PostForm("description")
	input.Language = c.PostForm("language")


	file, fileHeader, err := c.Request.FormFile("attachment")
	if err == nil {
		defer file.Close()

		
		if fileHeader.Header.Get("Content-Type") != "image/jpeg" &&
			fileHeader.Header.Get("Content-Type") != "image/png" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Only JPG or PNG images allowed"})
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
		"message": "Course created successfully",
		"data":    input,
	})
}


