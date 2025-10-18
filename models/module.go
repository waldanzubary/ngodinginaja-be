package models

import "time"

type Module struct {
	ID          uint      `gorm:"primaryKey"`
	CourseID    uint      `gorm:"not null;index"`
	Course      *Course   `gorm:"foreignKey:CourseID;references:ID;constraint:OnDelete:CASCADE"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Attachment  string    `gorm:"type:varchar(255)"`
	IsLocked    bool      `gorm:"default:false"`
	Order       int       `gorm:"default:0"`
	Lessons     []Lesson  `gorm:"foreignKey:ModuleID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
