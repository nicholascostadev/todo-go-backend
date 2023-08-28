package model

import "errors"

type User struct {
	ID           uint   `json:"id"           gorm:"primaryKey"`
	Username     string `json:"username"     gorm:"unique;not null"`
	HashPassword string `json:"hashPassword" gorm:"not null"`
}

func GetAllUsers() ([]User, error) {
	var users []User
	tx := db.Find(&users)

	if tx.Error != nil {
		return []User{}, tx.Error
	}

	return users, nil
}

func RegisterUser(user User) (User, error) {
	tx := db.Create(&user)

	return user, tx.Error
}

func GetUserByUsername(username string) (User, error) {
	var foundUser User

	db.Where("username = ?", username).First(&foundUser)

	return foundUser, nil
}

func LoginUser(user User) (User, error) {
	var foundUser User
	db.Where("username = ?", user.Username).First(&foundUser)

	if foundUser.Username != user.Username {
		return User{}, errors.New("No user found with given username, maybe you're trying to login with a different username?")
	}

	return User{
		ID:       foundUser.ID,
		Username: foundUser.Username,
	}, nil
}
