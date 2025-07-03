package dto

import "news-portal/internal/entity"

type CategoryResponse struct {
	CategoryID  int    `json:"categoryId"`
	Tittle      string `json:"tittle"`
	OrderNumber int    `json:"orderNumber"`
}

func ToCategoryResponse(cat entity.Category) CategoryResponse {
	return CategoryResponse{
		CategoryID:  cat.CategoryID,
		Tittle:      cat.Tittle,
		OrderNumber: cat.OrderNumber,
	}
}

func ToCategoryResponseSlice(cats []entity.Category) []CategoryResponse {
	resp := make([]CategoryResponse, len(cats))
	for i, cat := range cats {
		resp[i] = ToCategoryResponse(cat)
	}
	return resp
}
