package service

import (
	"errors"
	"forumv2/internal/models"
	"forumv2/internal/repository"
	"strings"
)

type PostService struct {
	repo     repository.Post
	userRepo repository.User
}

func NewPostService(repo repository.Post) Post {
	return &PostService{
		repo: repo,
	}
}

func (p *PostService) CheckPostInput(post models.Post) error {
	if len(post.Title) == 0 {
		return errors.New("empty title")
	}
	if title := strings.Trim(post.Title, "\r\n "); len(title) == 0 {
		return errors.New("empty title")
	}
	if content := strings.Trim(post.Content, "\r\n "); len(content) == 0 {
		return errors.New("empty title")
	}
	if len(post.Title) > 50 {
		return errors.New("title too long")
	}
	if len(post.Content) == 0 {
		return errors.New("empty content")
	}
	if len(post.Content) > 1000 {
		return errors.New("content too long")
	}
	return nil
}

func (p *PostService) CreatePostService(post models.Post) (models.PostID, error) {
	return p.repo.CreatePost(post)
}

func (p *PostService) GetAllPostService() ([]models.Post, error) {
	posts, err := p.repo.GetAllPost()
	if err != nil {
		return nil, err
	}
	for i := range posts {
		posts[i].Author, err = p.userRepo.GetUsersInfoByUUID(posts[i].Author.ID)
		if err != nil {
			return nil, err
		}
	}
	return posts, nil
}

func (p *PostService) GetPostByIDinService(id models.PostID) (models.Post, error) {
	post, err := p.repo.GetPostByID(id)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (p *PostService) GetUsersPostInService(id models.UserID) ([]models.Post, error) {
	posts, err := p.repo.GetPostsByUserID(id)
	if err != nil {
		return nil, err
	}
	user, err := p.userRepo.GetUsersInfoByUUID(id)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		posts[i].Author = user
	}
	return posts, nil
}

func (p *PostService) FilterPostsByCategories(categoriesString []string) ([]models.Post, error) {
	var categories []models.Category
	for _, val := range categoriesString {
		temp, err := p.repo.GetCategoryByName(val)
		if err != nil {
			return nil, err
		}
		categories = append(categories, temp)
	}
	posts, err := p.repo.FilterPostsByMultipleCategories(categories)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostService) CreateCategory(name string) error {
	return p.repo.CreateCategory(name)
}

func (p *PostService) GetCategoryByName(name string) (models.Category, error) {
	return p.repo.GetCategoryByName(name)
}

func (p *PostService) GetUsersLikedPosts(id models.UserID) ([]models.Post, error) {
	return p.repo.GetUsersLikePosts(id)
}
