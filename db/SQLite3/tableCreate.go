package db

import (
	"database/sql"
	"log"
	"os"
)

const sQLInitFile = "db/SQLite3/init.sql"

// Создание таблицы пользователя
func CreatTables(db *sql.DB) error {
	stmt, err := os.ReadFile(sQLInitFile)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(stmt))
	if err != nil {
		return err
	}
	log.Println("All table created successfully!")
	return nil
}
