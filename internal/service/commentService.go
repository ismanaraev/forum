package service

import (
	"forumv2/internal/models"
	"forumv2/internal/repository"
)

type CommentService struct {
	repo repository.Comments
}

func NewCommentsService(repo repository.Comments) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (c *CommentService) GetAllCommentsInService() ([]models.Comment, error) {
	return c.repo.GetAllComments()
}

func (c *CommentService) GetCommentsByIDinService(postID int64) ([]models.Comment, error) {
	return c.repo.GetCommentsByID(postID)
}

func (c *CommentService) CreateCommentsInService(com models.Comment) error {
	return c.repo.CreateComments(com)
}
