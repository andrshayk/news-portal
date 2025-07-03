package repository

import (
	"context"

	"news-portal/internal/entity"

	"fmt"

	"gorm.io/gorm"
)

type NewsRepository interface {
	GetByID(ctx context.Context, id int) (*entity.News, error)
	GetAll(ctx context.Context, tagID, categoryID, limit, offset int) ([]entity.News, error)
	Count(ctx context.Context, tagID, categoryID int) (int64, error)
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db: db}
}

func (r *newsRepository) GetByID(ctx context.Context, id int) (*entity.News, error) {
	var news entity.News
	if err := r.db.WithContext(ctx).First(&news, id).Error; err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *newsRepository) GetAll(ctx context.Context, tagID, categoryID, limit, offset int) ([]entity.News, error) {
	var newsList []entity.News
	query := r.db.WithContext(ctx).Model(&entity.News{}).Where(`"statusId" = ?`, 1)
	fmt.Printf("[DEBUG] Query: %+v\n", query)

	if tagID != 0 {
		query = query.Where(`? = ANY("tagIds")`, tagID)
	}
	if categoryID != 0 {
		query = query.Where(`"categoryId" = ?`, categoryID)
	}

	// DEBUG
	stmt := query.Session(&gorm.Session{DryRun: true}).Order(`"publishedAt" DESC`).Limit(limit).Offset(offset).Find(&newsList).Statement
	fmt.Printf("[DEBUG] SQL: %s | Vars: %v\n", stmt.SQL.String(), stmt.Vars)

	if err := query.Order(`"publishedAt" DESC`).Limit(limit).Offset(offset).Find(&newsList).Error; err != nil {
		fmt.Printf("[DEBUG] ERROR: %v\n", err)
		return nil, err
	}

	fmt.Printf("[DEBUG] RESULT: %+v\n", newsList)
	return newsList, nil
}

func (r *newsRepository) Count(ctx context.Context, tagID, categoryID int) (int64, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&entity.News{}).Where(`"statusId" = ?`, 1)

	if tagID != 0 {
		query = query.Where(`? = ANY("tagIds")`, tagID)
	}
	if categoryID != 0 {
		query = query.Where(`"categoryId" = ?`, categoryID)
	}

	// DEBUG
	stmt := query.Session(&gorm.Session{DryRun: true}).Count(&count).Statement
	fmt.Printf("[DEBUG] SQL: %s | Vars: %v\n", stmt.SQL.String(), stmt.Vars)

	if err := query.Count(&count).Error; err != nil {
		fmt.Printf("[DEBUG] ERROR: %v\n", err)
		return 0, err
	}

	fmt.Printf("[DEBUG] RESULT: %d\n", count)
	return count, nil
}
