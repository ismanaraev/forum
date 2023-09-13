package repository

import (
	"database/sql"
	"fmt"
	"forumv2/internal/models"
	"time"
)

type commentStorage struct {
	db *sql.DB
}

func newCommentsSQLite(db *sql.DB) *commentStorage {
	return &commentStorage{
		db: db,
	}
}

func (c *commentStorage) CreateComment(com models.Comment) error {
	query, err := c.db.Prepare(`INSERT INTO comments(postID,content,author,createdat) VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return fmt.Errorf("[CommentStorage]:Error with CreateComments method in repository: %w", err)
	}
	_, err = query.Exec(com.PostID, com.Content, com.Author.ID.String(), com.CreatedAt.Truncate(time.Second).Unix())
	if err != nil {
		return fmt.Errorf("create comment in repository: %w", err)
	}

	return nil
}

func (c *commentStorage) GetCommentsByPostID(postID models.PostID) ([]models.Comment, error) {
	row, err := c.db.Query("SELECT ID,postID,content,author,like,dislike,createdat FROM comments WHERE postID=$1", postID)
	if err != nil {
		return nil, fmt.Errorf("[CommentStorage]:Error with GetCommentsByID method in repository: %w", err)
	}

	temp := models.Comment{}
	allComments := []models.Comment{}

	for row.Next() {
		var tempPostID int64
		var tempUuidStr string
		var timestamp int64
		err := row.Scan(&temp.ID, &tempPostID, &temp.Content, &tempUuidStr, &temp.Like, &temp.Dislike, &timestamp)
		if err != nil {
			return nil, fmt.Errorf("[CommentStorage]:Error with GetCommentsByID method in repository: %w", err)
		}
		temp.PostID = models.PostID(tempPostID)
		temp.Author.ID, err = models.UserIDFromString(tempUuidStr)
		if err != nil {
			return nil, err
		}
		temp.CreatedAt = time.Unix(timestamp, 0)
		allComments = append(allComments, temp)
	}
	return allComments, nil
}

func (c *commentStorage) UpdateComment(comment models.Comment) error {
	stmt := `UPDATE comments SET ID = $1, postID = $2, content = $3, author = $4, like = $5, dislike = $6, createdat = $7 WHERE ID == $1`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return fmt.Errorf("error executing statement %v:\n%v", stmt, err)
	}
	_, err = query.Exec(comment.ID, comment.PostID, comment.Content, comment.Author.ID.String(), comment.Like, comment.Dislike, comment.CreatedAt.Truncate(time.Second).Unix())
	if err != nil {
		return fmt.Errorf("error executing statement %v: %v", stmt, err)
	}
	return nil
}

func (c *commentStorage) GetCommentByCommentID(commentID int) (models.Comment, error) {
	stmt := `SELECT ID, postID, content, author, like, dislike, createdat FROM comments WHERE ID == $1`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return models.Comment{}, fmt.Errorf("error executing statement %v: %v", stmt, err)
	}
	var res models.Comment
	var tempPostID int64
	var tempUuidStr string
	var timestamp int64
	err = query.QueryRow(commentID).Scan(&res.ID, &tempPostID, &res.Content, &tempUuidStr, &res.Like, &res.Dislike, &timestamp)
	if err != nil {
		return models.Comment{}, fmt.Errorf("error executing statement %v: %v", stmt, err)
	}
	res.PostID = models.PostID(tempPostID)
	res.Author.ID, err = models.UserIDFromString(tempUuidStr)
	if err != nil {
		return models.Comment{}, err
	}
	res.CreatedAt = time.Unix(timestamp, 0)
	return res, nil
}
