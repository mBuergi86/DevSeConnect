package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/cache"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/database"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/messaging"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/routing"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func main() {
	//if err := godotenv.Load(); err != nil {
	//	logger.Fatal().Err(err).Msg("Error loading .env file")
	//}

	// Initialize PostgreSQL
	db, err := database.InitPostgres()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to PostgreSQL")
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get SQL DB")
	}
	defer sqlDB.Close()

	// Initialize Redis
	redisClient, err := cache.InitRedis()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to Redis")
	}
	defer redisClient.Close()

	// Initialize RabbitMQ
	rabbitmqConn, err := messaging.InitRabbitMQ()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to RabbitMQ")
	}
	defer rabbitmqConn.Close()

	rabbitMQChan, err := rabbitmqConn.Channel()
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create RabbitMQ channel: %v", err)
	}
	defer rabbitMQChan.Close()

	// Declare exchange
	err = rabbitMQChan.ExchangeDeclare(
		"user_events", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to declare an exchange: %v", err)
	}

	producer, err := messaging.NewProducer(rabbitmqConn)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create producer: %v", err)
	}

	// Setup repositories
	userRepo := repository.NewUserRepository(db, redisClient)
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	tagRepo := repository.NewTagsRepository(db)
	postTagRepo := repository.NewPostTagsRepository(db)
	messageRepo := repository.NewMessageRepository(db)
	likeRepo := repository.NewLikeRepository(db)

	// Setup services
	userService, err := service.NewUserService(userRepo, rabbitMQChan, producer)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create user service: %v", err)
	}

	postService := service.NewPostService(postRepo, userRepo, rabbitMQChan)
	commentService := service.NewCommentService(commentRepo, postRepo, userRepo, rabbitMQChan)
	tagService := service.NewTagsService(tagRepo, rabbitMQChan)
	postTagService := service.NewPostTagsService(postTagRepo, rabbitMQChan)
	messageService := service.NewMessageService(messageRepo, userRepo, rabbitMQChan)
	likeService := service.NewLikeService(likeRepo, postRepo, commentRepo, userRepo, rabbitMQChan)

	// Setup consumer
	userConsumer, err := messaging.NewConsumer(rabbitmqConn, "user_queue")
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create user consumer: %v", err)
	}

	postConsumer, err := messaging.NewConsumer(rabbitmqConn, "post_queue")
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create post consumer: %v", err)
	}

	commentConsumer, err := messaging.NewConsumer(rabbitmqConn, "comment_queue")
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create comment consumer: %v", err)
	}

	tagConsumer, err := messaging.NewConsumer(rabbitmqConn, "tag_queue")
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create tag consumer: %v", err)
	}

	postTagConsumer, err := messaging.NewConsumer(rabbitmqConn, "post_tag_queue")
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create post tag consumer: %v", err)
	}

	messageConsumer, err := messaging.NewConsumer(rabbitmqConn, "message_queue")
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create message consumer: %v", err)
	}

	likeConsumer, err := messaging.NewConsumer(rabbitmqConn, "like_queue")
	if err != nil {
		logger.Fatal().Err(err).Msgf("Failed to create like consumer: %v", err)
	}

	// Start consumer
	uc := messaging.NewUserConsumer(userConsumer, userRepo)
	pc := messaging.NewPostConsumer(postConsumer, postRepo)
	cc := messaging.NewCommentConsumer(commentConsumer, commentRepo)
	tc := messaging.NewTagsConsumer(tagConsumer, tagRepo)
	ptc := messaging.NewPostTagsConsumer(postTagConsumer, postTagRepo)
	mc := messaging.NewMessageConsumer(messageConsumer, messageRepo)
	lc := messaging.NewLikeConsumer(likeConsumer, likeRepo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := uc.Start(ctx); err != nil {
			logger.Fatal().Err(err).Msgf("Failed to start user consumer: %v", err)
		}
	}()

	go func() {
		if err := pc.Start(); err != nil {
			logger.Fatal().Err(err).Msgf("Failed to start post consumer: %v", err)
		}
	}()

	go func() {
		if err := cc.Start(); err != nil {
			logger.Fatal().Err(err).Msgf("Failed to start comment consumer: %v", err)
		}
	}()

	go func() {
		if err := tc.Start(); err != nil {
			logger.Fatal().Err(err).Msgf("Failed to start tag consumer: %v", err)
		}
	}()

	go func() {
		if err := ptc.Start(); err != nil {
			logger.Fatal().Err(err).Msgf("Failed to start post tag consumer: %v", err)
		}
	}()

	go func() {
		if err := mc.Start(); err != nil {
			logger.Fatal().Err(err).Msgf("Failed to start message consumer: %v", err)
		}
	}()

	go func() {
		if err := lc.Start(); err != nil {
			logger.Fatal().Err(err).Msgf("Failed to start like consumer: %v", err)
		}
	}()

	// Setup router
	router := routing.SetupRouter(userService, postService, commentService, tagService, postTagService, messageService, likeService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	// Start server in a goroutine
	go func() {
		logger.Info().Msgf("Server starting on port %s", port)
		if err := router.Start("0.0.0.0:" + port); err != nil {
			logger.Fatal().Err(err).Msgf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info().Msg("Shutting down server...")

	if err := router.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msgf("Failed to shutdown server: %v", err)
	}

	logger.Info().Msg("Server exited properly")
}
