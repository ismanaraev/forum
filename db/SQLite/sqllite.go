package db

import (
	"database/sql"
	"fmt"
	"forum3/internal/repository"

	_ "github.com/mattn/go-sqlite3"
)

func Database() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../db/SQLite/store.db")
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	err = repository.CreatUsersTable(db)
	if err != nil {
		return nil, fmt.Errorf("can't create user table %v", err)
	}

	err = repository.CreatePostTable(db)
	if err != nil {
		return nil, fmt.Errorf("can't create post table %v", err)
	}

	err = repository.CreateCommentsTable(db)
	if err != nil {
		return nil, fmt.Errorf("can't create comments table %v", err)
	}

	err = repository.CreateTableForLikePost(db)
	if err != nil {
		return nil, fmt.Errorf("can't create like table %v", err)
	}

	err = repository.CreateTableForLikeComments(db)
	if err != nil {
		return nil, fmt.Errorf("can't create dislike table %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	return db, nil
}
