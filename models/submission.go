package models

import (
	"time"
)

type Submission struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	User        *User     `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	LessonID    uint      `gorm:"not null;index" json:"lesson_id"`
	Lesson      *Lesson   `gorm:"foreignKey:LessonID;references:ID;constraint:OnDelete:CASCADE" json:"lesson,omitempty"`
	Code        string    `gorm:"type:text;not null" json:"code"`
	Result      *string   `gorm:"type:text" json:"result,omitempty"`
	IsCompleted bool      `gorm:"default:false" json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
