package entity

type Category struct {
	CategoryID  int    `gorm:"column:categoryId;primaryKey"`
	Tittle      string `gorm:"column:tittle"`
	OrderNumber int    `gorm:"column:orderNumber"`
	StatusID    int    `gorm:"column:statusId"`
}

type CategoryResponse struct {
	CategoryID  int    `json:"categoryId"`
	Tittle      string `json:"tittle"`
	OrderNumber int    `json:"orderNumber"`
}
