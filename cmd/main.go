package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
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
	logger.Info().Msg("Starting server...")
	fmt.Println("Starting server...")
	if _, err := os.Stat(".env"); err == nil { // Datei existiert
		if err := godotenv.Load(); err != nil {
			fmt.Println("Error loading .env file")
			logger.Fatal().Err(err).Msg("Error loading .env file")
		}
	} else if os.IsNotExist(err) {
		fmt.Println("No .env file found, skipping load")
		logger.Info().Msg("No .env file found, skipping load")
	} else {
		fmt.Println("Error checking .env file existence")
		logger.Fatal().Err(err).Msg("Error checking .env file existence")
	}

	// Initialize PostgreSQL
	db, err := database.InitPostgres()
	if err != nil {
		fmt.Println("Failed to connect to PostgreSQL")
		logger.Fatal().Err(err).Msg("Failed to connect to PostgreSQL")
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to get SQL DB")
		logger.Fatal().Err(err).Msg("Failed to get SQL DB")
	}
	defer sqlDB.Close()

	// Initialize Redis
	redisClient, err := cache.InitRedis()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		logger.Fatal().Err(err).Msg("Failed to connect to Redis")
	}
	defer redisClient.Close()

	// Initialize RabbitMQ
	rabbitmqConn, err := messaging.InitRabbitMQ()
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
		logger.Fatal().Err(err).Msg("Failed to connect to RabbitMQ")
	}
	defer rabbitmqConn.Close()

	rabbitMQChan, err := rabbitmqConn.Channel()
	if err != nil {
		fmt.Println("Failed to create RabbitMQ channel")
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
		fmt.Println("Failed to declare an exchange")
		logger.Fatal().Err(err).Msgf("Failed to declare an exchange: %v", err)
	}

	producer, err := messaging.NewProducer(rabbitmqConn)
	if err != nil {
		fmt.Println("Failed to create producer")
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
		fmt.Println("Failed to create user service")
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
		fmt.Println("Failed to create user consumer")
		logger.Fatal().Err(err).Msgf("Failed to create user consumer: %v", err)
	}

	postConsumer, err := messaging.NewConsumer(rabbitmqConn, "post_queue")
	if err != nil {
		fmt.Println("Failed to create post consumer")
		logger.Fatal().Err(err).Msgf("Failed to create post consumer: %v", err)
	}

	commentConsumer, err := messaging.NewConsumer(rabbitmqConn, "comment_queue")
	if err != nil {
		fmt.Println("Failed to create comment consumer")
		logger.Fatal().Err(err).Msgf("Failed to create comment consumer: %v", err)
	}

	tagConsumer, err := messaging.NewConsumer(rabbitmqConn, "tag_queue")
	if err != nil {
		fmt.Println("Failed to create tag consumer")
		logger.Fatal().Err(err).Msgf("Failed to create tag consumer: %v", err)
	}

	postTagConsumer, err := messaging.NewConsumer(rabbitmqConn, "post_tag_queue")
	if err != nil {
		fmt.Println("Failed to create post tag consumer")
		logger.Fatal().Err(err).Msgf("Failed to create post tag consumer: %v", err)
	}

	messageConsumer, err := messaging.NewConsumer(rabbitmqConn, "message_queue")
	if err != nil {
		fmt.Println("Failed to create message consumer")
		logger.Fatal().Err(err).Msgf("Failed to create message consumer: %v", err)
	}

	likeConsumer, err := messaging.NewConsumer(rabbitmqConn, "like_queue")
	if err != nil {
		fmt.Println("Failed to create like consumer")
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
			fmt.Println("Failed to start user consumer")
			logger.Fatal().Err(err).Msgf("Failed to start user consumer: %v", err)
		}
	}()

	go func() {
		if err := pc.Start(); err != nil {
			fmt.Println("Failed to start post consumer")
			logger.Fatal().Err(err).Msgf("Failed to start post consumer: %v", err)
		}
	}()

	go func() {
		if err := cc.Start(); err != nil {
			fmt.Println("Failed to start comment consumer")
			logger.Fatal().Err(err).Msgf("Failed to start comment consumer: %v", err)
		}
	}()

	go func() {
		if err := tc.Start(); err != nil {
			fmt.Println("Failed to start tag consumer")
			logger.Fatal().Err(err).Msgf("Failed to start tag consumer: %v", err)
		}
	}()

	go func() {
		if err := ptc.Start(); err != nil {
			fmt.Println("Failed to start post tag consumer")
			logger.Fatal().Err(err).Msgf("Failed to start post tag consumer: %v", err)
		}
	}()

	go func() {
		if err := mc.Start(); err != nil {
			fmt.Println("Failed to start message consumer")
			logger.Fatal().Err(err).Msgf("Failed to start message consumer: %v", err)
		}
	}()

	go func() {
		if err := lc.Start(); err != nil {
			fmt.Println("Failed to start like consumer")
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
		fmt.Println("Server starting on port " + port)
		logger.Info().Msgf("Server starting on port %s", port)
		if err := router.Start("0.0.0.0:" + port); err != nil {
			fmt.Println("Failed to start server")
			logger.Fatal().Err(err).Msgf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")
	logger.Info().Msg("Shutting down server...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		fmt.Println("Failed to shutdown server")
		logger.Fatal().Err(err).Msgf("Failed to shutdown server: %v", err)
	}

	fmt.Println("Server exited properly")
	logger.Info().Msg("Server exited properly")
}
