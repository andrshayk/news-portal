package repository

import (
	"context"
	"log"

	"gorm.io/gorm"
)

// --- NewsRepository ---
type NewsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (r *NewsRepository) GetByID(ctx context.Context, id int) (*News, error) {
	var news News
	if err := r.db.WithContext(ctx).First(&news, id).Error; err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] NEWS BY ID: %+v", news)
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

func (r *NewsRepository) GetAll(ctx context.Context, tagID, categoryID, limit, offset int) ([]News, error) {
	var newsList []News
	query := applyNewsFilters(r.db.WithContext(ctx).Model(&News{}), tagID, categoryID)
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

func (r *NewsRepository) Count(ctx context.Context, tagID, categoryID int) (int64, error) {
	var count int64
	query := applyNewsFilters(r.db.WithContext(ctx).Model(&News{}), tagID, categoryID)
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

func (r *NewsRepository) GetAllWithCategory(ctx context.Context, tagID, categoryID, limit, offset int) ([]NewsWithCategory, error) {
	var newsList []NewsWithCategory
	query := r.db.WithContext(ctx).
		Model(&News{}).
		Select(`news.*, 
			categories."categoryId" as category__category_id, 
			categories."tittle" as category__tittle, 
			categories."orderNumber" as category__order_number, 
			categories."statusId" as category__status_id`).
		Joins(`INNER JOIN categories ON news."categoryId" = categories."categoryId"`)
	query = applyNewsFilters(query, tagID, categoryID)

	if err := query.Order(`news."publishedAt" DESC`).Limit(limit).Offset(offset).Scan(&newsList).Error; err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] RESULT: %+v", newsList)
	return newsList, nil
}

func (r *NewsRepository) GetAllTags(ctx context.Context) ([]Tag, error) {
	var tags []Tag
	err := r.db.WithContext(ctx).
		Where(`"statusId" = ?`, 1).
		Order(`"tittle"`).
		Find(&tags).Error
	return tags, err
}

func (r *NewsRepository) GetTagByIDs(ctx context.Context, ids []int64) ([]Tag, error) {
	var tags []Tag
	err := r.db.WithContext(ctx).
		Where(`"statusId" = ?`, 1).
		Where(`"tagId" IN ?`, ids).
		Order(`"tittle"`).
		Find(&tags).Error
	return tags, err
}

func (r *NewsRepository) GetCategories(ctx context.Context) ([]Category, error) {
	var categories []Category
	err := r.db.WithContext(ctx).
		Where(`"statusId" = ?`, 1).
		Order(`"orderNumber"`).
		Find(&categories).Error
	return categories, err
}
