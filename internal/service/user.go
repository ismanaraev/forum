package service

import (
	"forum/internal/models"
)

type UserRepository interface {
	AddUser(Username, Email, Password string) error
	GetUserById(Id string) (models.User, error)
	CheckToken(Id string, Token string) error
}

type userService struct {
	repo UserRepository
}

func newUserService(repo UserRepository) userService {
	return userService{repo: repo}
}

func (u userService) RegisterUser(Username, Email, Password string) (err error) {
	return nil
}

func (u userService) GetUser(Username, Password string) (user models.User, err error) {
	return models.User{}, nil
}

func (u userService) LogOut(Token string) error {
	return nil
}

func (u userService) CheckToken(Token string) error {
	return nil
}
