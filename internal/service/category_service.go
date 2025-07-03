package service

import (
	"context"
	"fmt"
	"news-portal/internal/entity"
	"news-portal/internal/repository"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (uc *CategoryService) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *CategoryService) GetCategoryByID(ctx context.Context, id int) (*entity.Category, error) {
	categories, err := uc.repo.GetAll(ctx)
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
