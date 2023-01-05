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
	GetUserInfoService(user models.Auth) (models.Auth, error)
	GetUsersInfoByUUIDtoRepo(id uuid.UUID) (models.Auth, error)
}

type Post interface {
	GetAllPostService() ([]models.Post, error)
	CreatePostService(post models.Post) (int, error)
	UpdatePostService(post models.Post) (int, error)
	DeletePostService(post models.Post) (int, error)
}

type Session interface {
	DeleteSessionRQtoRepo(uuid.UUID) error
	GetSessionRQtoRepo(token string) (uuid.UUID, error)
}

func NewService(repo repository.Repository) Service {
	return Service{
		Authorization: NewAuthService(repo.Authorization),
		Session:       NewSessionService(repo.Session),
		Post:          NewPostService(repo.Post),
	}
}
