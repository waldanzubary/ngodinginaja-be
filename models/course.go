package models

import "time"

type Course struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`                      // nullable
	Order       int       `gorm:"default:0"`
	Language string 	  `gorm:"type:varchar(255);not null"`
	Attachment  string    `gorm:"type:varchar(255)"`

	Modules     []Module  `gorm:"foreignKey:CourseID;references:ID;constraint:OnDelete:CASCADE"`         // relasi ke modules
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
