package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
)

type ReviewHandler struct {
	service services.ModelService
}

func (h *ReviewHandler) RegisterRoutes(r *gin.Engine) {
	reviews := r.Group("/reviews")
	{
		reviews.GET("/:id", h.GetByID)
		reviews.PATCH("/:id", h.UpdateReview)
		reviews.DELETE("/:id", h.DeleteReview)
	}

	medicines := r.Group("/medicines/:id")
	{
		medicines.POST("/reviews", h.CreateReview)
		medicines.GET("/reviews", h.ListByMedicineID)
		medicines.GET("/avg_rating", h.GetAvgRating)
	}
}

func NewReviewHandler(service services.ModelService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req models.ReviewForPost

	medicineID := c.Param("id")

	convert, err := strconv.ParseUint(medicineID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.MedicineID = uint(convert)

	review, err := h.service.CreateReview(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)

}

func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	id := c.Param("id")

	convert, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.DeleteReview(uint(convert))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})

}

func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	var input models.ReviewForUpdate
	id := c.Param("id")

	convert, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.UpdateReview(uint(convert), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)

}

func (h *ReviewHandler) GetByID(c *gin.Context) {

	id := c.Param("id")

	convert, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.GetByID(uint(convert))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)

}

func (h *ReviewHandler) ListByMedicineID(c *gin.Context) {
	id := c.Param("id")

	medicineID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	reviews, err := h.service.ListByMedicineID(uint(medicineID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) GetAvgRating(c *gin.Context) {
	id := c.Param("id")

	medicineID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	avg, err := h.service.GetAvgRating(uint(medicineID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, avg)
}
