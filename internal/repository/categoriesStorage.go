package repository

import (
	"database/sql"
	"forumv2/internal/models"
	"time"
)

type categoriesStorage struct {
	db *sql.DB
}

func newCategoriesStorage(db *sql.DB) *categoriesStorage {
	return &categoriesStorage{db: db}
}

func (p *categoriesStorage) GetCategoriesByPostID(postID models.PostID) ([]models.Category, error) {
	var categories []models.Category
	query, err := p.db.Prepare(`SELECT id, name FROM categories WHERE id IN (SELECT categoryID FROM categoriesPost WHERE PostID = $1)`)
	if err != nil {
		return nil, err
	}
	categoriesQuery, err := query.Query(postID)
	if err != nil {
		return nil, err
	}
	for categoriesQuery.Next() {
		var cat models.Category
		err = categoriesQuery.Scan(&cat.ID, &cat.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}

func (p *categoriesStorage) CreateCategory(name string) error {
	stmt, err := p.db.Prepare(`INSERT INTO categories (name) VALUES ($1)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name)
	return err
}

func (p *categoriesStorage) GetCategoryByName(name string) (models.Category, error) {
	stmt, err := p.db.Prepare(`SELECT id,name FROM categories WHERE name = $1`)
	if err != nil {
		return models.Category{}, err
	}
	row := stmt.QueryRow(name)
	var res models.Category
	err = row.Scan(&res.ID, &res.Name)
	if err != nil {
		return models.Category{}, err
	}
	return res, nil
}

func (p *categoriesStorage) AddCategoryToPost(postId models.PostID, categoryId models.CategoryID) error {
	stmt, err := p.db.Prepare(`INSERT INTO categoriesPost (postID, categoryID) VALUES ($1, $2)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(postId, categoryId)
	return err
}

func (c *categoriesStorage) GetPostsByCategory(category models.Category) ([]models.Post, error) {
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

func (c *categoriesStorage) DeleteCategory(name string) error {
	stmt, err := c.db.Prepare(`DELETE FROM categories WHERE name = $1`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}
	return nil
}

func (c *categoriesStorage) GetAllCategories() ([]models.Category, error) {
	stmt, err := c.db.Prepare(`SELECT ID,name FROM categories`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var res []models.Category
	for rows.Next() {
		var temp models.Category
		err = rows.Scan(&temp.ID, &temp.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil
}
