package filecontrollers

type FileInput struct{
	ID uint `gorm:"primary_key"`
	Type string `gorm:"not null"`
	Name string `gorm:"not null"`
	Url string `gorm:"not null"`
}