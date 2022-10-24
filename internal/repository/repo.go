package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Repository struct {
	User UserRepository
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		User: newUserRepo(db),
	}
}
