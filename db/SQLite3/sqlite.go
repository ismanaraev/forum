package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Database(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}
	err = CreatTables(db)
	if err != nil {
		return nil, fmt.Errorf("can't table create %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}
	query := `PRAGMA foreign_keys=1;`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	stmt.Exec()
	return db, nil
}
