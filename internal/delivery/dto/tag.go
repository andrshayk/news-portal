package dto

import "news-portal/internal/entity"

type TagResponse struct {
	TagID  int    `json:"tagId"`
	Tittle string `json:"tittle"`
}

func ToTagResponse(tag entity.Tag) TagResponse {
	return TagResponse{
		TagID:  tag.TagID,
		Tittle: tag.Tittle,
	}
}

func ToTagResponseSlice(tags []entity.Tag) []TagResponse {
	resp := make([]TagResponse, len(tags))
	for i, tag := range tags {
		resp[i] = ToTagResponse(tag)
	}
	return resp
}
