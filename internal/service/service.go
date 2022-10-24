package service

import "forum/internal/repository"

type Service struct {
	User UserService
}

type UserService interface {
	RegisterUser(Username, Email, Password string) error
	Authorize(Username, Password string) error
	LogOut(Token string) error
}

func NewService(repository repository.Repository) Service {
	return Service{
		User: newUserService(repository.User),
	}
}
