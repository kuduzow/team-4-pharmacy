package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
	"gorm.io/gorm"
)

type PromocodeHandler struct {
	service services.PromocodeService
}

func NewPromocodeHandler(
	service services.PromocodeService,
) *PromocodeHandler {
	return &PromocodeHandler{service: service}
}

func (h *PromocodeHandler) RegisterRoutes(r *gin.Engine) {
	promocodes := r.Group("/promocodes")
	{
		promocodes.POST("", h.Create)
		promocodes.GET("", h.GetAll)
		promocodes.PATCH("/:id", h.Update)
		promocodes.DELETE("/:id", h.Delete)
	}
}

func (h *PromocodeHandler) Create(c *gin.Context) {
	var req models.PromocodeCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный JSON"})
		return
	}

	promocode, err := h.service.CreatePromocode(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, promocode)
}

func (h *PromocodeHandler) GetAll(c *gin.Context) {
	promocodes, err := h.service.GetAllPromocodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, promocodes)
}

func (h *PromocodeHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный id"})
		return
	}

	var req models.PromocodeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный JSON"})
		return
	}

	promocode, err := h.service.UpdatePromocode(uint(id), req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, promocode)
}

func (h *PromocodeHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный id"})
		return
	}

	if err := h.service.DeletePromocode(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Успешно удалено"})
}
