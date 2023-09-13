package service

import (
	"errors"
	"forumv2/internal/models"
	"strings"
)

type commentService struct {
	repo Repository
}

func newCommentsService(repo Repository) *commentService {
	return &commentService{
		repo: repo,
	}
}

func (c *commentService) CheckCommentInput(comment models.Comment) error {
	if comment := strings.Trim(comment.Content, "\r\n "); len(comment) == 0 {
		return errors.New("empty title")
	}
	if len(comment.Content) == 0 {
		return errors.New("empty comment")
	}
	if len(comment.Content) > 500 {
		return errors.New("comment too long")
	}
	return nil
}

func (c *commentService) GetCommentsByPostID(postID models.PostID) ([]models.Comment, error) {
	comments, err := c.repo.GetCommentsByPostID(postID)
	if err != nil {
		return nil, err
	}
	for i := range comments {
		comments[i].Author, err = c.repo.GetUsersInfoByUUID(comments[i].Author.ID)
		if err != nil {
			return nil, err
		}
	}
	return comments, nil
}

func (c *commentService) CreateComment(com models.Comment) error {
	return c.repo.CreateComment(com)
}
