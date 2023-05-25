package models

import "time"


type UserEntity struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username;unique;not null"`
	Email        string  `gorm:"column:email;unique;not null"`
	Image        *string `gorm:"column:image"`
	Password string  `gorm:"column:password;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}