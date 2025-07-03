package entity

type Tag struct {
	TagID    int    `gorm:"column:tagId;primaryKey" json:"tagId"`
	Tittle   string `gorm:"column:tittle" json:"tittle"`
	StatusID int    `gorm:"column:statusId" json:"statusId"`
}

type TagResponse struct {
	TagID  int    `json:"tagId"`
	Tittle string `json:"tittle"`
}
