package repository

import (
	"context"

	"news-portal/internal/entity"

	"gorm.io/gorm"
)

type TagRepository interface {
	GetAll(ctx context.Context) ([]entity.Tag, error)
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) GetAll(ctx context.Context) ([]entity.Tag, error) {
	var tags []entity.Tag
	err := r.db.WithContext(ctx).
		Where(`"statusId" = ?`, 1).
		Order(`"tittle"`).
		Find(&tags).Error
	return tags, err
}
