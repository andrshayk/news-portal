package repository

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

type NewsWithCategory struct {
	News
	Category CategoryJoin `gorm:"embedded"`
}

type Category struct {
	CategoryID  int    `gorm:"column:categoryId"`
	Tittle      string `gorm:"column:tittle"`
	OrderNumber int    `gorm:"column:orderNumber"`
	StatusID    int    `gorm:"column:statusId"`
}

type CategoryJoin struct {
	CategoryID  int    `gorm:"column:category__category_id"`
	Tittle      string `gorm:"column:category__tittle"`
	OrderNumber int    `gorm:"column:category__order_number"`
	StatusID    int    `gorm:"column:category__status_id"`
}

func (cj CategoryJoin) ToCategory() Category {
	return Category{
		CategoryID:  cj.CategoryID,
		Tittle:      cj.Tittle,
		OrderNumber: cj.OrderNumber,
		StatusID:    cj.StatusID,
	}
}

type Tag struct {
	TagID    int    `gorm:"column:tagId;primaryKey"`
	Tittle   string `gorm:"column:tittle"`
	StatusID int    `gorm:"column:statusId"`
}

type Status struct {
	StatusID int    `gorm:"column:statusId;primaryKey" json:"statusId"`
	Tittle   string `gorm:"column:tittle" json:"tittle"`
}
