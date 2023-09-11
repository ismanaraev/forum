package db

import (
	"database/sql"
	"fmt"
	"log"
)

var tables []string = []string{
	`CREATE TABLE IF NOT EXISTS users (
		ID TEXT UNIQUE NOT NULL,
		name CHAR(50) NOT NULL,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(50) NOT NULL UNIQUE, 
		password VARCHAR(50) NOT NULL,
		token TEXT,
		expiretime
	);`,
	`CREATE TABLE IF NOT EXISTS post (
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL, 
		createdat INTEGER NOT NULL,
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		FOREIGN KEY (author) REFERENCES users(ID) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS comments (
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		postID INTEGER,
		content TEXT NOT NULL,
		author VARCHAR(50) NOT NULL, 
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		createdat VARCHAR(50) NOT NULL,
		FOREIGN KEY (postID) REFERENCES post(ID) ON DELETE CASCADE,
		FOREIGN KEY (author) REFERENCES users(username) ON DELETE CASCADE
	);`,
	`CREATE TABLE IF NOT EXISTS likePost (
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userID TEXT,
		postID INTEGER DEFAULT 0,
		status INTEGER DEFAULT 0,
		FOREIGN KEY (userID) REFERENCES users(ID) ON DELETE CASCADE,
		FOREIGN KEY (postID) REFERENCES post(ID) ON DELETE CASCADE
		);`,
	`CREATE TABLE IF NOT EXISTS likeComments(
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userID TEXT,
		commentsID INTEGER DEFAULT 0,
		status INTEGER DEFAULT 0,
		FOREIGN KEY (userID) REFERENCES users(ID) ON DELETE CASCADE,
		FOREIGN KEY (commentsID) REFERENCES comments(ID) ON DELETE CASCADE
		);`,
	`CREATE TABLE IF NOT EXISTS categories(
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(50) NOT NULL
		);`,
	`CREATE TABLE IF NOT EXISTS categoriesPost(
		categoryID INTEGER NOT NULL,
		postID INTEGER NOT NULL,
		FOREIGN KEY (categoryID) REFERENCES categories(ID) ON DELETE CASCADE,
		FOREIGN KEY (postID) REFERENCES post(ID) ON DELETE CASCADE
		);`,
}

// Создание таблицы пользователя
func CreatTables(db *sql.DB) error {
	for _, v := range tables {
		stmt, err := db.Prepare(v)
		if err != nil {
			return fmt.Errorf("Create table: %w", err)
		}
		_, err = stmt.Exec()
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("Create table: %w", err)
		}
	}
	log.Println("All table created successfully!")
	return nil
}
