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
	tagIDStr := c.Query("tagId")
	var tagID int
	if tagIDStr != "" {
		var err error
		tagID, err = strconv.Atoi(tagIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tagId"})
			return
		}
	}
	categoryIDStr := c.Query("categoryId")
	var categoryID int
	if categoryIDStr != "" {
		var err error
		categoryID, err = strconv.Atoi(categoryIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid categoryId"})
			return
		}
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	offset := (page - 1) * pageSize

	resp, err := h.uc.GetNewsResponses(c.Request.Context(), tagID, categoryID, pageSize, offset, h.categoryService, h.tagService)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Получение кол-ва новостей
func (h *NewsHandler) GetNewsCount(c *gin.Context) {
	tagIDStr := c.Query("tagId")
	var tagID int
	if tagIDStr != "" {
		var err error
		tagID, err = strconv.Atoi(tagIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tagId"})
			return
		}
	}
	categoryIDStr := c.Query("categoryId")
	var categoryID int
	if categoryIDStr != "" {
		var err error
		categoryID, err = strconv.Atoi(categoryIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid categoryId"})
			return
		}
	}

	count, err := h.uc.CountNews(c.Request.Context(), tagID, categoryID)
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

	resp, err := h.uc.GetNewsResponseByID(c.Request.Context(), id, h.categoryService, h.tagService)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
