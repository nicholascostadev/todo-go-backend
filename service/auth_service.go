package service

import (
	"errors"

	"github.com/nicholascostadev/todo-backend/model"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type NewAuthService struct{}

type RegisterUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (T *NewAuthService) RegisterUser(user RegisterUserData) (model.User, error) {
	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}

	newUser, err := model.RegisterUser(model.User{
		Username:     user.Username,
		HashPassword: hashPassword,
	})
	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

type LoginUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (T *NewAuthService) LoginUser(user LoginUserData) (model.Session, error) {
	foundUser, err := model.GetUserByUsername(user.Username)
	if err != nil {
		return model.Session{}, errors.New(
			"No user found with given username, maybe you're trying to login with a different username?",
		)
	}

	hasCorrectPassword := CheckPasswordHash(user.Password, foundUser.HashPassword)

	if !hasCorrectPassword {
		return model.Session{}, errors.New("Passwords don't match")
	}

	session, err := model.CreateSession(model.CreateSessionData{
		ID:       foundUser.ID,
		Username: foundUser.Username,
	})
	if err != nil {
		return model.Session{}, err
	}

	return model.Session{
		ID:        session.ID,
		UserID:    session.UserID,
		Token:     session.Token,
		CreatedAt: session.CreatedAt,
		ExpiresAt: session.ExpiresAt,
		UpdatedAt: session.UpdatedAt,
	}, nil
}

func (T *NewAuthService) GetSessionById(id uint) (model.Session, error) {
	session, err := model.GetSessionById(id)
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}
