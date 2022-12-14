package repository

import (
	"database/sql"
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

// Запрос на БД что бы получить username по токену
func (s *SessionStorage) GetSessionFromDB(token string) (uuid.UUID, error) {
	row := s.db.QueryRow("SELECT email FROM users WHERE email=$1", token)
	temp := models.Auth{}
	err := row.Scan(&temp.Uuid)
	if err != nil {
		return temp.Uuid, err
	}
	return temp.Uuid, nil
}

// Запрос на удаление по username токена и время токена
func (s *SessionStorage) DeleteSessionFromDB(user models.Auth) (int, error) {
	return 0, nil
}
