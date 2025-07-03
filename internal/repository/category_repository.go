package repository

import (
	"context"

	"news-portal/internal/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.WithContext(ctx).
		Where(`"statusId" = ?`, 1).
		Order(`"orderNumber"`).
		Find(&categories).Error
	return categories, err
}
