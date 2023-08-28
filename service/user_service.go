package service

import (
	"github.com/nicholascostadev/todo-backend/model"
)

type NewUserService struct{}

func (T *NewUserService) GetAllUsers() ([]model.User, error) {
	return model.GetAllUsers()
}
