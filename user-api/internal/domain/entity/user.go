package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(20)"`
	Name     string `gorm:"type:varchar(20)"`
	ImageUrl string `gorm:"type:varchar(100)"`
	Token    string `gorm:"type:varchar(200)"`
}
