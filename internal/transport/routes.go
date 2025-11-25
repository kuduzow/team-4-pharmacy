package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
)

func RegisterRoutes(
	router *gin.Engine,
	categoryService services.CategoryService,
	medicineService services.MedicineService,
	// orderService services.OrderService,   - хамзат доделает
	paymentService services.PaymentService,
	promocodeService services.PromocodeService,
	reviewService services.ReviewService,
	userService services.UserService,
	// нужно добавить КАРТ
) {
	categoryHandler := NewCategoryHandler(categoryService)
	medicineHandler := NewMedicineHandler(medicineService)
	// orderHandler := NewOrderHandler(orderService)  - хамзат доделает
	paymentHandler := NewPaymentHandler(paymentService)
	promocodeHandler := NewPromocodeHandler(promocodeService)
	reviewHandler := NewReviewHandler(&reviewService)
	UserHandler := NewUserHandler(userService)
		// нужно добавить КАРТ


	categoryHandler.RegisterRoutes(router)
	medicineHandler.RegisterRoutes(router)
	// orderHandler.RegisterRoutes(router)   - хамзат доделает
	paymentHandler.RegisterRoute(router)
	promocodeHandler.RegisterRoutes(router)
	reviewHandler.RegisterRoutes(router)
	UserHandler.RegisterRoute(router)
		// нужно добавить КАРТ


}
