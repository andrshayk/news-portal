package delivery

import (
	"net/http"
	"strconv"

	"news-portal/internal/entity"
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
	tagID, _ := strconv.Atoi(c.Query("tagId"))
	categoryID, _ := strconv.Atoi(c.Query("categoryId"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * pageSize

	newsList, err := h.uc.GetAllNews(c.Request.Context(), tagID, categoryID, pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []entity.NewsResponse
	for _, n := range newsList {
		category, _ := h.categoryService.GetCategoryByID(c.Request.Context(), n.CategoryID)
		tagIDs := make([]int32, len(n.TagIDs))
		for i, v := range n.TagIDs {
			tagIDs[i] = int32(v)
		}
		tags, _ := h.tagService.GetTagsByIDs(c.Request.Context(), tagIDs)

		// Преобразуем category
		catResp := entity.CategoryResponse{
			CategoryID:  category.CategoryID,
			Tittle:      category.Tittle,
			OrderNumber: category.OrderNumber,
		}

		// Преобразуем tags
		tagResps := make([]entity.TagResponse, len(tags))
		for i, tag := range tags {
			tagResps[i] = entity.TagResponse{
				TagID:  tag.TagID,
				Tittle: tag.Tittle,
			}
		}

		resp = append(resp, entity.NewsResponse{
			NewsID:      n.NewsID,
			Tittle:      n.Tittle,
			ShortText:   n.ShortText,
			FullText:    n.FullText,
			PublishedAt: n.PublishedAt,
			AuthorName:  n.AuthorName,
			Category:    catResp,
			Tags:        tagResps,
		})
	}

	c.JSON(http.StatusOK, resp)
}

// Получение кол-ва новостей
func (h *NewsHandler) GetNewsCount(c *gin.Context) {
	tagID, _ := strconv.Atoi(c.Query("tagId"))
	categoryID, _ := strconv.Atoi(c.Query("categoryId"))

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

	n, err := h.uc.GetNewsByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	category, _ := h.categoryService.GetCategoryByID(c.Request.Context(), n.CategoryID)
	tagIDs := make([]int32, len(n.TagIDs))
	for i, v := range n.TagIDs {
		tagIDs[i] = int32(v)
	}
	tags, _ := h.tagService.GetTagsByIDs(c.Request.Context(), tagIDs)

	catResp := entity.CategoryResponse{
		CategoryID:  category.CategoryID,
		Tittle:      category.Tittle,
		OrderNumber: category.OrderNumber,
	}
	tagResps := make([]entity.TagResponse, len(tags))
	for i, tag := range tags {
		tagResps[i] = entity.TagResponse{
			TagID:  tag.TagID,
			Tittle: tag.Tittle,
		}
	}

	resp := entity.NewsResponse{
		NewsID:      n.NewsID,
		Tittle:      n.Tittle,
		ShortText:   n.ShortText,
		FullText:    n.FullText,
		PublishedAt: n.PublishedAt,
		AuthorName:  n.AuthorName,
		Category:    catResp,
		Tags:        tagResps,
	}

	c.JSON(http.StatusOK, resp)
}
