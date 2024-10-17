package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
)

type PostHandler struct {
	postService *service.PostService
	userService *service.UserService
}

func NewPostHandler(postService *service.PostService, userService *service.UserService) *PostHandler {
	return &PostHandler{
		postService: postService,
		userService: userService,
	}
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	userID := c.Get("user_id").(uuid.UUID)

	var post entity.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	post.UserID = userID

	if err := h.postService.CreatePost(c.Request().Context(), &post); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create post"})
	}

	return c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetAllPosts(c echo.Context) error {
	posts, err := h.postService.GetAllPosts(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch posts"})
	}

	return c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetPost(c echo.Context) error {
	id, ok := c.Get("user_id").(uuid.UUID)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	post, err := h.postService.GetPostByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
	}
	return c.JSON(http.StatusOK, post)
}

func (h *PostHandler) UpdatePost(c echo.Context) error {
	id, ok := c.Get("user_id").(uuid.UUID)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	var updateData entity.Post

	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updateData.UserID = id

	updatedPost, err := h.postService.UpdatePost(c.Request().Context(), &updateData, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update post"})
	}

	return c.JSON(http.StatusOK, updatedPost)
}

func (h *PostHandler) DeletePost(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	err = h.postService.DeletePost(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete post"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Post successfully deleted"})
}
