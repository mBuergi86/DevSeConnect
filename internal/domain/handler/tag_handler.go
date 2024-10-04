package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
)

type TagHandler struct {
	TagService *service.TagsService
}

func NewTagHandler(tagService *service.TagsService) *TagHandler {
	if tagService == nil {
		panic("tagService is required")
	}
	return &TagHandler{
		TagService: tagService,
	}
}

func (h *TagHandler) GetTags(c echo.Context) error {
	tags, err := h.TagService.GetTags(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tags)
}

func (h *TagHandler) GetTag(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tag, err := h.TagService.GetTag(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tag)
}

func (h *TagHandler) CreateTag(c echo.Context) error {
	tag := new(entity.Tags)
	if err := c.Bind(tag); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.TagService.CreateTag(c.Request().Context(), tag); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, tag)
}

func (h *TagHandler) DeleteTag(c echo.Context) error {
	tagID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.TagService.DeleteTag(c.Request().Context(), tagID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
