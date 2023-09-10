package models

type LikeComment struct {
	UserID     UserID     `json:"userid"`
	CommentsID int        `json:"commentsid"`
	Status     LikeStatus `json:"status"`
}
