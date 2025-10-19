package models

import (
	"time"
)

type Submission struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null;index"`
	User        *User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	LessonID    uint      `gorm:"not null;index"`
	Lesson      *Lesson    `gorm:"foreignKey:LessonID;references:ID;constraint:OnDelete:CASCADE"`
	Code        string    `gorm:"type:text;not null"`
	Result      *string   `gorm:"type:text"`
	IsCompleted bool      `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
