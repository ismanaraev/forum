package service

import (
	"forum3/internal/models"
	"forum3/internal/repository"
	"log"
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

func (p *PostService) GetPostByIDinService(id int) (models.Post, error) {
	return p.repo.GetPostByID(id)
}

func (p *PostService) GetAllCommentsInService() ([]models.Comments, error) {
	return p.repo.GetAllComments()
}

func (p *PostService) GetCommentsByIDinService(postID int) ([]models.Comments, error) {
	return p.repo.GetCommentsByID(postID)
}

func (p *PostService) CreateCommentsInService(com models.Comments) (int, error) {
	return p.repo.CreateComments(com)
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

func (p *PostService) LikeInService(like models.LikePost) (models.LikePost, error) {
	log.Println("ABOBA")
	exist, err := p.repo.CheckLikeByPostAndUserID(like)
	if err != nil {
		return models.LikePost{}, err
	}
	log.Println("ABOBUS")
	if exist {
		return p.repo.UpdateLikeStatus(like)
	} else {
		return p.repo.CreateLikeForPost(like)
	}
}

// func (p *PostService) CreateLikeTable(like models.LikePost) (models.LikePost, error) {
// 	return p.repo.CreateLikeForPost(like)
// }

// func (p *PostService) CounterLikeInService() int {
// 	return p.repo.CounterLike()
// }
