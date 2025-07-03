package service

import (
	"context"
	"fmt"
	"news-portal/internal/delivery/dto"
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

func (uc *NewsService) GetAllNewsWithCategory(ctx context.Context, tagId, categoryId, limit, offset int) ([]entity.NewsWithCategory, error) {
	if repoWithCat, ok := uc.repo.(interface {
		GetAllWithCategory(context.Context, int, int, int, int) ([]entity.NewsWithCategory, error)
	}); ok {
		return repoWithCat.GetAllWithCategory(ctx, tagId, categoryId, limit, offset)
	}
	return nil, fmt.Errorf("repo does not support GetAllWithCategory")
}

func (uc *NewsService) GetNewsResponses(ctx context.Context, tagId, categoryId, limit, offset int, categoryService *CategoryService, tagService *TagService) ([]dto.NewsResponse, error) {
	newsList, err := uc.GetAllNewsWithCategory(ctx, tagId, categoryId, limit, offset)
	if err != nil {
		return nil, err
	}

	// Собираем все уникальные tagIds
	tagIdSet := make(map[int64]struct{})
	for _, n := range newsList {
		for _, tagId := range n.TagIDs {
			tagIdSet[tagId] = struct{}{}
		}
	}
	var tagIds []int64
	for id := range tagIdSet {
		tagIds = append(tagIds, id)
	}

	tags, _ := tagService.GetTagsByIDsFast(ctx, tagIds)
	tagMap := make(map[int64]entity.Tag)
	for _, t := range tags {
		tagMap[int64(t.TagID)] = t
	}

	var resp []dto.NewsResponse
	for _, n := range newsList {
		catResp := dto.ToCategoryResponse(n.Category)
		// Собираем теги для новости
		tagResps := make([]dto.TagResponse, 0, len(n.TagIDs))
		for _, tagId := range n.TagIDs {
			if tag, ok := tagMap[tagId]; ok {
				tagResps = append(tagResps, dto.ToTagResponse(tag))
			}
		}
		resp = append(resp, dto.ToNewsResponse(n.News, catResp, tagResps))
	}
	return resp, nil
}

func (uc *NewsService) GetNewsResponseByID(ctx context.Context, id int, categoryService *CategoryService, tagService *TagService) (*dto.NewsResponse, error) {
	n, err := uc.GetNewsByID(ctx, id)
	if err != nil {
		return nil, err
	}
	category, _ := categoryService.GetCategoryByID(ctx, n.CategoryID)
	tagIDs := make([]int32, len(n.TagIDs))
	for i, v := range n.TagIDs {
		tagIDs[i] = int32(v)
	}
	tags, _ := tagService.GetTagsByIDs(ctx, tagIDs)
	catResp := dto.ToCategoryResponse(*category)
	tagResps := dto.ToTagResponseSlice(tags)
	resp := dto.ToNewsResponse(*n, catResp, tagResps)
	return &resp, nil
}
