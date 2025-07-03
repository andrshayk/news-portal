package delivery

import (
	"net/http"
	"strconv"

	"news-portal/internal/service"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	uc              *service.NewsService
	categoryService *service.CategoryService
	tagService      *service.TagService
}

// Регистрируем эндпоинт GET /api/news
func RegisterNewsRoutes(rg *gin.RouterGroup, uc *service.NewsService, categoryService *service.CategoryService, tagService *service.TagService) {
	h := &NewsHandler{uc: uc, categoryService: categoryService, tagService: tagService}
	news := rg.Group("/news")

	news.GET("/", h.GetNews)
	news.GET("/count", h.GetNewsCount)
	news.GET("/:id", h.GetNewsByID)
}

// Получение всех новостей
func (h *NewsHandler) GetNews(c *gin.Context) {
	type NewsQuery struct {
		TagID      int `form:"tagId"`
		CategoryID int `form:"categoryId"`
		Page       int `form:"page,default=1"`
		Limit      int `form:"limit,default=10"`
	}
	var q NewsQuery
	if err := c.ShouldBind(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if q.Page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}
	if q.Limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}
	offset := (q.Page - 1) * q.Limit

	resp, err := h.uc.GetNewsResponses(c.Request.Context(), q.TagID, q.CategoryID, q.Limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Получение кол-ва новостей
func (h *NewsHandler) GetNewsCount(c *gin.Context) {
	type NewsCountQuery struct {
		TagID      int `form:"tagId"`
		CategoryID int `form:"categoryId"`
	}
	var q NewsCountQuery
	if err := c.ShouldBind(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := h.uc.CountNews(c.Request.Context(), q.TagID, q.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

// Получение новости по id
func (h *NewsHandler) GetNewsByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	resp, err := h.uc.GetNewsResponseByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
