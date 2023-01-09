package repository

import (
	"database/sql"
	"fmt"
	"forumv2/internal/models"
	"net/http"
)

type CommentStorage struct {
	db *sql.DB
}

func NewCommentsSQLite(db *sql.DB) *CommentStorage {
	return &CommentStorage{
		db: db,
	}
}

func (c *CommentStorage) CreateComments(com models.Comment) (int, error) {
	query, err := c.db.Prepare(`INSERT INTO comments(postID,content,author,like,dislike,createdat) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("[CommentStorage]:Error with CreateComments method in repository: %w", err)
	}

	_, err = query.Exec(com.PostID, com.Content, com.Author, com.Like, com.Dislike, com.CreatedAt)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Create comment in repository: %w", err)
	}
	fmt.Println("Comment created successfully!")

	return http.StatusOK, nil
}

func (c *CommentStorage) GetAllComments() ([]models.Comment, error) {
	row, err := c.db.Query("SELECT id, postID,content,author,like,dislike,createdat FROM comments")
	if err != nil {
		return nil, fmt.Errorf("[CommentStorage]:Error with GetAllComments method in repository: %w", err)
	}

	temp := models.Comment{}
	allComments := []models.Comment{}

	for row.Next() {
		err := row.Scan(&temp.PostID, &temp.Content, &temp.Author, &temp.Like, &temp.Dislike, &temp.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("[CommentStorage]:Error with GetAllComments method in repository: %w", err)
		}
		allComments = append(allComments, temp)
	}
	return allComments, nil
}

func (c *CommentStorage) GetCommentsByID(postID int) ([]models.Comment, error) {
	row, err := c.db.Query("SELECT id,postID,content,author,like,dislike,createdat FROM comments WHERE postID=$1", postID)
	if err != nil {
		return nil, fmt.Errorf("[CommentStorage]:Error with GetCommentsByID method in repository: %w", err)
	}

	temp := models.Comment{}
	allComments := []models.Comment{}

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.PostID, &temp.Content, &temp.Author, &temp.Like, &temp.Dislike, &temp.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("[CommentStorage]:Error with GetCommentsByID method in repository: %w", err)
		}
		allComments = append(allComments, temp)
	}
	return allComments, nil
}
