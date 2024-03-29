package services

import (
	"github.com/Ghalbavieira/e-commerce.git/internal/database"
	"github.com/Ghalbavieira/e-commerce.git/internal/entity"
)

type CategoryService struct {
	categoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{categoryDB: categoryDB}
}

func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.categoryDB.GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	_, err := cs.categoryDB.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := cs.categoryDB.GetCategory(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

