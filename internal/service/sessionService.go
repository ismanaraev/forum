package service

import (
	"forumv2/internal/models"
)

type sessionService struct {
	repo Session
}

func newSessionService(repo Session) *sessionService {
	return &sessionService{
		repo: repo,
	}
}

func (s *sessionService) DeleteSessionService(uuid models.UserID) error {
	err := s.repo.DeleteSessionFromDB(uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionService) GetSessionService(token string) (models.UserID, error) {
	uuid, err := s.repo.GetSessionFromDB(token)
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}
