package transport

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	service services.CategoryService
	logger  *slog.Logger
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	logger := slog.Default()
	return &CategoryHandler{
		service: service,
		logger:  logger,
	}
}

func (h *CategoryHandler) RegisterRoutes(r *gin.Engine) {
	categories := r.Group("/categories")
	{
		categories.GET("", h.GetAll)
		categories.POST("", h.Create)
		categories.POST("/subcategory", h.CreateSubcategory)
		categories.GET("/:id/subcategory", h.GetSubcategoriesByCategoryID)

	}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	h.logger.Info("category_handler.Create: processing POST /categories request")

	var req models.CreateCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("category_handler.Create: invalid request body", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Debug("category_handler.Create: attempting to create category", slog.String("name", req.Name))

	category, err := h.service.CreateCategory(req)
	if err != nil {
		h.logger.Error("category_handler.Create: service error", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("category_handler.Create: category created successfully", slog.Uint64("id", uint64(category.ID)))
	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	h.logger.Info("category_handler.GetAll: processing GET /categories request")

	categories, err := h.service.GetAll()
	if err != nil {
		h.logger.Error("category_handler.GetAll: service error", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("category_handler.GetAll: categories retrieved successfully", slog.Int("count", len(categories)))
	c.JSON(http.StatusOK, categories)
}
func (h *CategoryHandler) CreateSubcategory(c *gin.Context) {
	h.logger.Info("category_handler.CreateSubcategory: processing POST /categories/subcategory request")

	var req models.CreateSubcategory
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("category_handler.CreateSubcategory: invalid request body", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Debug("category_handler.CreateSubcategory: attempting to create subcategory", slog.String("name", req.Name), slog.Uint64("category_id", uint64(req.CategoryID)))

	subcategory, err := h.service.CreateSubcategory(req)
	if err != nil {
		h.logger.Error("category_handler.CreateSubcategory: service error", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("category_handler.CreateSubcategory: subcategory created successfully", slog.Uint64("id", uint64(subcategory.ID)))
	c.JSON(http.StatusCreated, subcategory)
}
func (h *CategoryHandler) GetSubcategoriesByCategoryID(c *gin.Context) {
	h.logger.Info("category_handler.GetSubcategoriesByCategoryID: processing GET /categories/:id/subcategory request")

	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Error("category_handler.GetSubcategoriesByCategoryID: invalid id parameter", slog.String("id", idStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Debug("category_handler.GetSubcategoriesByCategoryID: fetching subcategories", slog.Uint64("category_id", id))

	subcategory, err := h.service.GetSubcategoriesByCategoryID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Warn("category_handler.GetSubcategoriesByCategoryID: category not found", slog.Uint64("category_id", id))
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.logger.Error("category_handler.GetSubcategoriesByCategoryID: service error", slog.String("error", err.Error()), slog.Uint64("category_id", id))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("category_handler.GetSubcategoriesByCategoryID: subcategories retrieved successfully", slog.Uint64("category_id", id), slog.Int("count", len(subcategory)))
	c.JSON(http.StatusOK, subcategory)
}
