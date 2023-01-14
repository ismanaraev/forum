package service

import (
	"errors"
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

func (p *PostService) CreatePostService(post models.Post) (int64, error) {
	return p.repo.CreatePost(post)
}

func (p *PostService) GetAllPostService() ([]models.Post, error) {
	return p.repo.GetAllPost()
}

func (p *PostService) GetPostByIDinService(id int64) (models.Post, error) {
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

func (p *PostService) FilterPostsByCategories(categories []string) ([]models.Post, error) {
	category, err := p.CreateCategory(categories)
	if err != nil {
		return nil, err
	}
	return p.repo.GetPostsByCategory(category)
}

func (p *PostService) CreateCategory(categories []string) (models.Category, error) {
	var category models.Category
	for _, val := range categories {
		switch val {
		case "Coding":
			category = category | models.Coding
		case "Music":
			category = category | models.Music
		case "Art":
			category = category | models.Art
		case "Sports":
			category = category | models.Sports
		case "Cooking":
			category = category | models.Cooking
		case "Other":
			category = category | models.Other

		default:
			return models.Other, errors.New("invalid category")
		}
	}
	return category, nil
}
