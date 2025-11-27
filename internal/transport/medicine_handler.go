package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
	"gorm.io/gorm"
)

type MedicineHandler struct {
	service services.MedicineService
}

func NewMedicineHandler(service services.MedicineService) *MedicineHandler {
	return &MedicineHandler{service: service}
}

func (h *MedicineHandler) RegisterRoutes(r *gin.Engine) {
	medicines := r.Group("/medicines")
	{
		medicines.POST("", h.Create)
		medicines.GET("/:id", h.Get)
		medicines.DELETE("/:id", h.Delete)
		medicines.PATCH("/:id", h.Update)
		medicines.GET("", h.List)
	}
}

func (h *MedicineHandler) Create(c *gin.Context) {
	var req models.MedicineCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректный JSON",
		})
		return
	}

	medicine, err := h.service.CreateMedicine(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, medicine)
}

func (h *MedicineHandler) Get(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id"})
		return
	}

	medicine, err := h.service.GetMedicineByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHandler) Update(c *gin.Context) {

	idstr := c.Param("id")

	id, err := strconv.ParseUint(idstr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id"})
		return
	}

	var req models.MedicineUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
		return
	}

	medicine, err := h.service.UpdateMedicine(uint(id), req)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id"})
		return
	}

	if err := h.service.DeleteMedicine(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "удалено успешшно"})
}

func (h *MedicineHandler) List(c *gin.Context) {
	var filter repository.MedicineFilter

	if categoryIDStr := c.Query("category_id"); categoryIDStr != "" {
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err == nil {
			categoryIDUint := uint(categoryID)
			filter.CategoryID = &categoryIDUint
		}
	}

	if subcategoryIDStr := c.Query("subcategory_id"); subcategoryIDStr != "" {
		subcategoryID, err := strconv.ParseUint(subcategoryIDStr, 10, 64)
		if err == nil {
			subcategoryIDUint := uint(subcategoryID)
			filter.SubcategoryID = &subcategoryIDUint
		}
	}

	if inStockStr := c.Query("in_stock"); inStockStr != "" {
		inStock, err := strconv.ParseBool(inStockStr)
		if err == nil {
			filter.InStock = &inStock
		}
	}

	medicines, err := h.service.ListMedicines(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicines)
}
