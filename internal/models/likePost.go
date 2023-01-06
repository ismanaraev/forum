package models

import "github.com/gofrs/uuid"

type LikePost struct {
	UserID uuid.UUID `json:"userid"`
	PostID int       `json:"postid"`
	Status int       `json:"status"`
}
