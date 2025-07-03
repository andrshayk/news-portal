package service

import (
	"context"
	"news-portal/internal/repository"
	"news-portal/internal/response"
)

type NewsService struct {
	newsRepo        *repository.NewsRepository
	categoryService *CategoryService
	tagService      *TagService
}

func NewNewsService(newsRepo *repository.NewsRepository, categoryService *CategoryService, tagService *TagService) *NewsService {
	return &NewsService{
		newsRepo:        newsRepo,
		categoryService: categoryService,
		tagService:      tagService,
	}
}

func (uc *NewsService) GetAllNews(ctx context.Context, tagId, categoryId, limit, offset int) ([]repository.News, error) {
	return uc.newsRepo.GetAll(ctx, tagId, categoryId, limit, offset)
}

func (uc *NewsService) CountNews(ctx context.Context, tagId, categoryId int) (int64, error) {
	return uc.newsRepo.Count(ctx, tagId, categoryId)
}

func (uc *NewsService) GetNewsByID(ctx context.Context, id int) (*repository.News, error) {
	return uc.newsRepo.GetByID(ctx, id)
}

func (uc *NewsService) GetAllNewsWithCategory(ctx context.Context, tagId, categoryId, limit, offset int) ([]repository.NewsWithCategory, error) {
	return uc.newsRepo.GetAllWithCategory(ctx, tagId, categoryId, limit, offset)
}

func (uc *NewsService) GetNewsResponses(ctx context.Context, tagId, categoryId, limit, offset int) ([]response.NewsResponse, error) {
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

	var tags []repository.Tag
	if len(tagIds) > 0 {
		tags, err = uc.tagService.GetTagsByIDsFast(ctx, tagIds)
		if err != nil {
			return nil, err
		}
	}
	tagMap := make(map[int64]repository.Tag)
	for _, t := range tags {
		tagMap[int64(t.TagID)] = t
	}

	var resp []response.NewsResponse
	for _, n := range newsList {
		catResp := response.ToCategoryResponse(n.Category.ToCategory())
		// Собираем теги для новости
		tagResps := make([]response.TagResponse, 0, len(n.TagIDs))
		for _, tagId := range n.TagIDs {
			if tag, ok := tagMap[tagId]; ok {
				tagResps = append(tagResps, response.ToTagResponse(tag))
			}
		}
		resp = append(resp, response.ToNewsResponse(n.News, catResp, tagResps))
	}
	return resp, nil
}

func (uc *NewsService) GetNewsResponseByID(ctx context.Context, id int) (*response.NewsResponse, error) {
	n, err := uc.GetNewsByID(ctx, id)
	if err != nil {
		return nil, err
	}
	category, err := uc.categoryService.GetCategoryByID(ctx, n.CategoryID)
	if err != nil {
		return nil, err
	}
	tagIDs := make([]int32, len(n.TagIDs))
	for i, v := range n.TagIDs {
		tagIDs[i] = int32(v)
	}
	var tags []repository.Tag
	if len(tagIDs) > 0 {
		tags, err = uc.tagService.GetTagsByIDs(ctx, tagIDs)
		if err != nil {
			return nil, err
		}
	}
	catResp := response.ToCategoryResponse(*category)
	tagResps := response.ToTagResponseSlice(tags)
	resp := response.ToNewsResponse(*n, catResp, tagResps)
	return &resp, nil
}
