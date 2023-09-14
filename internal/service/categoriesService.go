package service

import (
	"database/sql"
	"errors"
	"forumv2/internal/models"
)

type categoriesService struct {
	repo Repository
}

func newCategoriesService(repo Repository) *categoriesService {
	return &categoriesService{
		repo: repo,
	}
}

func (c *categoriesService) CreateCategory(name string) error {
	cat, err := c.repo.GetCategoryByName(name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}
	if cat != nil {
		return errors.New("category exists")
	}
	return c.repo.CreateCategory(name)
}

func (c *categoriesService) DeleteCategory(name string) error {
	return c.repo.DeleteCategory(name)
}

func (c *categoriesService) GetAllCategories() ([]models.Category, error) {
	return c.repo.GetAllCategories()
}
