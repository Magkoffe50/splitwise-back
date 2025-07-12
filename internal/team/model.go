package team

import (
	"splitbill-back/internal/user"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name        string
	Description string
	Users       []user.User `gorm:"many2many:user_teams;"`
}
