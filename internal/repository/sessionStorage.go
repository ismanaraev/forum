package repository

import (
	"database/sql"
	"fmt"
	"forumv2/internal/models"
)

type SessionStorage struct {
	db *sql.DB
}

func NewSessionSQLite(db *sql.DB) *SessionStorage {
	return &SessionStorage{
		db: db,
	}
}

func (s *SessionStorage) GetSessionFromDB(token string) (models.UserID, error) {
	row := s.db.QueryRow("SELECT ID FROM users WHERE token=$1", token)
	var temp string
	err := row.Scan(&temp)
	if err != nil {
		return models.UserID{}, err
	}
	res, err := models.UserIDFromString(temp)
	if err != nil {
		return res, fmt.Errorf("[SessionStorage]:Error with GetSessionFromDB method in repository: %w", err)
	}
	return res, nil
}

// Запрос на удаление по uuid токена и время токена
func (s *SessionStorage) DeleteSessionFromDB(uuid models.UserID) error {
	records := ("UPDATE users SET token = NULL, expiretime = NULL WHERE ID = $1")

	query, err := s.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("[SessionStorage]:Error with DeleteSessionFromDB method in repository: %w", err)
	}

	_, err = query.Exec(uuid.String())
	if err != nil {
		return fmt.Errorf("[SessionStorage]:Error with DeleteSessionFromDB method in repository: %w", err)
	}

	fmt.Println("Session deleted successfully!")
	return nil
}
