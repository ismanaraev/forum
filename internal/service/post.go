package service

import "forum/internal/models"

type PostRepository interface {
	AddPost(Title, Text string, Author models.User) (Id string, err error)
	GetPostById(Id string) (models.Post, error)
}

type postService struct {
	repo PostRepository
}

func (p postService) CreatePost(Title, Text string, token string) {

}
