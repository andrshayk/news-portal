package entity

type Category struct {
	CategoryID  int    `gorm:"column:categoryId;primaryKey" json:"categoryId"`
	Tittle      string `gorm:"column:tittle" json:"tittle"`
	OrderNumber int    `gorm:"column:orderNumber" json:"orderNumber"`
	StatusID    int    `gorm:"column:statusId" json:"statusId"`
}

type CategoryResponse struct {
	CategoryID  int    `json:"categoryId"`
	Tittle      string `json:"tittle"`
	OrderNumber int    `json:"orderNumber"`
}
