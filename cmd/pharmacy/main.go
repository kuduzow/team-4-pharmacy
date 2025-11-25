package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/config"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
	"github.com/kuduzow/team-4-pharmacy/internal/transport"
)

func main() {
	db := config.SetUpDatabaseConnection()

	if err := db.AutoMigrate(
		&models.Cart{},
		&models.Category{},
		&models.Medicine{},
		&models.Order{},
		&models.Payment{},
		&models.Promocode{},
		&models.Review{},
		&models.User{},
	); err != nil {
		log.Fatalf("не удалось мигрировать: %v ", err)
	}
	// cartRepo := repository.NewCartRepository(db) - не доделан - усман доделает
	categoryRepo := repository.NewCategoryRepository(db)
	medicineRepo := repository.NewMedicineRepository(db)
	// orderRepo := repository.NewOrderRepository(db) - хамзат доделает
	paymentRepo := repository.NewPaymentRepository(db)
	promocodeRepo := repository.NewPromocodeRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	userRepo := repository.NewUserRepository(db)

	// cartService := services.NewCartService - не доделан - усман доделает
 	// orderService := services.NewOrderService(orderRepo, paymentRepo) - хамзат доделает
	categoryService := services.NewCategoryService(categoryRepo)
	medicineService := services.NewMedicineService(medicineRepo, categoryRepo)

	paymentService := services.NewPaymentService(paymentRepo)
	promocodeService := services.NewPromocodeService(promocodeRepo)
	reviewService := services.NewReviewService(reviewRepo)
	userService := services.NewUserService(userRepo)

	router := gin.Default()

	transport.RegisterRoutes(
		router,
		categoryService,
		medicineService,
		// cartService, усман доделает
		// orderService, хамзат доделает
		paymentService,
		promocodeService,
		reviewService,
		userService,
	)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
