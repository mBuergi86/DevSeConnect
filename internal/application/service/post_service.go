package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/rabbitmq/amqp091-go"
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

func (s *PostService) CreatePost(ctx context.Context, post *entity.Post, username string) error {
	if post == nil {
		return errors.New("Post is nil")
	}
	if username == "" {
		return errors.New("Username is empty")
	}

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return fmt.Errorf("Failed to find user by username: %s: %w", username, err)
	}

	post.UserID = user.UserID

	if err := s.postRepo.Create(post, username); err != nil {
		return err
	}
	s.publishPostEvent("post_created", post)

	return nil
}

func (s *PostService) GetAllPosts() ([]*entity.Post, error) {
	return s.postRepo.FindAll()
}

func (s *PostService) GetPostByID(id uuid.UUID) (*entity.Post, error) {
	return s.postRepo.FindByID(id)
}

func (s *PostService) GetPostByTitle(title string) (*entity.Post, error) {
	return s.postRepo.FindByTitle(title)
}

func (s *PostService) UpdatePost(updateData map[string]interface{}) (*entity.Post, error) {
	postID, ok := updateData["post_id"].(uuid.UUID)
	if !ok {
		return nil, errors.New("Invalid post ID")
	}

	existingPost, err := s.postRepo.FindByID(postID)
	if err != nil {
		return nil, err
	}

	// Actually update the post with the new data
	if userID, ok := updateData["user_id"].(uuid.UUID); ok {
		existingPost.UserID = userID
	}
	if title, ok := updateData["title"].(string); ok {
		existingPost.Title = title
	}
	if content, ok := updateData["content"].(string); ok {
		existingPost.Content = content
	}
	if mediaType, ok := updateData["media_type"].(string); ok {
		existingPost.MediaType = mediaType
	}
	if mediaURL, ok := updateData["media_url"].(string); ok {
		existingPost.MediaURL = mediaURL
	}
	if isDeleted, ok := updateData["is_deleted"].(bool); ok {
		existingPost.IsDeleted = isDeleted
	}

	if err := s.postRepo.Update(existingPost); err != nil {
		return nil, err
	}

	s.publishPostEvent("post_updated", existingPost)

	return existingPost, nil
}

func (s *PostService) DeletePost(id uuid.UUID) error {
	if err := s.postRepo.Delete(id); err != nil {
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
