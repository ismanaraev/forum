package repository

import (
	"database/sql"
	"forum3/internal/models"
	"time"

	"github.com/gofrs/uuid"
)

type Repository struct {
	Authorization
	Post
	Session
}

type Authorization interface {
	SetSession(user models.Auth, token string, time time.Time) error
	CreateUser(user models.Auth) (int, error)
	GetUserInfo(user models.Auth) (models.Auth, error)
	GetUsersEmail(user models.Auth) (models.Auth, error)
}

type Post interface {
	GetPost(post models.Post) (int, error)
	CreatePost(post models.Post) (int, error)
	UpdatePost(post models.Post) (int, error)
	DeletePost(post models.Post) (int, error)
}

type Session interface {
	GetSessionFromDB(token string) (uuid.UUID, error)
	DeleteSessionFromDB(user models.Auth) (int, error)
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Authorization: NewAuthSQLite(db),
		Post:          NewPostSQLite(db),
		Session:       NewSessionSQLite(db),
	}
}
