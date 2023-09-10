package service

import (
	"forumv2/internal/models"
	"forumv2/internal/repository"
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
func (s *SessionService) DeleteSessionService(uuid models.UserID) error {
	err := s.repo.DeleteSessionFromDB(uuid)
	if err != nil {
		return err
	}
	return nil
}

// Получаем по токену username
func (s *SessionService) GetSessionService(token string) (models.UserID, error) {
	uuid, err := s.repo.GetSessionFromDB(token)
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}
