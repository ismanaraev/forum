package service

import (
	"forum3/internal/models"
	"forum3/internal/repository"
	"net/http"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (p *PostService) GetPostService(post models.Post) (int, error) {
	return http.StatusOK, nil
}

func (p *PostService) CreatePostService(post models.Post) (int, error) {
	return http.StatusOK, nil
}

func (p *PostService) UpdatePostService(post models.Post) (int, error) {
	return http.StatusOK, nil
}

func (p *PostService) DeletePostService(post models.Post) (int, error) {
	return http.StatusOK, nil
}
