package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/config"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
	"github.com/kuduzow/team-4-pharmacy/internal/transport"
)

func initLogger() *slog.Logger {
	level := slog.LevelInfo

	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	case "info":
		level = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return slog.New(handler)
}

func main() {
	db := config.SetUpDatabaseConnection()

	logger := initLogger()

	if err := db.AutoMigrate(
		&models.Cart{},
		&models.CartItem{},
		&models.Category{},
		&models.Medicine{},
		&models.Order{},
		&models.Payment{},
		&models.Promocode{},
		&models.Review{},
		&models.User{},
	); err != nil {
		logger.Error("Ошибка при миграции базы данных", slog.Any("error", err))
	}
	cartRepo := repository.NewCartRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	medicineRepo := repository.NewMedicineRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	promocodeRepo := repository.NewPromocodeRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	userRepo := repository.NewUserRepository(logger, db)

	cartService := services.NewCartService(cartRepo)
	orderService := services.NewOrderService(orderRepo, paymentRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	medicineService := services.NewMedicineService(medicineRepo, categoryRepo)

	paymentService := services.NewPaymentService(paymentRepo)
	promocodeService := services.NewPromocodeService(promocodeRepo)
	reviewService := services.NewReviewService(reviewRepo)
	userService := services.NewUserService(logger, userRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	router := gin.Default()

	transport.RegisterRoutes(
		router,
		categoryService,
		medicineService,
		orderService,
		paymentService,
		promocodeService,
		reviewService,
		userService,
		logger,
		cartService,
	)

	logger.Info("server started",
		slog.String("addr=", ":"+port),
		slog.String("env=", env))

	if err := router.Run(":" + port); err != nil {
		logger.Error("не удалось запустить HTTP-сервер", slog.Any("error", err))
	}
}
