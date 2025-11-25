package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
)

type PaymentHandler struct {
	service services.PaymentService
}

func NewPaymentHandler(service services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) RegisterRoute(r *gin.Engine) {
	payments := r.Group("/payment")
	{

		payments.POST("", h.Create)
		payments.GET("",h.Get)
		payments.PATCH("",h.Update)
		payments.DELETE("",h.Delete)
	}
}
func (h *PaymentHandler) Create(c *gin.Context) {
	var req models.PaymentCreate

	if err := c.ShouldBindJSON(&req).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	payment, err := h.service.CreatePayment(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err})
		return
	}
	c.JSON(http.StatusCreated, payment)
}
func (h *PaymentHandler) Get(c *gin.Context) {
	idstr := c.Param("id")

	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}
	payment, err := h.service.GetPaymentByID(uint(id))
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"error"})
		return
	}
	var req models.PaymentUpdate

	payment ,err := h.service.UpdatePayment(uint(id),req)
	if err != nil{
		if errors.Is(err, services.ErrUserNotFound){
			c.JSON(http.StatusNotFound, gin.H{"ERROR":"ERR"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"erroe":err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func(h *PaymentHandler)Delete(c *gin.Context){
	idStr := c.Param("id")

	payment ,err := strconv.ParseUint(idStr, 10,64)
	if err != nil{
		if errors.Is(err, services.ErrUserNotFound){
		c.JSON(http.StatusNotFound,gin.H{"error":"err"})
		return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,payment)
}