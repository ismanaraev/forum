package repository

import (
	"fmt"
	"forumv2/internal/models"
)

func (p *PostStorage) GetCategoriesByPostID(postID models.PostID) ([]models.Category, error) {
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

func (p *PostStorage) FilterPostsByMultipleCategories(categories []models.Category) ([]models.Post, error) {
	if len(categories) == 0 {
		return nil, nil
	}
	var s string
	for i := 1; i < len(categories); i++ {
		s += fmt.Sprintf(" AND categoryID=%d", i+1)
	}
	var ids []interface{}
	for _, val := range categories {
		ids = append(ids, val.ID)
	}
	stmt, err := p.db.Prepare(`SELECT postID FROM categoriesPost WHERE categoryID = $1` + s)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(ids...)
	if err != nil {
		return nil, err
	}
	var postIDS []models.PostID
	for rows.Next() {
		var temp models.PostID
		err = rows.Scan(&temp)
		if err != nil {
			return nil, err
		}
		postIDS = append(postIDS, temp)
	}
	var res []models.Post
	for _, val := range postIDS {
		post, err := p.GetPostByID(val)
		if err != nil {
			return nil, err
		}
		res = append(res, post)
	}
	return res, nil
}

func (p *PostStorage) CreateCategory(name string) error {
	stmt, err := p.db.Prepare(`INSERT INTO categories (name) VALUES ($1)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name)
	return err
}

func (p *PostStorage) GetCategoryByName(name string) (models.Category, error) {
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