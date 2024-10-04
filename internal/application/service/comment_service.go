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

type CommentService struct {
	commentRepo  repository.CommentRepository
	postRepo     repository.PostRepository
	userRepo     repository.UserRepository
	rabbitMQChan *amqp091.Channel
}

func NewCommentService(
	commentRepo repository.CommentRepository,
	postRepo repository.PostRepository,
	userRepo repository.UserRepository,
	rabbitMQChan *amqp091.Channel,
) *CommentService {
	if commentRepo == nil || postRepo == nil || userRepo == nil {
		log.Fatal("commentRepo and postRepo and userRepo must not be nil")
	}

	return &CommentService{
		commentRepo:  commentRepo,
		postRepo:     postRepo,
		userRepo:     userRepo,
		rabbitMQChan: rabbitMQChan,
	}
}

func (s *CommentService) FindAllComments() ([]*entity.Comments, error) {
	return s.commentRepo.FindAll()
}

func (s *CommentService) FindCommentByID(ctx context.Context, commentID uuid.UUID) (*entity.Comments, error) {
	return s.commentRepo.FindByID(ctx, commentID)
}

func (s *CommentService) CreateComment(ctx context.Context, comment *entity.Comments, title, username string) error {
	if comment == nil {
		return errors.New("Comment is nil")
	}
	if title == "" {
		return errors.New("Title is empty")
	}
	if username == "" {
		return errors.New("Username is empty")
	}
	post, err := s.postRepo.FindByTitle(title)
	if err != nil {
		return fmt.Errorf("Failed to find post by title: %s: %w", title, err)
	}
	comment.PostID = post.PostID

	if err := s.commentRepo.Create(ctx, comment, title, username); err != nil {
		return err
	}
	s.publishCommentEvent("comment_created", comment)
	return nil
}

func (s *CommentService) UpdateComment(ctx context.Context, comment *entity.Comments, commentID uuid.UUID) (*entity.Comments, error) {
	if comment == nil {
		return nil, errors.New("Comment is nil")
	}
	if commentID == uuid.Nil {
		return nil, errors.New("CommentID is nil")
	}
	if err := s.commentRepo.Update(ctx, comment, commentID); err != nil {
		return nil, err
	}
	s.publishCommentEvent("comment_updated", comment)
	return comment, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, commentID uuid.UUID) error {
	if commentID == uuid.Nil {
		return errors.New("CommentID is nil")
	}
	if err := s.commentRepo.Delete(ctx, commentID); err != nil {
		return err
	}
	s.publishCommentEvent("comment_deleted", &entity.Comments{CommentID: commentID})
	return nil
}

func (s *CommentService) publishCommentEvent(eventType string, comment *entity.Comments) {
	event := map[string]interface{}{
		"type":    eventType,
		"comment": comment,
	}
	eventJSON, _ := json.Marshal(event)
	err := s.rabbitMQChan.Publish(
		"comment_events", // exchange
		"comment",        // routing key
		false,            // mandatory
		false,            // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		})
	if err != nil {
		// Log the error, but don't stop the operation
		// TODO: Implement proper error logging
		log.Println("Failed to publish comment event:", err)
	}
}
