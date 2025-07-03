package delivery

import (
	"net/http"
	"news-portal/internal/response"
	"news-portal/internal/service"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	uc *service.TagService
}

// Регистрируем эндпоинт GET /api/tags
func RegisterTagRoutes(rg *gin.RouterGroup, uc *service.TagService) {
	h := &TagHandler{uc: uc}
	rg.GET("/tags", h.GetAll)
}

// Получение всех тегов
func (h *TagHandler) GetAll(c *gin.Context) {
	tags, err := h.uc.GetAllTags(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Преобразуем в []TagResponse
	resp := response.ToTagResponseSlice(tags)
	c.JSON(http.StatusOK, resp)
}
