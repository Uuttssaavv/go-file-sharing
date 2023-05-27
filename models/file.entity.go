package models

import "time"

type FileModel struct{
	ID uint `gorm:"primary_key"`
	Type string `gorm:"not null"`
	Name string `gorm:"not null"`
	Url string `gorm:"not null"`
	AccessKey string `gorm:""`
	CreatedAt time.Time
	UpdatedAt time.Time
}