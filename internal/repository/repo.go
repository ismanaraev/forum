package repository

import (
	"database/sql"
	"forum/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

type UserRepository interface {
	AddUser(Username, Email, Password string) error
	GetUserById(Id string) (models.User, error)
	CheckToken(Id string, Token string) error
}

type Repository struct {
	User UserRepository
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		User: newUserRepo(db),
	}
}
