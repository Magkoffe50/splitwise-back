package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Login    string `gorm:"unique"`
	Password string
}
