package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
)

type LikeHandler struct {
	likeService    *service.LikeService
	postService    *service.PostService
	commentService *service.CommentService
	userService    *service.UserService
}

func NewLikeHandler(
	likeService *service.LikeService,
	postService *service.PostService,
	commentService *service.CommentService,
	userService *service.UserService,
) *LikeHandler {
	return &LikeHandler{
		likeService:    likeService,
		postService:    postService,
		commentService: commentService,
		userService:    userService,
	}
}

func (h *LikeHandler) GetAllLikes(c echo.Context) error {
	likes, err := h.likeService.FindAllLikes(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch likes"})
	}

	return c.JSON(http.StatusOK, likes)
}

func (h *LikeHandler) GetLike(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid like ID"})
	}
	like, err := h.likeService.FindLikeByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Like not found"})
	}
	return c.JSON(http.StatusOK, like)
}

func (h *LikeHandler) CreateLike(c echo.Context) error {
	title := c.Param("title")
	username := c.Param("username")

	if title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid title"})
	}

	if username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid username"})
	}

	post, err := h.postService.GetPostByTitle(c.Request().Context(), title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get post by title"})
	}

	user, err := h.userService.GetUserByUsername(c.Request().Context(), username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user by username"})
	}

	var like entity.Likes
	if err := c.Bind(&like); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	like.PostID = post.PostID
	like.UserID = user.UserID

	if err := h.likeService.CreateByComment(
		c.Request().Context(),
		&like,
		title,
		username,
	); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create like"})
	}

	return c.JSON(http.StatusCreated, like)
}

func (h *LikeHandler) CreateByComment(c echo.Context) error {
	content := c.Param("content")
	username := c.Param("username")

	if content == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid content"})
	}

	if username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid username"})
	}

	comment, err := h.commentService.FindCommentByContent(c.Request().Context(), content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get comment by content"})
	}

	user, err := h.userService.GetUserByUsername(c.Request().Context(), username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user by username"})
	}

	var like entity.Likes
	if err := c.Bind(&like); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	like.CommentID = comment.CommentID
	like.UserID = user.UserID

	if err := h.likeService.CreateByComment(
		c.Request().Context(),
		&like,
		content,
		username,
	); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create like"})
	}

	return c.JSON(http.StatusCreated, like)
}

func (h *LikeHandler) DeleteLike(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid like ID"})
	}

	err = h.likeService.DeleteLike(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete like"})
	}

	return c.JSON(http.StatusOK, map[string]string{"like": "Like successfully deleted"})
}
