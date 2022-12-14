package db

import (
	"database/sql"
	"fmt"
	"forum3/internal/repository"
)

func Database() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../db/SQLite/store.db")
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	err = repository.CreatTable(db)
	if err != nil {
		return nil, fmt.Errorf("can't create table %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	return db, nil
}
