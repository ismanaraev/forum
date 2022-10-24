package service

import "forum/internal/repository"

type Service struct {
	User UserService
}

func NewService(repository repository.Repository) Service {
	return Service{
		User: newUserService(repository.User),
	}
}
