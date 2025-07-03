package dto

import (
	"news-portal/internal/entity"
	"time"
)

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

func ToNewsResponse(n entity.News, category CategoryResponse, tags []TagResponse) NewsResponse {
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
