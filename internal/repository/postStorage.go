package repository

import (
	"database/sql"
	"forum3/internal/models"
	"net/http"
)

type PostStorage struct {
	db *sql.DB
}

func NewPostSQLite(db *sql.DB) *PostStorage {
	return &PostStorage{
		db: db,
	}
}

func (p *PostStorage) GetPost(post models.Post) (int, error) {
	return http.StatusOK, nil
}

func (p *PostStorage) CreatePost(post models.Post) (int, error) {
	return http.StatusOK, nil
}

func (p *PostStorage) UpdatePost(post models.Post) (int, error) {
	return http.StatusOK, nil
}

func (p *PostStorage) DeletePost(post models.Post) (int, error) {
	return http.StatusOK, nil
}
