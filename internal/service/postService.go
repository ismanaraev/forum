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

func (p *PostService) LikePost(like models.LikePost) (models.LikePost, error) {
	status, err := p.repo.GetLikeStatusByPostAndUserID(like)
	if err != nil {
		return models.LikePost{}, err
	}
	if status == models.NoLike {
		return p.repo.CreateLikeForPost(like)
	} else {
		if status != like.Status {
			return p.repo.UpdatePostLikeStatus(like)
		}
		return like, nil
	}
}

func (p *PostService) LikeComment(like models.LikeComments) (models.LikeComments, error) {
	status, err := p.repo.GetLikeStatusByCommentAndUserID(like)
	if err != nil {
		return models.LikeComments{}, err
	}
	if status == models.NoLike {
		return p.repo.CreateLikeForComment(like)
	} else {
		if status != like.Status {
			return p.repo.UpdateCommentLikeStatus(like)
		}
		return like, nil
	}
}

// func (p *PostService) CreateLikeTable(like models.LikePost) (models.LikePost, error) {
// 	return p.repo.CreateLikeForPost(like)
// }

// func (p *PostService) CounterLikeInService() int {
// 	return p.repo.CounterLike()
// }
