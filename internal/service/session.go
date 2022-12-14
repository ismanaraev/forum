package service

import (
	"forum3/internal/repository"

	"github.com/gofrs/uuid"
)

type SessionService struct {
	repo repository.Session
}

func NewSessionService(repo repository.Session) *SessionService {
	return &SessionService{
		repo: repo,
	}
}

// RQ-request
// Запрос на удаления токена и время токена
func (s *SessionService) DeleteSessionRQtoRepo(tokenString string) {
	// delete(s.session, tokenString)
}

// Получаем по токену username
func (s *SessionService) GetSessionRQtoRepo(token string) (uuid.UUID, error) {
	uuid, err := s.repo.GetSessionFromDB(token)
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}
