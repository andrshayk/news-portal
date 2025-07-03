package service

import (
	"context"
	"fmt"
	"news-portal/internal/repository"
)

type CategoryService struct {
	repo *repository.NewsRepository
}

func NewCategoryService(repo *repository.NewsRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (uc *CategoryService) GetAllCategories(ctx context.Context) ([]repository.Category, error) {
	return (*uc.repo).GetCategories(ctx)
}

func (uc *CategoryService) GetCategoryByID(ctx context.Context, id int) (*repository.Category, error) {
	categories, err := (*uc.repo).GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	for _, c := range categories {
		if c.CategoryID == id {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("category not found")
}
