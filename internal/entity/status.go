package entity

type Status struct {
	StatusID int    `gorm:"column:statusId;primaryKey" json:"statusId"`
	Tittle   string `gorm:"column:tittle" json:"tittle"`
}
