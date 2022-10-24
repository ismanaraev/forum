package models

// TODO: add Comments field which contains []models.Comment and Likes field, which contains []models.Like
type Post struct {
	Id     string
	Title  string
	Text   string
	Author User
}
