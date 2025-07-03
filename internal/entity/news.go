package entity

import (
	"time"

	"github.com/lib/pq"
)

type News struct {
	NewsID      int           `gorm:"column:newsId;primaryKey" json:"newsId"`
	Tittle      string        `gorm:"column:tittle" json:"tittle"`
	ShortText   string        `gorm:"column:shortText" json:"shortText"`
	FullText    string        `gorm:"column:fullText" json:"fullText,omitempty"`
	PublishedAt time.Time     `gorm:"column:publishedAt" json:"publishedAt"`
	AuthorName  string        `gorm:"column:authorName" json:"authorName"`
	CategoryID  int           `gorm:"column:categoryId" json:"categoryId"`
	TagIDs      pq.Int64Array `gorm:"column:tagIds;type:int[]" json:"tagIds"`
	CreatedAt   time.Time     `gorm:"column:createdAt" json:"createdAt"`
	StatusID    int           `gorm:"column:statusId" json:"statusId"`
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
