package routing

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/handler"
)

func SetupRouter(userService *service.UserService, postService *service.PostService, commentService *service.CommentService) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Create handlers
	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService, userService)
	commentHandler := handler.NewCommentHandler(commentService, postService, userService)

	// User routes
	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)
	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUser)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)

	e.POST("/posts/:username", postHandler.CreatePost)
	e.GET("/posts", postHandler.GetAllPosts)
	e.GET("/posts/:id", postHandler.GetPost)
	e.PUT("/posts/:id", postHandler.UpdatePost)
	e.DELETE("/posts/:id", postHandler.DeletePost)

	e.GET("/comments", commentHandler.GetAllComments)
	e.GET("/comments/:id", commentHandler.GetComment)
	e.POST("/comments/:title/:username", commentHandler.CreateComment)
	e.PUT("/comments/:id", commentHandler.UpdateComment)
	e.DELETE("/comments/:id", commentHandler.DeleteComment)

	return e
}
