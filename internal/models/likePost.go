package models

import "github.com/gofrs/uuid"

type LikePost struct {
	UserID uuid.UUID  `json:"userid"`
	PostID int        `json:"postid"`
	Status LikeStatus `json:"status"`
}

type LikeStatus int

const Like LikeStatus = 1
const DisLike LikeStatus = -1
