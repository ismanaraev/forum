package repository

import (
	"database/sql"
	"fmt"
	"forum3/internal/models"
	"net/http"
)

type PostStorage struct {
	db *sql.DB
}

func NewPostSQLite(db *sql.DB) *PostStorage {
	return &PostStorage{
		db: db,
	}
}

func (p *PostStorage) GetAllPost() ([]models.Post, error) {
	temp := models.Post{}
	allPost := []models.Post{}
	n := 1

	for i := 1; i <= n; i++ {
		row := p.db.QueryRow("SELECT id,uuid,title,content,author,createdAt,categories FROM post WHERE id=$1", n)
		err := row.Scan(&temp.ID, &temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories)
		if err != nil {
			//	fmt.Println(err)
			return allPost, nil
		}
		allPost = append(allPost, temp)
		n++
	}

	// uuid,title,content,author,createdat,categories
	return allPost, nil
}

func (p *PostStorage) GetAllComments() ([]models.Comments, error) {
	row, err := p.db.Query("SELECT postID,content,author,like,dislike,createdat FROM comments")
	if err != nil {
		return nil, fmt.Errorf("SELECT allcomments in repository: %w", err)
	}

	temp := models.Comments{}
	allComments := []models.Comments{}

	for row.Next() {
		err := row.Scan(&temp.PostID, &temp.Content, &temp.Author, &temp.Like, &temp.Dislike, &temp.CreatedAt)
		if err != nil {
			return nil, err
		}
		allComments = append(allComments, temp)
	}
	return allComments, nil
}

func (p *PostStorage) CreatePost(post models.Post) (int, error) {
	query, err := p.db.Prepare(`INSERT INTO post(uuid,title,content,author,createdat,categories) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Create post in repository: %w", PrepareNotCorrect)
	}

	_, err = query.Exec(post.Uuid, post.Title, post.Content, post.Author, post.CreatedAt, post.Categories)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("Create post in repository: %w", err)
	}
	fmt.Println("Post created successfully!")

	return http.StatusOK, nil
}

func (p *PostStorage) GetPostByID(id int) (models.Post, error) {
	row := p.db.QueryRow("SELECT id,uuid,title,content,author,createdAt,categories FROM post WHERE id=$1", id)

	temp := models.Post{}
	err := row.Scan(&temp.ID, &temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories)
	if err != nil {
		return temp, fmt.Errorf("Create post in repository: %w", err)
	}
	return temp, nil
}

func (p *PostStorage) GetCommentsByID(postID int) ([]models.Comments, error) {
	row, err := p.db.Query("SELECT postID,content,author,like,dislike,createdat FROM comments WHERE postID=$1", postID)
	if err != nil {
		return nil, fmt.Errorf("SELECT comments by id in repository: %w", err)
	}

	temp := models.Comments{}
	allComments := []models.Comments{}

	for row.Next() {
		err := row.Scan(&temp.PostID, &temp.Content, &temp.Author, &temp.Like, &temp.Dislike, &temp.CreatedAt)
		if err != nil {
			return nil, err
		}
		allComments = append(allComments, temp)
	}
	return allComments, nil
}

func (p *PostStorage) CreateComments(com models.Comments) (int, error) {
	query, err := p.db.Prepare(`INSERT INTO comments(postID,content,author,like,dislike,createdat) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Create comment in repository: %w", err)
	}

	_, err = query.Exec(com.PostID, com.Content, com.Author, com.Like, com.Dislike, com.CreatedAt)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("Create comment in repository: %w", err)
	}
	fmt.Println("Comment created successfully!")

	return http.StatusOK, nil
}

func (p *PostStorage) UpdatePost(post models.Post) (int, error) {
	return http.StatusOK, nil
}

func (p *PostStorage) DeletePost(post models.Post) (int, error) {
	return http.StatusOK, nil
}
