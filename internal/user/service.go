package user

import (
	"splitbill-back/internal/db"
)

func CreateUser(u *User) error {
	return db.DB.Create(u).Error
}

func GetUserByEmailAndPassword(email, password string) (*User, error) {
	var user User
	if err := db.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsersByQuery(email, login string) ([]User, error) {
	var users []User
	query := db.DB

	if email != "" {
		query = query.Where("email = ?", email)
	}
	if login != "" {
		query = query.Or("login = ?", login)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
