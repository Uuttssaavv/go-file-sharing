package models

import "time"

type FileModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Type      string    `gorm:"not null" json:"type"`
	Name      string    `gorm:"not null" json:"name"`
	Url       string    `gorm:"not null" json:"url"`
	AccessKey string    `gorm:"" json:"access_key"`
	CreatedAt time.Time `gorm:"" json:"createdAt"`
	UpdatedAt time.Time `gorm:"" json:"updatedAt"`
}
