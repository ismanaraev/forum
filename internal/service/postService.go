package service

import (
	"forum3/internal/models"
	"forum3/internal/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (p *PostService) GetAllPostService() ([]models.Post, error) {
	return p.repo.GetAllPost()
}

func (p *PostService) CreatePostService(post models.Post) (int, error) {
	return p.repo.CreatePost(post)
}

func (p *PostService) UpdatePostService(post models.Post) (int, error) {
	return p.repo.UpdatePost(post)
}

func (p *PostService) DeletePostService(post models.Post) (int, error) {
	return p.repo.DeletePost(post)
}
