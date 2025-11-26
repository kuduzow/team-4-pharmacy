package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
)

func RegisterRoutes(
	router *gin.Engine,
	categoryService services.CategoryService,
	medicineService services.MedicineService,
	orderService services.OrderService,
	paymentService services.PaymentService,
	promocodeService services.PromocodeService,
	reviewService services.ModelService,
	userService services.UserService,
	cartService services.CartService,
) {
	categoryHandler := NewCategoryHandler(categoryService)
	medicineHandler := NewMedicineHandler(medicineService)
	orderHandler := NewOrderHandler(orderService)
	paymentHandler := NewPaymentHandler(paymentService)
	promocodeHandler := NewPromocodeHandler(promocodeService)
	reviewHandler := NewReviewHandler(reviewService)
	userHandler := NewUserHandler(userService)
	cartHandler := NewCartHandler(cartService)

	categoryHandler.RegisterRoutes(router)
	medicineHandler.RegisterRoutes(router)
	orderHandler.RegisterRoute(router)
	paymentHandler.RegisterRoutes(router)
	promocodeHandler.RegisterRoutes(router)
	reviewHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	cartHandler.RegisterRoutes(router)

}
