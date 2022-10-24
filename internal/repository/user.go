package repository

import (
	"database/sql"
	"forum/internal/models"
)

func newUserRepo(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

type userRepo struct {
	session map[string]string //token to uuid map
	db      *sql.DB
}

func (u userRepo) AddUser(Username, Email, Password string) error {
	return nil
}

func (u userRepo) GetUserById(Id string) (models.User, error) {
	return models.User{}, nil
}

func (u userRepo) CheckToken(Id string, Token string) error {
	return nil
}
