package models

import (
	"github.com/jinzhu/gorm"
	"go-crud/utils"
	"time"
)

type UserEntity struct {
	ID        uint    `gorm:"primary_key"`
	Username  string  `gorm:"column:username;unique;not null"`
	Email     string  `gorm:"column:email;unique;not null"`
	Image     *string `gorm:"column:image"`
	Password  string  `gorm:"column:password;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *UserEntity) BeforeCreate(db *gorm.DB) error {
	entity.Password = utils.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}


func (entity *UserEntity) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}