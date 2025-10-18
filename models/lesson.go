package models

import (
	"time"
	
)


type Difficult string

const (
	DifficultEasy   Difficult = "easy"
	DifficultNormal Difficult = "normal"
	DifficultHard   Difficult = "hard"
	DifficultExtreme Difficult = "extreme"
)



type Lesson struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	ModuleID        uint      `gorm:"not null;index"`
	Module          Module    `gorm:"foreignKey:ModuleID;constraint:OnDelete:CASCADE"`
	Title           string    `gorm:"type:varchar(255);not null"`
	Description     string    `gorm:"type:text;not null"`
	Difficult       Difficult    `gorm:"type:varchar(20);not null"`
	Input           *string   `gorm:"type:text"`
	ExpectedOutput  *string   `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
