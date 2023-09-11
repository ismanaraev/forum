package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserID uuid.UUID

type User struct {
	ID         UserID    `json:"uuid"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Token      string    `json:"token"`
	ExpireTime time.Time `json:"expiretime"`
}

func (u UserID) String() string {
	res := uuid.UUID(u)
	return res.String()
}

func UserIDFromString(s string) (UserID, error) {
	res, err := uuid.FromString(s)
	if err != nil {
		return UserID{}, err
	}
	return UserID(res), nil
}
