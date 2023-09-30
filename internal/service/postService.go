package service

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"forumv2/internal/models"
	"strings"
)

type postService struct {
	repo Repository
}

func newPostService(repo Repository) *postService {
	return &postService{
		repo: repo,
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
	if len(post.Title) > models.PostMaxTitleLen {
		return errors.New("title too long")
	}
	if len(post.Content) == 0 {
		return errors.New("empty content")
	}
	if len(post.Content) > models.PostMaxContentLen {
		return errors.New("content too long")
	}
	return nil
}

func (p *postService) CreatePostService(post models.Post) (models.PostID, error) {
	postId, err := p.repo.CreatePost(post)
	if err != nil {
		return 0, err
	}
	for _, val := range post.Categories {
		err = p.repo.AddCategoryToPost(postId, val.ID)
		if err != nil {
			return 0, err
		}
	}
	if len(post.Pictures) != 0 {

		for _, pic := range post.Pictures {
			temp := base64.StdEncoding.EncodeToString([]byte(pic.Value))
			err = p.repo.AddPictureToPost(postId, models.Picture{Value: temp, Type: pic.Type, Size: pic.Size})
			if err != nil {
				return 0, err
			}
		}
	}
	return postId, nil
}

func (p *postService) GetAllPostService() ([]models.Post, error) {
	posts, err := p.repo.GetAllPost()
	if err != nil {
		return nil, err
	}
	for i := range posts {
		posts[i].Author, err = p.repo.GetUsersInfoByUUID(posts[i].Author.ID)
		if err != nil {
			return nil, err
		}
		categories, err := p.repo.GetCategoriesByPostID(posts[i].ID)
		if err != nil {
			return nil, err
		}
		posts[i].Categories = categories
		posts[i].Pictures, err = p.repo.GetPicturesByPostID(posts[i].ID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}
		}
	}
	return posts, nil
}

func (p *postService) GetPostByIDinService(id models.PostID) (models.Post, error) {
	post, err := p.repo.GetPostByID(id)
	if err != nil {
		return models.Post{}, err
	}
	author, err := p.repo.GetUsersInfoByUUID(post.Author.ID)
	if err != nil {
		return models.Post{}, err
	}
	categories, err := p.repo.GetCategoriesByPostID(post.ID)
	if err != nil {
		return models.Post{}, err
	}
	picture, err := p.repo.GetPicturesByPostID(id)
	if err != nil {
		return models.Post{}, err
	}
	post.Author = author
	post.Categories = categories
	post.Pictures = picture
	return post, nil
}

func (p *postService) GetUsersPostInService(id models.UserID) ([]models.Post, error) {
	posts, err := p.repo.GetPostsByUserID(id)
	if err != nil {
		return nil, err
	}
	user, err := p.repo.GetUsersInfoByUUID(id)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		categories, err := p.repo.GetCategoriesByPostID(posts[i].ID)
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
		temp, err := p.repo.GetCategoryByName(val)
		if err != nil {
			return nil, err
		}
		categories = append(categories, *temp)
	}
	posts, err := p.repo.FilterPostsByMultipleCategories(categories)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		author, err := p.repo.GetUsersInfoByUUID(posts[i].Author.ID)
		if err != nil {
			return nil, err
		}
		postCats, err := p.repo.GetCategoriesByPostID(posts[i].ID)
		if err != nil {
			return nil, err
		}
		postPics, err := p.repo.GetPicturesByPostID(posts[i].ID)
		if err != nil {
			return nil, err
		}
		posts[i].Author = author
		posts[i].Categories = postCats
		posts[i].Pictures = postPics
	}
	return posts, nil
}

func (p *postService) GetCategoryByName(name string) (models.Category, error) {
	res, err := p.repo.GetCategoryByName(name)
	return *res, err
}

func (p *postService) GetUsersLikedPosts(id models.UserID) ([]models.Post, error) {
	posts, err := p.repo.GetUsersLikePosts(id)
	if err != nil {
		return nil, err
	}
	author, err := p.repo.GetUsersInfoByUUID(id)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		cats, err := p.repo.GetCategoriesByPostID(posts[i].ID)
		if err != nil {
			return nil, err
		}
		posts[i].Categories = cats
		posts[i].Author = author
	}
	return posts, nil
}

func (p *postService) UpdatePost(post models.Post) error {
	return p.repo.UpdatePost(post)
}

func (p *postService) DeletePostByID(ID models.PostID) error {
	return p.repo.DeletePostByID(ID)
}
