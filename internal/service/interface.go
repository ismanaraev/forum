package service

import (
	"forum3/internal/models"
	"forum3/internal/repository"

	"github.com/gofrs/uuid"
)

type Service struct {
	Authorization
	Post
	Session
}

type Authorization interface {
	CreateSession(user models.Auth) (string, error)
	CreateUserService(user models.Auth) (int, error)
	AuthorizationUserService(models.Auth) (string, error)
}

type Post interface {
	GetPostService(post models.Post) (int, error)
	CreatePostService(post models.Post) (int, error)
	UpdatePostService(post models.Post) (int, error)
	DeletePostService(post models.Post) (int, error)
}

type Session interface {
	DeleteSessionRQtoRepo(tokenString string)
	GetSessionRQtoRepo(token string) (uuid.UUID, error)
}

func NewService(repo repository.Repository) Service {
	return Service{
		Authorization: NewAuthService(repo.Authorization),
		Session:       NewSessionService(repo.Session),
		Post:          NewPostService(repo.Post),
	}
}
