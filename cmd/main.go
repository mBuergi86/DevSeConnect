package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/mBuergi86/devseconnect/internal/application/service"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/cache"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/database"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/messaging"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/routing"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize PostgreSQL
	db, err := database.InitPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}
	defer sqlDB.Close()

	// Initialize Redis
	redisClient, err := cache.InitRedis()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	// Initialize RabbitMQ
	rabbitmqConn, err := messaging.InitRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitmqConn.Close()

	rabbitMQChan, err := rabbitmqConn.Channel()
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ channel: %v", err)
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
		log.Fatalf("Failed to declare an exchange: %v", err)
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
	userService, err := service.NewUserService(userRepo, rabbitMQChan)
	if err != nil {
		log.Fatalf("Failed to create user service: %v", err)
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
		log.Fatalf("Failed to create user consumer: %v", err)
	}

	postConsumer, err := messaging.NewConsumer(rabbitmqConn, "post_queue")
	if err != nil {
		log.Fatalf("Failed to create post consumer: %v", err)
	}

	commentConsumer, err := messaging.NewConsumer(rabbitmqConn, "comment_queue")
	if err != nil {
		log.Fatalf("Failed to create comment consumer: %v", err)
	}

	tagConsumer, err := messaging.NewConsumer(rabbitmqConn, "tag_queue")
	if err != nil {
		log.Fatalf("Failed to create tag consumer: %v", err)
	}

	postTagConsumer, err := messaging.NewConsumer(rabbitmqConn, "post_tag_queue")
	if err != nil {
		log.Fatalf("Failed to create post tag consumer: %v", err)
	}

	messageConsumer, err := messaging.NewConsumer(rabbitmqConn, "message_queue")
	if err != nil {
		log.Fatalf("Failed to create message consumer: %v", err)
	}

	likeConsumer, err := messaging.NewConsumer(rabbitmqConn, "like_queue")
	if err != nil {
		log.Fatalf("Failed to create like consumer: %v", err)
	}

	// Start consumer
	uc := messaging.NewUserConsumer(userConsumer, userRepo)
	pc := messaging.NewPostConsumer(postConsumer, postRepo)
	cc := messaging.NewCommentConsumer(commentConsumer, commentRepo)
	tc := messaging.NewTagsConsumer(tagConsumer, tagRepo)
	ptc := messaging.NewPostTagsConsumer(postTagConsumer, postTagRepo)
	mc := messaging.NewMessageConsumer(messageConsumer, messageRepo)
	lc := messaging.NewLikeConsumer(likeConsumer, likeRepo)

	go func() {
		if err := uc.Start(); err != nil {
			log.Fatalf("Failed to start user consumer: %v", err)
		}
	}()

	go func() {
		if err := pc.Start(); err != nil {
			log.Fatalf("Failed to start post consumer: %v", err)
		}
	}()

	go func() {
		if err := cc.Start(); err != nil {
			log.Fatalf("Failed to start comment consumer: %v", err)
		}
	}()

	go func() {
		if err := tc.Start(); err != nil {
			log.Fatalf("Failed to start tag consumer: %v", err)
		}
	}()

	go func() {
		if err := ptc.Start(); err != nil {
			log.Fatalf("Failed to start post tag consumer: %v", err)
		}
	}()

	go func() {
		if err := mc.Start(); err != nil {
			log.Fatalf("Failed to start message consumer: %v", err)
		}
	}()

	go func() {
		if err := lc.Start(); err != nil {
			log.Fatalf("Failed to start like consumer: %v", err)
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
		log.Printf("Server starting on port %s", port)
		if err := router.Start(":" + port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// Context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}

	log.Println("Server exiting properly")
}
