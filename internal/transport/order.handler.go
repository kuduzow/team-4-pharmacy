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

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) RegisterRoute(r *gin.Engine) {
	orders := r.Group("/order")
	{
		orders.GET("/:id", h.Get)
		orders.PATCH("/:id/status",h.Update)
	}
	users := r.Group("/users")
	{
		users.POST("/:id/orders")
		users.GET("/:id/orders")
		users.PATCH("/:id/status")
	}
}
func (h *OrderHandler) Get(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}
	order, err := h.service.GetOrderByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 61)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	var req models.OrderUpdate
	if err := c.ShouldBindJSON(&req);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	order , err := h.service.UpdateOrder(uint(id),req)
	if err != nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,order)
}
