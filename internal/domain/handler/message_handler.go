package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
)

type MessageHandler struct {
	messageService *service.MessageService
	userService    *service.UserService
}

func NewMessageHandler(
	messageService *service.MessageService,
	userService *service.UserService,
) *MessageHandler {
	return &MessageHandler{
		messageService: messageService,
		userService:    userService,
	}
}

func (h *MessageHandler) GetAllMessages(c echo.Context) error {
	messages, err := h.messageService.FindAllMessages(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch messages"})
	}

	return c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) GetMessage(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid message ID"})
	}
	message, err := h.messageService.FindMessageByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Message not found"})
	}
	return c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) CreateMessage(c echo.Context) error {
	username1 := c.Param("username1")
	username2 := c.Param("username2")

	if username1 == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid username1"})
	}

	if username2 == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid username2"})
	}

	user1, err := h.userService.GetUserByUsername(c.Request().Context(), username1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get sender by username1"})
	}

	user2, err := h.userService.GetUserByUsername(c.Request().Context(), username2)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get receiver by username2"})
	}

	var message entity.Messages
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	message.SenderID = user1.UserID
	message.ReceiverID = user2.UserID

	if err := h.messageService.CreateMessage(
		c.Request().Context(),
		&message,
		username1,
		username2,
	); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create message"})
	}

	message.User1 = user1.Username
	message.User2 = user2.Username

	return c.JSON(http.StatusCreated, message)
}

func (h *MessageHandler) UpdateMessage(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid message ID"})
	}

	var updateData entity.Messages

	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updateData.MessageID = id

	updatedMessage, err := h.messageService.UpdateMessage(
		c.Request().Context(),
		&updateData,
		id,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update message"})
	}

	return c.JSON(http.StatusOK, updatedMessage)
}

func (h *MessageHandler) DeleteMessage(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid message ID"})
	}

	err = h.messageService.DeleteMessage(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete message"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Message successfully deleted"})
}
