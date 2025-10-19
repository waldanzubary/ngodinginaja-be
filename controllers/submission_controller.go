package controllers

import (
	"net/http"
	"ngodinginaja-be/config"
	"ngodinginaja-be/models"

	"github.com/gin-gonic/gin"
)

func GetSubmission(c *gin.Context) {
	lessonID := c.Param("lesson_id")

	var submissions []models.Submission
	if err := config.DB.Where("lesson_id = ?", lessonID).Preload("User").Find(&submissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(submissions) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No submissions found for this lesson"})
		return
	}

	c.HTML(http.StatusOK, "submission.html", gin.H{
		"title":       "Submissions",
		"submissions": submissions,
	})
}

func CreateSubmission(c *gin.Context) {
	var submission models.Submission

	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var lesson models.Lesson
	if err := config.DB.First(&lesson, submission.LessonID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	submission.IsCompleted = false

	inputMatch := lesson.Input != nil && submission.Code == *lesson.Input
	resultMatch := lesson.ExpectedOutput != nil && submission.Result != nil && *submission.Result == *lesson.ExpectedOutput

	if inputMatch && resultMatch {
		submission.IsCompleted = true
	}

	if err := config.DB.Create(&submission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"is_completed": submission.IsCompleted,
		"data":         submission,
	})
}
