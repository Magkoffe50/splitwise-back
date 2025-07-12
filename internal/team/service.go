package team

import (
	"splitbill-back/internal/db"
	"splitbill-back/internal/user"
)

func CreateTeam(name, desc string, userIDs []uint) (*Team, error) {
	var users []user.User
	if err := db.DB.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		return nil, err
	}

	team := &Team{Name: name, Description: desc, Users: users}
	if err := db.DB.Create(&team).Error; err != nil {
		return nil, err
	}

	return team, nil
}

func GetTeamByID(id string) (*Team, error) {
	var team Team
	if err := db.DB.Preload("Users").First(&team, id).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
