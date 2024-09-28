package handler

import (
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	service "github.com/mBuergi86/devseconnect/internal/application"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/pkg/response"
)

type UserHandler struct {
	Service *service.UserService
	lock    sync.Mutex
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	users, err := h.Service.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewError("Failed to get user", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUsersByID(c echo.Context) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError("Invalid user ID format", http.StatusBadRequest, err))
	}

	user, err := h.Service.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.NewError("User not found", http.StatusNotFound, err))
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) FindUserByEmail(c echo.Context) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	email := c.Param("email")

	user, err := h.Service.FindByEmail(c.Request().Context(), email)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.NewError("User not found", http.StatusNotFound, err))
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	var user entity.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError("Invalid input", http.StatusBadRequest, err))
	}

	newUser, err := h.Service.Create(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewError("Failed to create user", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError("Invalid user ID format", http.StatusBadRequest, err))
	}

	var updateUser entity.User

	if err := c.Bind(&updateUser); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError("Invalid input", http.StatusBadRequest, err))
	}

	updateUser.UserID = id

	user, err := h.Service.Update(c.Request().Context(), updateUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewError("Failed to update user", http.StatusInternalServerError, err))
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError("Invalid user ID format", http.StatusBadRequest, err))
	}

	if err = h.Service.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewError("Failed to update user", http.StatusInternalServerError, err))
	}
	return c.NoContent(http.StatusNoContent)
}
