package filecontrollers

type FileInput struct {
	ID     string `gorm:"primary_key"`
	Type   string `gorm:"not null"`
	Name   string `gorm:"not null"`
	Url    string `gorm:"not null"`
	UserId uint   `gorm:"not null"`
}
