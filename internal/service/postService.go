package service

import (
	"errors"
	"forumv2/internal/models"
	"strings"
)

type postService struct {
	repo           Post
	userRepo       User
	categoriesRepo Categories
}

func newPostService(repo Post, userRepo User, catRepo Categories) *postService {
	return &postService{
		repo:           repo,
		userRepo:       userRepo,
		categoriesRepo: catRepo,
	}
}

func (p *postService) CheckPostInput(post models.Post) error {
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

func (p *postService) CreatePostService(post models.Post) (models.PostID, error) {
	return p.repo.CreatePost(post)
}

func (p *postService) GetAllPostService() ([]models.Post, error) {
	posts, err := p.repo.GetAllPost()
	if err != nil {
		return nil, err
	}
	for i := range posts {
		posts[i].Author, err = p.userRepo.GetUsersInfoByUUID(posts[i].Author.ID)
		if err != nil {
			return nil, err
		}
		categories, err := p.categoriesRepo.GetCategoriesByPostID(posts[i].ID)
		if err != nil {
			return nil, err
		}
		posts[i].Categories = categories
	}
	return posts, nil
}

func (p *postService) GetPostByIDinService(id models.PostID) (models.Post, error) {
	post, err := p.repo.GetPostByID(id)
	if err != nil {
		return models.Post{}, err
	}
	author, err := p.userRepo.GetUsersInfoByUUID(post.Author.ID)
	if err != nil {
		return models.Post{}, err
	}
	categories, err := p.categoriesRepo.GetCategoriesByPostID(post.ID)
	if err != nil {
		return models.Post{}, err
	}
	post.Author = author
	post.Categories = categories
	return post, nil
}

func (p *postService) GetUsersPostInService(id models.UserID) ([]models.Post, error) {
	posts, err := p.repo.GetPostsByUserID(id)
	if err != nil {
		return nil, err
	}
	user, err := p.userRepo.GetUsersInfoByUUID(id)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		categories, err := p.categoriesRepo.GetCategoriesByPostID(posts[i].ID)
		if err != nil {
			return nil, err
		}
		posts[i].Categories = categories
		posts[i].Author = user
	}
	return posts, nil
}

func (p *postService) FilterPostsByCategories(categoriesString []string) ([]models.Post, error) {
	var categories []models.Category
	for _, val := range categoriesString {
		temp, err := p.categoriesRepo.GetCategoryByName(val)
		if err != nil {
			return nil, err
		}
		categories = append(categories, temp)
	}
	posts, err := p.repo.FilterPostsByMultipleCategories(categories)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		author, err := p.userRepo.GetUsersInfoByUUID(posts[i].Author.ID)
		if err != nil {
			return nil, err
		}
		postCats, err := p.categoriesRepo.GetCategoriesByPostID(posts[i].ID)
		if err != nil {
			return nil, err
		}
		posts[i].Author = author
		posts[i].Categories = postCats
	}
	return posts, nil
}

func (p *postService) CreateCategory(name string) error {
	return p.categoriesRepo.CreateCategory(name)
}

func (p *postService) GetCategoryByName(name string) (models.Category, error) {
	return p.categoriesRepo.GetCategoryByName(name)
}

func (p *postService) GetUsersLikedPosts(id models.UserID) ([]models.Post, error) {
	posts, err := p.repo.GetUsersLikePosts(id)
	if err != nil {
		return nil, err
	}
	author, err := p.userRepo.GetUsersInfoByUUID(id)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		cats, err := p.categoriesRepo.GetCategoriesByPostID(posts[i].ID)
		if err != nil {
			return nil, err
		}
		posts[i].Categories = cats
		posts[i].Author = author
	}
	return posts, nil
}
