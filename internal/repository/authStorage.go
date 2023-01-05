package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"forum3/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

type AuthStorage struct {
	db *sql.DB
}

func NewAuthSQLite(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

var (
	UniqueConstraintFailed = errors.New("Such data already exists")
	QueryExecFailed        = errors.New("Query exec failed")
	PrepareNotCorrect      = errors.New("Prepare not correct")
)

// Добавление нового пользователя в базу
func (a *AuthStorage) CreateUser(user models.Auth) (int, error) {
	records := `INSERT INTO users(uuid,name,username,email,password) VALUES ($1,$2,$3,$4,$5)`

	query, err := a.db.Prepare(records)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Create user in repository: %w", PrepareNotCorrect)
	}

	_, err = query.Exec(user.Uuid, user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("Create user in repository: %w", UniqueConstraintFailed)
	}
	fmt.Println("User created successfully!")
	return http.StatusOK, err
}

// Создает токен и время для токена по uuid
func (a *AuthStorage) SetSession(user models.Auth, token string, time time.Time) error {
	records := `UPDATE users SET token=$1,expiretime=$2 WHERE uuid=$3`

	query, err := a.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("Set session in repository: %w", PrepareNotCorrect)
	}

	_, err = query.Exec(token, time, user.Uuid)
	if err != nil {
		return fmt.Errorf("Set session in repository: %w", UniqueConstraintFailed)
	}

	fmt.Println("Session created successfully!")
	return nil
}

// Получить полную информация о юзере с помощью почты
func (a *AuthStorage) GetUserInfo(user models.Auth) (models.Auth, error) {
	row := a.db.QueryRow("SELECT uuid,name,username,email,password FROM users WHERE email=$1", user.Email)

	temp := models.Auth{}
	err := row.Scan(&temp.Uuid, &temp.Name, &temp.Username, &temp.Email, &temp.Password)
	if err != nil {
		log.Printf("GetUserInfo error: %v\n", err)
		return models.Auth{}, err
	}
	return temp, nil
}

// Получить почту юзера по username
func (a *AuthStorage) GetUsersEmail(user models.Auth) (models.Auth, error) {
	row := a.db.QueryRow("SELECT email FROM users WHERE username=$1", user.Username)

	temp := models.Auth{}
	err := row.Scan(&temp.Email)
	if err != nil {
		log.Printf("GetUsersEmail error: %v\n", err)
		return models.Auth{}, err
	}
	return temp, nil
}

func (a *AuthStorage) GetUsersInfoByUUID(id uuid.UUID) (models.Auth, error) {
	row := a.db.QueryRow("SELECT name,username,email,password FROM users WHERE uuid=$1", id)

	temp := models.Auth{}
	err := row.Scan(&temp.Name, &temp.Username, &temp.Email, &temp.Password)
	if err != nil {
		log.Printf("GetUsersInfoByUUID error: %v\n", err)
		return models.Auth{}, err
	}
	return temp, nil
}

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
		fmt.Println(err)
		return fmt.Errorf("Create table in repository: %w", PrepareNotCorrect)
	}

	_, err = query.Exec()
	if err != nil {
		return fmt.Errorf("Create table in repository: %w", QueryExecFailed)
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
		categories
	);`

	query, err := db.Prepare(post_table)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Create post in repository: %w", PrepareNotCorrect)
	}

	_, err = query.Exec()
	if err != nil {
		return fmt.Errorf("Create post in repository: %w", QueryExecFailed)
	}

	fmt.Println("Post table created successfully!")
	return nil
}
