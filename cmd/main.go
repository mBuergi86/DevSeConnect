package main

import (
	"log"
	"os"

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

	// Setup consumer
	userConsumer, err := messaging.NewConsumer(rabbitmqConn, userRepo)
	if err != nil {
		log.Fatalf("Failed to create user consumer: %v", err)
	}

	// Setup services
	userService, err := service.NewUserService(userRepo, rabbitMQChan)
	if err != nil {
		log.Fatalf("Failed to create user service: %v", err)
	}

	// Start consumer
	go func() {
		if err := userConsumer.Start(); err != nil {
			log.Fatalf("Failed to start user consumer: %v", err)
		}
	}()

	// Setup router
	router := routing.SetupRouter(userService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Start(":" + port))
}
