package routing

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/handler"
	"github.com/rs/zerolog"
)

func SetupRouter(
	userService *service.UserService,
	postService *service.PostService,
	commentService *service.CommentService,
	tagService *service.TagsService,
	postTagService *service.PostTagsService,
	messageService *service.MessageService,
	likeService *service.LikeService,
) *echo.Echo {
	e := echo.New()

	// Middleware
	logger := zerolog.New(os.Stdout)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("Status", v.Status).
				Msg("handled request")
			return nil
		},
	}))
	e.Use(middleware.Recover())

	// Create handlers
	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService, userService)
	commentHandler := handler.NewCommentHandler(commentService, postService, userService)
	tagHandler := handler.NewTagHandler(tagService)
	postTagHandler := handler.NewPostTagsHandler(postTagService)
	messageHandler := handler.NewMessageHandler(messageService, userService)
	likeHandler := handler.NewLikeHandler(likeService, postService, commentService, userService)

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

	e.GET("/tags", tagHandler.GetTags)
	e.GET("/tags/:id", tagHandler.GetTag)
	e.POST("/tags", tagHandler.CreateTag)
	e.DELETE("/tags/:id", tagHandler.DeleteTag)

	e.GET("/posttags", postTagHandler.GetPostTags)
	e.GET("/posttags/:id", postTagHandler.GetPostTag)
	e.POST("/posttags", postTagHandler.CreatePostTag)

	e.GET("/messages", messageHandler.GetAllMessages)
	e.GET("/messages/:id", messageHandler.GetMessage)
	e.POST("/messages", messageHandler.CreateMessage)
	e.PUT("/messages/:id", messageHandler.UpdateMessage)
	e.DELETE("/messages/:id", messageHandler.DeleteMessage)

	e.GET("/likes", likeHandler.GetAllLikes)
	e.GET("/likes/:id", likeHandler.GetLike)
	e.POST("/likes/:title/:username", likeHandler.CreateLike)

	return e
}
