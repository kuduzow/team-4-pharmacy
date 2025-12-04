package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/config"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/repository"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
	"github.com/kuduzow/team-4-pharmacy/internal/transport"
)

func main() {

	logger := setupLogger()

	db := config.SetUpDatabaseConnection()

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
		logger.Error("не удалось мигрировать базу данных", slog.Any("error", err))
		os.Exit(1)
	}
	cartRepo := repository.NewCartRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	medicineRepo := repository.NewMedicineRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	promocodeRepo := repository.NewPromocodeRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	userRepo := repository.NewUserRepository(db)

	cartService := services.NewCartService(cartRepo)
	orderService := services.NewOrderService(orderRepo, paymentRepo)
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
		orderService,
		paymentService,
		promocodeService,
		reviewService,
		userService,
		cartService,
	)

	addr := getServerAddress()
	env := getEnvironment()
	logger.Info("сервер запустился",
		slog.String("addr", addr),
		slog.String("env", env),
	)
	if err := router.Run(addr); err != nil {
		logger.Error("не удалось запустить HTTP-сервер", slog.Any("error", err))
		os.Exit(1)
	}
}
func setupLogger() *slog.Logger {
	var level slog.Level
	logLevel := os.Getenv("LOG_LEVEL")

	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}),
	)
}

func getServerAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func getEnvironment() string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "local"
	}
	return env
}
