package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	Uuid       uuid.UUID `json:"uuid"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	CreatedAt  time.Time `json:"createdat"`
	Categories string    `json:"categories"`
}
