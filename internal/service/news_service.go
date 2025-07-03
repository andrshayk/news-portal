package service

import (
	"context"
	"news-portal/internal/entity"
	"news-portal/internal/repository"
)

type NewsService struct {
	repo repository.NewsRepository
}

func NewNewsService(repo repository.NewsRepository) *NewsService {
	return &NewsService{repo: repo}
}

func (uc *NewsService) GetAllNews(ctx context.Context, tagId, categoryId, limit, offset int) ([]entity.News, error) {
	return uc.repo.GetAll(ctx, tagId, categoryId, limit, offset)
}

func (uc *NewsService) CountNews(ctx context.Context, tagId, categoryId int) (int64, error) {
	return uc.repo.Count(ctx, tagId, categoryId)
}

func (uc *NewsService) GetNewsByID(ctx context.Context, id int) (*entity.News, error) {
	return uc.repo.GetByID(ctx, id)
}
