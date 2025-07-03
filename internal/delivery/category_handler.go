package delivery

import (
	"net/http"
	"news-portal/internal/entity"
	"news-portal/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	uc *service.CategoryService
}

// Регистрируем эндпоинт GET /api/categories
func RegisterCategoryRoutes(rg *gin.RouterGroup, uc *service.CategoryService) {
	h := &CategoryHandler{uc: uc}
	rg.GET("/categories", h.GetAll)
}

// Получение всех категорий
func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.uc.GetAllCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Преобразуем в []CategoryResponse
	resp := make([]entity.CategoryResponse, len(categories))
	for i, cat := range categories {
		resp[i] = entity.CategoryResponse{
			CategoryID:  cat.CategoryID,
			Tittle:      cat.Tittle,
			OrderNumber: cat.OrderNumber,
		}
	}
	c.JSON(http.StatusOK, resp)
}
