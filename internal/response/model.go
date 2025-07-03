package response

import (
	"news-portal/internal/repository"
	"time"
)

type CategoryResponse struct {
	CategoryID  int    `json:"categoryId"`
	Tittle      string `json:"tittle"`
	OrderNumber int    `json:"orderNumber"`
}

func ToCategoryResponse(cat repository.Category) CategoryResponse {
	return CategoryResponse{
		CategoryID:  cat.CategoryID,
		Tittle:      cat.Tittle,
		OrderNumber: cat.OrderNumber,
	}
}

func ToCategoryResponseSlice(cats []repository.Category) []CategoryResponse {
	resp := make([]CategoryResponse, len(cats))
	for i, cat := range cats {
		resp[i] = ToCategoryResponse(cat)
	}
	return resp
}

type NewsResponse struct {
	NewsID      int              `json:"newsId"`
	Tittle      string           `json:"tittle"`
	ShortText   string           `json:"shortText"`
	FullText    string           `json:"fullText,omitempty"`
	PublishedAt time.Time        `json:"publishedAt"`
	AuthorName  string           `json:"authorName"`
	Category    CategoryResponse `json:"category"`
	Tags        []TagResponse    `json:"tags"`
}

func ToNewsResponse(n repository.News, category CategoryResponse, tags []TagResponse) NewsResponse {
	return NewsResponse{
		NewsID:      n.NewsID,
		Tittle:      n.Tittle,
		ShortText:   n.ShortText,
		FullText:    n.FullText,
		PublishedAt: n.PublishedAt,
		AuthorName:  n.AuthorName,
		Category:    category,
		Tags:        tags,
	}
}

type TagResponse struct {
	TagID  int    `json:"tagId"`
	Tittle string `json:"tittle"`
}

func ToTagResponse(tag repository.Tag) TagResponse {
	return TagResponse{
		TagID:  tag.TagID,
		Tittle: tag.Tittle,
	}
}

func ToTagResponseSlice(tags []repository.Tag) []TagResponse {
	resp := make([]TagResponse, len(tags))
	for i, tag := range tags {
		resp[i] = ToTagResponse(tag)
	}
	return resp
}
