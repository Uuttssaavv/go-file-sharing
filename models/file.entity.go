package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type FileModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Type      string     `gorm:"not null" json:"type"`
	Name      string     `gorm:"not null" json:"name"`
	Url       string     `gorm:"not null" json:"url"`
	AccessKey string     `gorm:"" json:"access_key"`
	CreatedAt time.Time  `gorm:"" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"" json:"updatedAt"`
	UserID    uint       `gorm:"foreignkey:UserID" json:"-"`
	User      UserEntity `gorm:"foreignkey:UserID" json:"user"`
}


func (entity *FileModel) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	entity.UpdatedAt = time.Now().Local()

	return nil
}
