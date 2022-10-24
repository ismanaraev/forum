package repository

import (
	"database/sql"
	"forum/internal/models"
)

type UserRepository interface {
	AddUser(Username, Email, Password string) error
	GetUser(Id string) (models.User, error)
	CheckToken(Id string, Token string) error
}

func newUserRepo(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

type userRepo struct {
	db *sql.DB
}

func (u userRepo) AddUser(Username, Email, Password string) error {
	return nil
}

func (u userRepo) GetUser(Id string) (models.User, error) {
	return models.User{}, nil
}

func (u userRepo) CheckToken(Id string, Token string) error {
	return nil
}
