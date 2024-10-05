package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
)

type CommentHandler struct {
	commentService *service.CommentService
	postService    *service.PostService
	userService    *service.UserService
}

func NewCommentHandler(
	commentService *service.CommentService,
	postService *service.PostService,
	userService *service.UserService,
) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
		postService:    postService,
		userService:    userService,
	}
}

func (h *CommentHandler) GetAllComments(c echo.Context) error {
	comments, err := h.commentService.FindAllComments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch posts"})
	}

	return c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) GetComment(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}
	comment, err := h.commentService.FindCommentByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Comment not found"})
	}
	return c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) CreateComment(c echo.Context) error {
	title := c.Param("title")

	if title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid title"})
	}

	post, err := h.postService.GetPostByTitle(c.Request().Context(), title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get post by title"})
	}

	username := c.Param("username")

	if username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid username"})
	}

	user, err := h.userService.GetUserByUsername(c.Request().Context(), username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}

	var comment entity.Comments
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	comment.PostID = post.PostID
	comment.UserID = user.UserID

	if err := h.commentService.CreateComment(
		c.Request().Context(),
		&comment,
		title,
		username,
	); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create post"})
	}

	comment.Post = post
	comment.User = user

	return c.JSON(http.StatusCreated, comment)
}

func (h *CommentHandler) UpdateComment(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	var updateData entity.Comments

	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updateData.CommentID = id

	updatedComment, err := h.commentService.UpdateComment(
		c.Request().Context(),
		&updateData,
		id,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update post"})
	}

	return c.JSON(http.StatusOK, updatedComment)
}

func (h *CommentHandler) DeleteComment(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid post ID"})
	}

	err = h.commentService.DeleteComment(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete post"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Comment successfully deleted"})
}
