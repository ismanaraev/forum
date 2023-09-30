package models

import "time"

type PostID int64

const PostMaxTitleLen = 50

const PostMaxContentLen = 1000

type Post struct {
	ID         PostID     `json:"id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	Author     User       `json:"author"`
	CreatedAt  time.Time  `json:"createdat"`
	Categories []Category `json:"categories"`
	Pictures   []Picture  `json:"pictures"`
	Like       int        `json:"like"`
	Dislike    int        `json:"dislike"`
}
