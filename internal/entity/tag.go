package entity

type Tag struct {
	TagID    int    `gorm:"column:tagId;primaryKey"`
	Tittle   string `gorm:"column:tittle"`
	StatusID int    `gorm:"column:statusId"`
}

type TagResponse struct {
	TagID  int    `json:"tagId"`
	Tittle string `json:"tittle"`
}
