package models

import "time"

type Comment struct {
	ID        int
	PostID    PostID    `json:"postid"`
	Author    User      `json:"author"`
	Content   string    `json:"content"`
	Like      int       `json:"like"`
	Dislike   int       `json:"dislike"`
	CreatedAt time.Time `json:"createdat"`
}
