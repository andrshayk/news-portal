package entity

import (
	"time"

	"github.com/lib/pq"
)

type News struct {
	NewsID      int           `gorm:"column:newsId;primaryKey"`
	Tittle      string        `gorm:"column:tittle"`
	ShortText   string        `gorm:"column:shortText"`
	FullText    string        `gorm:"column:fullText"`
	PublishedAt time.Time     `gorm:"column:publishedAt"`
	AuthorName  string        `gorm:"column:authorName"`
	CategoryID  int           `gorm:"column:categoryId"`
	TagIDs      pq.Int64Array `gorm:"column:tagIds;type:int[]"`
	CreatedAt   time.Time     `gorm:"column:createdAt"`
	StatusID    int           `gorm:"column:statusId"`
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

type NewsWithCategory struct {
	News
	Category Category `gorm:"foreignKey:CategoryID;references:CategoryID"`
}
