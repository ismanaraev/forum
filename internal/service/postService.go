package service

import (
	"forumv2/internal/models"
	"forumv2/internal/repository"

	"github.com/gofrs/uuid"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (p *PostService) CreatePostService(post models.Post) (int, error) {
	return p.repo.CreatePost(post)
}

func (p *PostService) GetAllPostService(category string) ([]models.Post, error) {
	if category != "all" {
		return p.repo.GetPostWithCategory(category)
	} else {
		return p.repo.GetAllPost()
	}
}

func (p *PostService) GetPostByIDinService(id int) (models.Post, error) {
	return p.repo.GetPostByID(id)
}

func (p *PostService) GetUsersPostInService(uuid uuid.UUID) ([]models.Post, error) {
	return p.repo.GetUsersPost(uuid)
}

func (p *PostService) GetUserLikePostsInService(uuid uuid.UUID) ([]models.Post, error) {
	temp, err := p.repo.GetPostIdWithUUID(uuid)
	if err != nil {
		return nil, err
	}
	return p.repo.GetUsersLikePosts(temp)
}
