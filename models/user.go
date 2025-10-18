package models

import (
    "time"
    "gorm.io/gorm"
)


type Role string

const (
  	RoleAdmin   Role = "admin"
    RoleInstruction Role = "instruction"
    RoleUser    Role = "user"
)


type Plan string

const (
    Free    Plan = "free"
    Premium Plan = "premium"
)


type User struct {
    ID             uint       `gorm:"primaryKey"`
    FirebaseUID    *string    `gorm:"unique"`                         
    Username       string     `gorm:"type:varchar(100);not null"`
    Email          string     `gorm:"unique;not null"`
    Password       string     `gorm:"not null"`
    Role           Role       `gorm:"type:varchar(20);default:'user'"`
    Plan           Plan       `gorm:"type:varchar(20);default:'free'"`
    ProfilePicture string     `gorm:"type:varchar(255);null"`
    Bio            string     `gorm:"type:varchar(500);null"`
	VerificationCode string   `gorm:"type:varchar(255);null"`
	IsVerified   bool      	  `gorm:"default:false"`
    Submission []Submission   `gorm:"foreignKey:UserID"`
    CreatedAt      time.Time
    UpdatedAt      time.Time
    DeletedAt      gorm.DeletedAt `gorm:"index"` 
}
