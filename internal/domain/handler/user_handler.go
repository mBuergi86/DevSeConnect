package handler

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/rs/zerolog"
)

type UserHandler struct {
	userService *service.UserService
}

type jwtCustomClaims struct {
	UserID   string          `json:"user_id"`
	Username string          `json:"username"`
	ExpireAt jwt.NumericDate `json:"exp"`
	jwt.RegisteredClaims
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

var logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func (h *UserHandler) Register(c echo.Context) error {
	var input struct {
		Username  string `json:"username" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=6"`
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		logger.Error().Err(err).Msg("Failed to bind input")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user := &entity.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: input.Password,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
	}

	if err := h.userService.Register(c.Request().Context(), user); err != nil {
		logger.Error().Err(err).Msg("Failed to register user")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c echo.Context) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user, token, err := h.userService.Login(c.Request().Context(), input.Username, input.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   int(time.Hour.Seconds() * 24),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.userService.GetUsers(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Get("user_id").(uuid.UUID)

	user, err := h.userService.GetUserByID(c.Request().Context(), id)
	if err != nil {
		logger.Error().Msgf("User not found: %v", err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Get("user_id").(uuid.UUID)

	updateData := make(map[string]interface{})

	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updateData["user_id"] = id

	updatedUser, err := h.userService.UpdateUser(c.Request().Context(), updateData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Get("user_id").(uuid.UUID)

	err := h.userService.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User successfully deleted"})
}
