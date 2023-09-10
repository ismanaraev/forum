package repository

import (
	"database/sql"
	"fmt"
	"forumv2/internal/models"
	"log"
)

type PostStorage struct {
	db *sql.DB
}

func NewPostSQLite(db *sql.DB) *PostStorage {
	return &PostStorage{
		db: db,
	}
}

// Создать пост
func (p *PostStorage) CreatePost(post models.Post) (models.PostID, error) {
	query, err := p.db.Prepare(`INSERT INTO post(title,content,author,createdat) VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return 0, fmt.Errorf("[PostStorage]:Error with CreatePost method in repository: %w", err)
	}

	res, err := query.Exec(post.Title, post.Content, post.Author.ID, post.CreatedAt.Unix())
	if err != nil {
		return 0, fmt.Errorf("[PostStorage]:Error with CreatePost method in repository: %w", err)
	}
	postId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	for _, val := range post.Categories {
		query, err := p.db.Prepare(`INSERT INTO categoriesPost (postID, categoryID) VALUES ($1,$2)`)
		if err != nil {
			return 0, err
		}
		_, err = query.Exec(postId, val.ID)
		if err != nil {
			return 0, err
		}
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
		var author models.User
		err := row.Scan(&temp.ID, &temp.Title, &temp.Content, &author.ID, &temp.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("[PostStorage]:Error with GetAllPost method in repository: %w", err)
		}
		categories, err := p.GetCategoriesByPostID(temp.ID)
		if err != nil {
			return nil, err
		}
		temp.Categories = categories
		allPost = append(allPost, temp)
	}

	return allPost, nil
}

func (p *PostStorage) GetPostByID(id models.PostID) (models.Post, error) {
	row := p.db.QueryRow("SELECT ID,title,content,author,createdAt,categories, like, dislike FROM post WHERE ID=$1", id)

	var temp models.Post
	err := row.Scan(&temp.ID, &temp.Author.ID, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Categories, &temp.Like, &temp.Dislike)
	if err != nil {
		return temp, fmt.Errorf("[PostStorage]:Error with GetPostByID method in repository: %w", err)
	}
	var categories []models.Category
	categories, err = p.GetCategoriesByPostID(temp.ID)
	temp.Categories = categories
	return temp, nil
}

func (p *PostStorage) GetPostsByUserID(uuid models.UserID) ([]models.Post, error) {
	row, err := p.db.Query("SELECT ID,title,content,createdAt,like,dislike FROM post WHERE author=$1", uuid)
	if err != nil {
		return nil, fmt.Errorf("[PostStorage]:Error with GetUsersPost method in repository: %w", err)
	}

	temp := models.Post{}
	usersPost := []models.Post{}

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.Title, &temp.Content, &temp.CreatedAt, &temp.Like, &temp.Dislike)
		if err != nil {
			return nil, fmt.Errorf("[PostStorage]:Error with GetUsersPost method in repository: %w", err)
		}
		usersPost = append(usersPost, temp)
	}
	return usersPost, nil
}

// Запрос на ID поста юзера
func (p *PostStorage) GetPostIdWithUUID(uuid models.UserID) ([]models.PostID, error) {
	row, err := p.db.Query("SELECT postID FROM likePost WHERE userID==$1 AND status==$2", uuid, 1)
	if err != nil {
		return nil, fmt.Errorf("[PostStorage]:Error with GetPostIdWithUUID method in repository: %w", err)
	}

	var result []models.PostID

	for row.Next() {
		var temp models.PostID
		err = row.Scan(&temp)
		if err != nil {
			return nil, fmt.Errorf("[PostStorage]:Error with GetPostIdWithUUID method in repository: %w", err)
		}
		result = append(result, temp)
	}

	return result, nil
}

func (p *PostStorage) GetUsersLikePosts(id models.UserID) ([]models.Post, error) {
	result := []models.Post{}

	rows, err := p.db.Query("SELECT id,title,content,author,createdAt,like,dislike FROM post WHERE id IN (SELECT postID FROM likePost WHERE userID = $1)", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var temp models.Post
		err := rows.Scan(&temp.ID, &temp.Title, &temp.Content, &temp.Author, &temp.CreatedAt, &temp.Like, &temp.Dislike)
		if err != nil {
			return nil, fmt.Errorf("[ReactionStorage]:Error with GetUsersLikePosts method in repository: %w", err)
		}
		result = append(result, temp)
	}

	return result, nil
}

func (p *PostStorage) UpdatePost(post models.Post) error {
	stmt := `UPDATE post SET ID=$1,title=$2,content=$3,author=$4,createdat=$5,like=$6,dislike=$7 WHERE id == $1`
	query, err := p.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = query.Exec(&post.ID, &post.Title, &post.Content, &post.Author, &post.CreatedAt, &post.Like, &post.Dislike)
	if err != nil {
		return err
	}
	return nil
}

func (c *PostStorage) GetPostsByCategory(category models.Category) ([]models.Post, error) {
	stmt := `SELECT id, title, content, author, createdat, (SELECT name FROM categories INNER JOIN categoriesPost ON categories.ID = categoriesPost.categoryID) AS categories, like, dislike FROM post WHERE `
	query, err := c.db.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	var res []models.Post
	values, err := query.Query(category)
	if err != nil {
		return nil, err
	}
	for values.Next() {
		var post models.Post
		if err := values.Scan(&post.ID, &post.Title, &post.Content, &post.Author.ID, &post.CreatedAt, &post.Categories, &post.Like, &post.Dislike); err != nil {
			return nil, err
		}
		res = append(res, post)
	}
	return res, nil
}
