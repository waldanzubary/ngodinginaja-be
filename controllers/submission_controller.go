package controllers

import (
	"net/http"
	"ngodinginaja-be/config"
	"ngodinginaja-be/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSubmission(c *gin.Context) {
	lessonID := c.Param("id")

	var submissions []models.Submission
	if err := config.DB.Where("lesson_id = ?", lessonID).Find(&submissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(submissions) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No submissions found for this lesson"})
		return
	}

	c.JSON(http.StatusOK, submissions)
}





func UpdateSubmission(c *gin.Context) {
	idParam := c.Param("id")

	
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid submission ID"})
		return
	}

	
	userValue, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user := userValue.(models.User)


	var existingSubmission models.Submission
	if err := config.DB.First(&existingSubmission, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}

	
	if existingSubmission.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to edit this submission"})
		return
	}

	
	var input struct {
		Code   string  `json:"code"`
		Result *string `json:"result"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	existingSubmission.Code = input.Code
	existingSubmission.Result = input.Result

	
	var lesson models.Lesson
	if err := config.DB.First(&lesson, existingSubmission.LessonID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	
	existingSubmission.IsCompleted = false
	inputMatch := lesson.Input != nil && existingSubmission.Code == *lesson.Input
	resultMatch := lesson.ExpectedOutput != nil && existingSubmission.Result != nil && *existingSubmission.Result == *lesson.ExpectedOutput
	if inputMatch && resultMatch {
		existingSubmission.IsCompleted = true
	}

	if err := config.DB.Save(&existingSubmission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"message":      "Submission updated successfully",
		"is_completed": existingSubmission.IsCompleted,
		"data":         existingSubmission,
	})
}

func CreateSubmission(c *gin.Context) {
	var submission models.Submission

	
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	userValue, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	
	user := userValue.(models.User)
	submission.UserID = user.ID 

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

