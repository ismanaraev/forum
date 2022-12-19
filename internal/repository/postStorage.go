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

func (p *PostStorage) GetAllPost() (models.Post, error) {
	row := p.db.QueryRow("SELECT * FROM post")

	temp := models.Post{}

	err := row.Scan(&temp.Uuid, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories)
	if err != nil {
		fmt.Println(err)
		return temp, fmt.Errorf("Get post in repository: %w", UniqueConstraintFailed)
	}
	// uuid,title,content,author,createdat,categories
	return temp, nil
}

func (p *PostStorage) CreatePost(post models.Post) (int, error) {
	query, err := p.db.Prepare(`INSERT INTO post(uuid,title,content,author,createdat,categories) VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Create post in repository: %w", PrepareNotCorrect)
	}

	_, err = query.Exec(post.Uuid, post.Title, post.Content, post.Author, post.CreatedAt, post.Categories)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("Create post in repository: %w", UniqueConstraintFailed)
	}
	fmt.Println("Post created successfully!")

	return http.StatusOK, nil
}

func (p *PostStorage) UpdatePost(post models.Post) (int, error) {
	return http.StatusOK, nil
}

func (p *PostStorage) DeletePost(post models.Post) (int, error) {
	return http.StatusOK, nil
}
