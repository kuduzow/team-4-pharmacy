package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
)

type CartItemHandler struct {
	service services.CartItemService
}

func NewCartItemHandler(service services.CartItemService) *CartItemHandler {
	return &CartItemHandler{service: service}
}

func (h *CartItemHandler) RegisterRoutes(r *gin.Engine) {
	items := r.Group("/users/:id/cart/items")
	{
		items.POST("", h.AddItem)
	}
}

func (h *CartItemHandler) AddItem(c *gin.Context) {
	userIDStr := c.Param("id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id обязателен"})
		return
	}

	id, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "недействительный user_id"})
		return
	}

	var req models.CartCreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "недействительный JSON"})
		return
	}

	updated, err := h.service.AddItem(uint(id), req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidQuantity) || errors.Is(err, services.ErrOutOfStock) || errors.Is(err, services.ErrMedicineMissing) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, updated)
}
