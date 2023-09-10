package models

type LikeStatus int

const (
	Like    LikeStatus = 1
	DisLike LikeStatus = -1
	NoLike  LikeStatus = 0
)

type LikePost struct {
	UserID UserID     `json:"userid"`
	PostID PostID     `json:"postid"`
	Status LikeStatus `json:"status"`
}
