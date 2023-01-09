package repository

import (
	"database/sql"
	"fmt"
)

// Создание таблицы пользователя
func CreatUsersTable(db *sql.DB) error {
	users_table := `CREATE TABLE IF NOT EXISTS users (
		uuid TEXT PRIMARY KEY NOT NULL,
		name CHAR(50) NOT NULL,
		username CHAR(50) NOT NULL UNIQUE,
		email CHAR(50) NOT NULL UNIQUE, 
		password CHAR(50) NOT NULL,
		token TEXT,
		expiretime 
	);`

	query, err := db.Prepare(users_table)
	if err != nil {
		return fmt.Errorf("create table in repository: %v", err)
	}

	_, err = query.Exec()
	if err != nil {
		return fmt.Errorf("create table in repository: %v", err)
	}

	fmt.Println("Users table created successfully!")
	return nil
}

// Создание таблицы для поста
func CreatePostTable(db *sql.DB) error {
	post_table := `CREATE TABLE IF NOT EXISTS post (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		uuid TEXT NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author CHAR(50) NOT NULL, 
		createdat CHAR(50) NOT NULL,
		categories,
		like INTEGER DEFAULT 0, 
		dislike INTEGER DEFAULT 0
	);`

	query, err := db.Prepare(post_table)
	if err != nil {
		return fmt.Errorf("create post table in repository: %v", err)
	}

	_, err = query.Exec()
	if err != nil {
		return fmt.Errorf("create post table in repository: %v", err)
	}

	fmt.Println("Post table created successfully!")
	return nil
}

func CreateCommentsTable(db *sql.DB) error {
	comments_table := `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		postID INTEGER,
		content TEXT NOT NULL,
		author CHAR(50) NOT NULL, 
		like INTEGER DEFAULT 0,
		dislike INTEGER DEFAULT 0,
		createdat CHAR(50) NOT NULL
	);`

	query, err := db.Prepare(comments_table)
	if err != nil {
		return fmt.Errorf("create comments table in repository: %v", err)
	}

	_, err = query.Exec()
	if err != nil {
		return fmt.Errorf("create comments table in repository: %v", err)
	}

	fmt.Println("Comments table created successfully!")
	return nil
}

func CreateTableForLikePost(db *sql.DB) error {
	like_table := `CREATE TABLE IF NOT EXISTS likePost(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userID TEXT NOT NULL,
		postID INTEGER NOT NULL,
		status INTEGER DEFAULT 0
		);`

	query, err := db.Prepare(like_table)
	if err != nil {
		return fmt.Errorf("create like table in repository: %v", err)
	}

	_, err = query.Exec()
	if err != nil {
		return fmt.Errorf("create likePost table in repository: %v", err)
	}

	fmt.Println("LikePost table created successfully!")
	return nil
}

func CreateTableForLikeComments(db *sql.DB) error {
	like_table := `CREATE TABLE IF NOT EXISTS likeComments(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userID TEXT NOT NULL,
		commentsID INTEGER NOT NULL,
		status INTEGER DEFAULT 0
		);`

	query, err := db.Prepare(like_table)
	if err != nil {
		return fmt.Errorf("create likeComments table in repository: %v", err)
	}

	_, err = query.Exec()
	if err != nil {
		return fmt.Errorf("create like table in repository: %v", err)
	}

	fmt.Println("LikeComments table created successfully!")
	return nil
}

// func CreateTableForDislike(db *sql.DB) error {
// 	dislike_table := `CREATE TABLE IF NOT EXISTS like(
// 		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
// 		userID TEXT,
// 		postID INTEGER,
// 		status INTEGER
// 		);`

// 	query, err := db.Prepare(dislike_table)
// 	if err != nil {
// 		return fmt.Errorf("create dislike table in repository: %v", err)
// 	}

// 	_, err = query.Exec()
// 	if err != nil {
// 		return fmt.Errorf("create dislike table in repository: %v", err)
// 	}

// 	fmt.Println("Dislike table created successfully!")
// 	return nil
// }
