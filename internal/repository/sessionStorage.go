package repository

import (
	"database/sql"
	"fmt"
	"forum3/internal/models"

	"github.com/gofrs/uuid"
)

type SessionStorage struct {
	db *sql.DB
}

func NewSessionSQLite(db *sql.DB) *SessionStorage {
	return &SessionStorage{
		db: db,
	}
}

// Запрос на БД что бы получить uuid по токену
func (s *SessionStorage) GetSessionFromDB(token string) (uuid.UUID, error) {
	row := s.db.QueryRow("SELECT uuid FROM users WHERE token=$1", token)
	temp := models.Auth{}
	err := row.Scan(&temp.Uuid)
	if err != nil {
		return temp.Uuid, err
	}
	return temp.Uuid, nil
}

// Запрос на удаление по username токена и время токена
func (s *SessionStorage) DeleteSessionFromDB(r uuid.UUID) error {
	records := ("UPDATE users SET token = NULL, expiretime = NULL WHERE uuid = $1")

	query, err := s.db.Prepare(records)
	if err != nil {
		return fmt.Errorf("Delete session in repository: %w", PrepareNotCorrect)
	}

	_, err = query.Exec(r)
	if err != nil {
		return fmt.Errorf("Delete session in repository: %w", UniqueConstraintFailed)
	}

	fmt.Println("Session deleted successfully!")
	return nil
}
