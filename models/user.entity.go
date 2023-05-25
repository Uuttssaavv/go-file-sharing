package models

import "time"


type UserEntity struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:usernameunique;not null"`
	Email        string  `gorm:"column:email;unique;not null"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
	AccessKey string  `gorm:"column:access_key;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}