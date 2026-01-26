package transport

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-4-pharmacy/internal/models"
	"github.com/kuduzow/team-4-pharmacy/internal/services"
)

type UserHandler struct {
	service services.UserService
	logger  *slog.Logger
}

func NewUserHandler(logger *slog.Logger, service services.UserService) *UserHandler {
	return &UserHandler{service: service, logger: logger}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/:id", h.Get)
		users.GET("/", h.GetAll)
		users.POST("/", h.Create)
		users.PATCH("/:id", h.Update)
		users.DELETE("/:id", h.Delete)
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	h.logger.Info("Обработка запроса на создание пользователя")
	var req models.UserCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Ошибка привязки JSON при создании пользователя", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.service.CreateUser(req)
	if err != nil {
		h.logger.Error("Ошибка при создании пользователя", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info("Пользователь успешно создан", "user", user)
	c.JSON(http.StatusCreated, user)

}

func (h *UserHandler) Get(c *gin.Context) {

	h.logger.Info("Обработка запроса на получение пользователя")

	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Warn("Некорректный идентификатор пользователя", "id", idStr, "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный идентификатор"})
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			h.logger.Warn("Пользователь не найден", "id", id)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.logger.Error("Ошибка при получении пользователя", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info("Пользователь успешно получен", "user", user)
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Update(c *gin.Context) {

	h.logger.Info("Обработка запроса на обновление пользователя")

	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Warn("Некорректный идентификатор пользователя", "id", idStr, "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный идентификатор"})
		return
	}

	var req models.UserUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Ошибка привязки JSON при обновлении пользователя", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.service.UpdateUser(uint(id), req)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			h.logger.Error("Пользователь не найден для обновления", "id", id)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		h.logger.Error("Ошибка при обновлении пользователя", "id", id, "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info("Пользователь успешно обновлен", "user", user)
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	h.logger.Info("Обработка запроса на удаление пользователя")
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Warn("Некорректный идентификатор пользователя", "id", idStr, "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный идентификатор"})
		return
	}
	if err := h.service.DeleteUser(uint(id)); err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			h.logger.Error("Пользователь не найден для удаления", "id", id)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		h.logger.Error("Ошибка при удалении пользователя", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info("Пользователь успешно удален", "id", id)
	c.Status(http.StatusNoContent)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	h.logger.Info("Обработка запроса на получение всех пользователей")
	users, err := h.service.GetAllUsers()
	if err != nil {
		h.logger.Error("Ошибка при получении всех пользователей", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info("Все пользователи успешно получены", "count", len(users))
	c.JSON(http.StatusOK, users)
}
