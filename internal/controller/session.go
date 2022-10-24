package controller

import (
	"forum/internal/models"
	"forum/internal/service"
)

type UserRepository interface {
	GetUserByUUID(UUID string) (models.User, error)
}

type SessionManager struct {
	Service service.Service
	Session map[string]string
}

func (s SessionManager) CreateSession(Username, Password string) (Token string, err error) {
	err = s.Service.User.Authorize(Username, Password)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (s SessionManager) ValidateSession(Token string) error {
	return nil
}

func (s SessionManager) DeleteSession(Token string) error {
	return nil
}
