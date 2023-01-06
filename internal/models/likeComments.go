package models

import "github.com/gofrs/uuid"

type LikeComments struct {
	UserID     uuid.UUID `json:"userid"`
	CommentsID int       `json:"commentsid"`
	Status     int       `json:"status"`
}
