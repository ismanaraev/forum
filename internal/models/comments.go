package models

type Comments struct {
	PostID    int    `json:"postid"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	Like      int    `json:"like"`
	Dislike   int    `json:"dislike"`
	CreatedAt string `json:"createdat"`
}
