package service

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type PostService struct {
	postRepo     repository.PostRepository
	userRepo     repository.UserRepository
	rabbitMQChan *amqp091.Channel
}

func NewPostService(postRepo repository.PostRepository, userRepo repository.UserRepository, rabbitMQChan *amqp091.Channel) *PostService {
	if postRepo == nil || userRepo == nil {
		log.Fatal("postRepo and userRepo must not be nil")
	}

	return &PostService{
		postRepo:     postRepo,
		userRepo:     userRepo,
		rabbitMQChan: rabbitMQChan,
	}
}

var logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func (s *PostService) CreatePost(ctx context.Context, post *entity.Post) error {
	if err := s.postRepo.Create(ctx, post); err != nil {
		return err
	}
	s.publishPostEvent("post_created", post)

	return nil
}

func (s *PostService) GetAllPosts(ctx context.Context) ([]*entity.Post, error) {
	return s.postRepo.FindAll(ctx)
}

func (s *PostService) GetPostByID(ctx context.Context, id uuid.UUID) ([]*entity.Post, error) {
	return s.postRepo.FindByID(ctx, id)
}

func (s *PostService) GetPostByTitle(ctx context.Context, title string) (*entity.Post, error) {
	return s.postRepo.FindByTitle(ctx, title)
}

func (s *PostService) UpdatePost(ctx context.Context, post *entity.Post, userID uuid.UUID) (*entity.Post, error) {
	if err := s.postRepo.Update(ctx, post, userID); err != nil {
		return nil, err
	}

	s.publishPostEvent("post_updated", post)

	return post, nil
}

func (s *PostService) DeletePost(ctx context.Context, id uuid.UUID) error {
	if err := s.postRepo.Delete(ctx, id); err != nil {
		return err
	}
	s.publishPostEvent("post_deleted", &entity.Post{PostID: id})
	return nil
}

func (s *PostService) publishPostEvent(eventType string, post *entity.Post) {
	event := map[string]interface{}{
		"type": eventType,
		"post": post,
	}
	eventJSON, _ := json.Marshal(event)
	err := s.rabbitMQChan.Publish(
		"post_events", // exchange
		"post",        // routing key
		false,         // mandatory
		false,         // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		})
	if err != nil {
		// Log the error, but don't stop the operation
		// TODO: Implement proper error logging
		log.Println("Failed to publish post event:", err)
	}
}
