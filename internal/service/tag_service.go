package service

import (
	"context"
	"news-portal/internal/entity"
	"news-portal/internal/repository"
)

type TagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) *TagService {
	return &TagService{repo: repo}
}

func (uc *TagService) GetAllTags(ctx context.Context) ([]entity.Tag, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *TagService) GetTagsByIDs(ctx context.Context, ids []int32) ([]entity.Tag, error) {
	allTags, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var tags []entity.Tag
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
