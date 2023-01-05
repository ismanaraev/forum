package models

import (
	"github.com/gofrs/uuid"
)

type Post struct {
	Uuid       uuid.UUID `json:"uuid"`
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	CreatedAt  string    `json:"createdat"`
	Categories string    `json:"categories"`
}
