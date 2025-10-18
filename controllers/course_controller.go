	package controllers

	import (
		"net/http"
		"ngodinginaja-be/config"
		"ngodinginaja-be/models"

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
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
			
		}

		course := models.Course{
			Title: input.Title,
			Description: input.Description,
			Language: input.Language,

		}

		if err := config.DB.Create(&course).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, course)


	}