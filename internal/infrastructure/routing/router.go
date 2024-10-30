package routing

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/handler"
	jwtMiddleware "github.com/mBuergi86/devseconnect/internal/infrastructure/middleware"
	"github.com/rs/zerolog"
)

type jwtCustomClaims struct {
	UserID   string          `json:"user_id"`
	Username string          `json:"username"`
	ExpireAt jwt.NumericDate `json:"exp"`
	jwt.RegisteredClaims
}

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
				Str("Time", v.StartTime.Format("2006-01-02 15:04:05")).
				Str("IP", c.RealIP()).AnErr("Error", v.Error).
				Str("Host", c.Request().Host).AnErr("Error", v.Error).
				Str("URI", v.URI).
				Int("Status", v.Status).
				Str("Method", c.Request().Method).
				Int64("ResponseTime", v.Latency.Milliseconds()).
				AnErr("Error", v.Error).
				Msg("Handled request")
			return nil
		},
	}))
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders:  []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		ExposeHeaders: []string{echo.HeaderAuthorization},
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            3600,
		HSTSExcludeSubdomains: true,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	// Create handlers
	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService, userService)
	commentHandler := handler.NewCommentHandler(commentService, postService, userService)
	tagHandler := handler.NewTagHandler(tagService)
	postTagHandler := handler.NewPostTagsHandler(postTagService)
	messageHandler := handler.NewMessageHandler(messageService, userService)
	likeHandler := handler.NewLikeHandler(likeService, postService, commentService, userService)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		logger.Error().Msg("JWT_SECRET is not set")
	}

	/* config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(secret),
		ErrorHandler: func(c echo.Context, err error) error {
			logger.Error().Err(err).Msgf("JWT Middleware error: %v", err)
			return c.JSON(401, map[string]string{"error": "Unauthorized"})
		},
	}*/

	jwtMiddleware := jwtMiddleware.JWTMiddleware(logger)

	// User routes login and register
	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)

	authenticated := e.Group("")
	authenticated.Use(jwtMiddleware)

	authenticated.GET("/users", userHandler.GetUsers)
	authenticated.GET("/user", userHandler.GetUser)
	authenticated.PUT("/user/update", userHandler.UpdateUser)
	authenticated.DELETE("/user/delete", userHandler.DeleteUser)

	authenticated.POST("/posts", postHandler.CreatePost)
	authenticated.GET("/posts", postHandler.GetAllPosts)
	authenticated.GET("/post", postHandler.GetPost)
	authenticated.PUT("/post", postHandler.UpdatePost)
	authenticated.DELETE("/post", postHandler.DeletePost)

	authenticated.GET("/comments", commentHandler.GetAllComments)
	authenticated.GET("/comment", commentHandler.GetComment)
	authenticated.POST("/comments/:title/:username", commentHandler.CreateComment)
	authenticated.PUT("/comment", commentHandler.UpdateComment)
	authenticated.DELETE("/comment", commentHandler.DeleteComment)

	authenticated.GET("/tags", tagHandler.GetTags)
	authenticated.GET("/tag", tagHandler.GetTag)
	authenticated.POST("/tags", tagHandler.CreateTag)
	authenticated.DELETE("/tag", tagHandler.DeleteTag)

	authenticated.GET("/posttags", postTagHandler.GetPostTags)
	authenticated.GET("/posttag", postTagHandler.GetPostTag)
	authenticated.POST("/posttags", postTagHandler.CreatePostTag)

	authenticated.GET("/messages", messageHandler.GetAllMessages)
	authenticated.GET("/message", messageHandler.GetMessage)
	authenticated.POST("/messages", messageHandler.CreateMessage)
	authenticated.PUT("/message", messageHandler.UpdateMessage)
	authenticated.DELETE("/message", messageHandler.DeleteMessage)

	authenticated.GET("/likes", likeHandler.GetAllLikes)
	authenticated.GET("/like", likeHandler.GetLike)
	authenticated.POST("/likes/:title/:username", likeHandler.CreateLike)

	return e
}
