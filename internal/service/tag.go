package service

import (
	"context"
	"news-portal/internal/repository"
)

type TagService struct {
	repo *repository.NewsRepository
}

func NewTagService(repo *repository.NewsRepository) *TagService {
	return &TagService{repo: repo}
}

func (uc *TagService) GetAllTags(ctx context.Context) ([]repository.Tag, error) {
	return (*uc.repo).GetAllTags(ctx)
}

func (uc *TagService) GetTagsByIDs(ctx context.Context, ids []int32) ([]repository.Tag, error) {
	allTags, err := (*uc.repo).GetAllTags(ctx)
	if err != nil {
		return nil, err
	}
	var tags []repository.Tag
	for _, tag := range allTags {
		for _, id := range ids {
			if int32(tag.TagID) == id {
				tags = append(tags, tag)
				break
			}
		}
	}
	return tags, nil
}

func (uc *TagService) GetTagsByIDsFast(ctx context.Context, ids []int64) ([]repository.Tag, error) {
	return (*uc.repo).GetTagByIDs(ctx, ids)
}
