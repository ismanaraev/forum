package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"forumv2/internal/models"
	"log"
	"time"
)

type userStorage struct {
	db *sql.DB
}

func newUserSQLite(db *sql.DB) *userStorage {
	return &userStorage{
		db: db,
	}
}

func (u *userStorage) CreateUser(user models.User) error {
	records := `INSERT INTO users(ID,name,username,email,password) VALUES ($1,$2,$3,$4,$5)`

	query, err := u.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("Error in CreateUser method in repository: %w", err)
	}

	_, err = query.Exec(user.ID.String(), user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("Error in CreateUser method in repository: %w", err)
	}

	return nil
}

func (u *userStorage) SetSession(user models.User, token string, time time.Time) error {
	records := `UPDATE users SET token=$1,expiretime=$2 WHERE ID=$3`

	query, err := u.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("Error in SetSession method in repository: %w", err)
	}

	_, err = query.Exec(token, time, user.ID.String())
	if err != nil {
		return fmt.Errorf("error in SetSession method in repository: %w", err)
	}

	fmt.Println("Session created successfully!")
	return nil
}

func (u *userStorage) GetUserInfoByEmail(email string) (models.User, error) {
	row := u.db.QueryRow("SELECT ID,name,username,email,password FROM users WHERE email=$1", email)

	temp := models.User{}
	var userIdStr string
	err := row.Scan(&userIdStr, &temp.Name, &temp.Username, &temp.Email, &temp.Password)
	if err != nil {
		log.Printf("Error with GetUserInfo in repository: %v\n", err)
		return models.User{}, err
	}
	temp.ID, err = models.UserIDFromString(userIdStr)
	if err != nil {
		return models.User{}, err
	}
	return temp, nil
}

func (u *userStorage) GetUsersEmail(user models.User) (models.User, error) {
	row := u.db.QueryRow("SELECT email FROM users WHERE username=$1", user.Username)

	temp := models.User{}
	err := row.Scan(&temp.Email)
	if err != nil {
		log.Printf("Error with GetUsersEmail method in repository: %v\n", err)
		return models.User{}, err
	}
	return temp, nil
}

func (u *userStorage) GetUsersInfoByUUID(id models.UserID) (models.User, error) {
	row := u.db.QueryRow("SELECT name,username,email,password FROM users WHERE ID=$1", id.String())

	temp := models.User{}
	temp.ID = id
	err := row.Scan(&temp.Name, &temp.Username, &temp.Email, &temp.Password)
	if err != nil {
		log.Printf("GetUsersInfoByUUID error: %v\n", err)
		return models.User{}, err
	}
	return temp, nil
}

// CheckUserEmail - returns true if user by this email exists
func (u *userStorage) CheckUserEmail(email string) (UserExist bool, err error) {
	stmt := `SELECT email FROM users WHERE email == $1`
	query, err := u.db.Prepare(stmt)
	if err != nil {
		return false, err
	}
	row := query.QueryRow(email)
	var mail string
	err = row.Scan(&mail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (u *userStorage) CheckUserUsername(username string) (UserExist bool, err error) {
	stmt := `SELECT username FROM users WHERE username == $1`
	query, err := u.db.Prepare(stmt)
	if err != nil {
		return false, err
	}
	row := query.QueryRow(username)
	var name string
	err = row.Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
