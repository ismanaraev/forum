package repository

import (
	"database/sql"
	"fmt"
	"forumv2/internal/models"
	"log"
	"time"
)

type PostStorage struct {
	db *sql.DB
}

func NewPostSQLite(db *sql.DB) *PostStorage {
	return &PostStorage{
		db: db,
	}
}

func (p *PostStorage) CreatePost(post models.Post) (models.PostID, error) {
	query, err := p.db.Prepare(`INSERT INTO post(title,content,author,createdat) VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return 0, fmt.Errorf("[PostStorage]:Error with CreatePost method in repository: %w", err)
	}

	res, err := query.Exec(post.Title, post.Content, post.Author.ID.String(), post.CreatedAt.Truncate(time.Second).Unix())
	if err != nil {
		return 0, fmt.Errorf("[PostStorage]:Error with CreatePost method in repository: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	log.Println("Post created successfully!")

	return models.PostID(id), nil
}

// Запрос на все посты
func (p *PostStorage) GetAllPost() ([]models.Post, error) {
	row, err := p.db.Query("SELECT ID,title,content,author,createdAt FROM post")
	if err != nil {
		return nil, fmt.Errorf("[PostStorage]:Error with GetAllPost method in repository: %w", err)
	}

	allPost := []models.Post{}

	for row.Next() {
		var temp models.Post
		var id string
		var timestamp int64
		err := row.Scan(&temp.ID, &temp.Title, &temp.Content, &id, &timestamp)
		if err != nil {
			return nil, fmt.Errorf("[PostStorage]:Error with GetAllPost method in repository: %w", err)
		}
		temp.Author.ID, err = models.UserIDFromString(id)
		if err != nil {
			return nil, err
		}
		temp.CreatedAt = time.Unix(timestamp, 0)
		allPost = append(allPost, temp)
	}

	return allPost, nil
}

func (p *PostStorage) GetPostByID(id models.PostID) (models.Post, error) {
	row := p.db.QueryRow("SELECT ID,title,content,author,createdAt,like, dislike FROM post WHERE ID=$1", id)

	var temp models.Post
	var userIdStr string
	var timeStamp int64
	err := row.Scan(&temp.ID, &temp.Title, &temp.Content, &userIdStr, &timeStamp, &temp.Like, &temp.Dislike)
	if err != nil {
		return temp, fmt.Errorf("[PostStorage]:Error with GetPostByID method in repository: %w", err)
	}
	temp.Author.ID, err = models.UserIDFromString(userIdStr)
	if err != nil {
		return models.Post{}, err
	}
	temp.CreatedAt = time.Unix(timeStamp, 0)
	return temp, nil
}

func (p *PostStorage) GetPostsByUserID(uuid models.UserID) ([]models.Post, error) {
	row, err := p.db.Query("SELECT ID,title,content,createdAt,like,dislike FROM post WHERE author=$1", uuid.String())
	if err != nil {
		return nil, fmt.Errorf("[PostStorage]:Error with GetUsersPost method in repository: %w", err)
	}

	temp := models.Post{}
	usersPost := []models.Post{}

	for row.Next() {
		var timestamp int64
		err := row.Scan(&temp.ID, &temp.Title, &temp.Content, &timestamp, &temp.Like, &temp.Dislike)
		if err != nil {
			return nil, fmt.Errorf("[PostStorage]:Error with GetUsersPost method in repository: %w", err)
		}
		temp.CreatedAt = time.Unix(timestamp, 0)
		temp.Author.ID = uuid
		usersPost = append(usersPost, temp)
	}
	return usersPost, nil
}

func (p *PostStorage) GetUsersLikePosts(id models.UserID) ([]models.Post, error) {
	result := []models.Post{}

	rows, err := p.db.Query("SELECT ID,title,content,author,createdAt,like,dislike FROM post WHERE ID IN (SELECT postID FROM likePost WHERE userID = $1 AND status = $2)", id.String(), models.Like)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp models.Post
		var userIdStr string
		var timestamp int64
		err := rows.Scan(&temp.ID, &temp.Title, &temp.Content, &userIdStr, &timestamp, &temp.Like, &temp.Dislike)
		if err != nil {
			return nil, fmt.Errorf("[ReactionStorage]:Error with GetUsersLikePosts method in repository: %w", err)
		}
		temp.Author.ID, err = models.UserIDFromString(userIdStr)
		if err != nil {
			return nil, err
		}
		temp.CreatedAt = time.Unix(timestamp, 0)
		result = append(result, temp)
	}

	return result, nil
}

func (p *PostStorage) UpdatePost(post models.Post) error {
	stmt := `UPDATE post SET ID=$1,title=$2,content=$3,author=$4,createdat=$5,like=$6,dislike=$7 WHERE ID == $1`
	query, err := p.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = query.Exec(post.ID, post.Title, post.Content, post.Author.ID.String(), post.CreatedAt.Truncate(time.Second).Unix(), post.Like, post.Dislike)
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoriesStorage) GetPostsByCategory(category models.Category) ([]models.Post, error) {
	stmt := `SELECT ID, title, content, author, createdat, like, dislike FROM post WHERE id IN (SELECT postID FROM categoriesPost WHERE categoryID = $1)`
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	var res []models.Post
	values, err := query.Query(category.ID)
	if err != nil {
		return nil, err
	}
	for values.Next() {
		var post models.Post
		var userIdStr string
		var timestamp int64
		if err = values.Scan(&post.ID, &post.Title, &post.Content, &userIdStr, &timestamp, &post.Like, &post.Dislike); err != nil {
			return nil, err
		}
		post.Author.ID, err = models.UserIDFromString(userIdStr)
		if err != nil {
			return nil, err
		}
		post.CreatedAt = time.Unix(timestamp, 0)
		res = append(res, post)
	}
	return res, nil
}
