package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
)

type PostTagsHandler struct {
	postTagsService *service.PostTagsService
}

func NewPostTagsHandler(postTagsService *service.PostTagsService) *PostTagsHandler {
	return &PostTagsHandler{
		postTagsService: postTagsService,
	}
}

func (h *PostTagsHandler) GetPostTags(c echo.Context) error {
	postTags, err := h.postTagsService.GetPostTags(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postTags)
}

func (h *PostTagsHandler) GetPostTag(c echo.Context) error {
	postTagID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	postTag, err := h.postTagsService.GetTag(c.Request().Context(), postTagID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postTag)
}

func (h *PostTagsHandler) CreatePostTag(c echo.Context) error {
	postTag := new(entity.PostTags)
	title := c.Param("title")
	tags := c.Param("tags")

	if err := c.Bind(postTag); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := h.postTagsService.CreateTag(c.Request().Context(), postTag, title, tags)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, postTag)
}
