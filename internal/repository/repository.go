package repository

import (
	"context"
	"log"

	"news-portal/internal/entity"

	"gorm.io/gorm"
)

// --- NewsRepository ---
type NewsRepository interface {
	GetByID(ctx context.Context, id int) (*entity.News, error)
	GetAll(ctx context.Context, tagID, categoryID, limit, offset int) ([]entity.News, error)
	Count(ctx context.Context, tagID, categoryID int) (int64, error)
	GetAllWithCategory(ctx context.Context, tagID, categoryID, limit, offset int) ([]entity.NewsWithCategory, error)
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

// applyNewsFilters применяет фильтры statusId, tagID, categoryID к запросу
func applyNewsFilters(query *gorm.DB, tagID, categoryID int) *gorm.DB {
	query = query.Where(`news."statusId" = ?`, 1)
	if tagID != 0 {
		query = query.Where(`? = ANY(news."tagIds")`, tagID)
	}
	if categoryID != 0 {
		query = query.Where(`news."categoryId" = ?`, categoryID)
	}
	return query
}

func (r *newsRepository) GetAll(ctx context.Context, tagID, categoryID, limit, offset int) ([]entity.News, error) {
	var newsList []entity.News
	query := applyNewsFilters(r.db.WithContext(ctx).Model(&entity.News{}), tagID, categoryID)
	// DEBUG
	stmt := query.Session(&gorm.Session{DryRun: true}).Order(`"publishedAt" DESC`).Limit(limit).Offset(offset).Find(&newsList).Statement
	log.Printf("[DEBUG] SQL: %s | Vars: %v", stmt.SQL.String(), stmt.Vars)

	if err := query.Order(`"publishedAt" DESC`).Limit(limit).Offset(offset).Find(&newsList).Error; err != nil {
		log.Printf("[ERROR] DB error: %v", err)
		return nil, err
	}

	log.Printf("[DEBUG] RESULT: %+v", newsList)
	return newsList, nil
}

func (r *newsRepository) Count(ctx context.Context, tagID, categoryID int) (int64, error) {
	var count int64
	query := applyNewsFilters(r.db.WithContext(ctx).Model(&entity.News{}), tagID, categoryID)
	// DEBUG
	stmt := query.Session(&gorm.Session{DryRun: true}).Count(&count).Statement
	log.Printf("[DEBUG] SQL: %s | Vars: %v", stmt.SQL.String(), stmt.Vars)

	if err := query.Count(&count).Error; err != nil {
		log.Printf("[ERROR] DB error: %v", err)
		return 0, err
	}

	log.Printf("[DEBUG] RESULT: %d", count)
	return count, nil
}

func (r *newsRepository) GetAllWithCategory(ctx context.Context, tagID, categoryID, limit, offset int) ([]entity.NewsWithCategory, error) {
	var newsList []entity.NewsWithCategory
	query := r.db.WithContext(ctx).
		Model(&entity.News{}).
		Select(`news.*, categories."categoryId", categories."tittle", categories."orderNumber", categories."statusId"`).
		Joins(`LEFT JOIN categories ON news."categoryId" = categories."categoryId"`)
	query = applyNewsFilters(query, tagID, categoryID)

	if err := query.Order(`news."publishedAt" DESC`).Limit(limit).Offset(offset).Scan(&newsList).Error; err != nil {
		return nil, err
	}
	return newsList, nil
}

// --- TagRepository ---
type TagRepository interface {
	GetAll(ctx context.Context) ([]entity.Tag, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.Tag, error)
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

func (r *tagRepository) GetByIDs(ctx context.Context, ids []int64) ([]entity.Tag, error) {
	var tags []entity.Tag
	err := r.db.WithContext(ctx).
		Where(`"statusId" = ?`, 1).
		Where(`"tagId" IN ?`, ids).
		Order(`"tittle"`).
		Find(&tags).Error
	return tags, err
}

// --- CategoryRepository ---
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
