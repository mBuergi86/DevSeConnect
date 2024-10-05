package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/rabbitmq/amqp091-go"
)

type LikeService struct {
	likeRepo     repository.LikeRepository
	postRepo     repository.PostRepository
	commentRepo  repository.CommentRepository
	userRepo     repository.UserRepository
	rabbitMQChan *amqp091.Channel
}

func NewLikeService(
	likeRepo repository.LikeRepository,
	postRepo repository.PostRepository,
	commentRepo repository.CommentRepository,
	userRepo repository.UserRepository,
	rabbitMQChan *amqp091.Channel,
) *LikeService {
	if likeRepo == nil || postRepo == nil || commentRepo == nil || userRepo == nil {
		log.Fatal("likeRepo and postRepo and commentRepo and userRepo must not be nil")
	}

	return &LikeService{
		likeRepo:     likeRepo,
		postRepo:     postRepo,
		commentRepo:  commentRepo,
		userRepo:     userRepo,
		rabbitMQChan: rabbitMQChan,
	}
}

func (s *LikeService) FindAllLikes(ctx context.Context) ([]*entity.Likes, error) {
	return s.likeRepo.FindAll(ctx)
}

func (s *LikeService) FindLikeByID(ctx context.Context, likeID uuid.UUID) (*entity.Likes, error) {
	return s.likeRepo.FindByID(ctx, likeID)
}

func (s *LikeService) CreateByPost(ctx context.Context, like *entity.Likes, title, username string) error {
	if like == nil {
		return errors.New("Like is nil")
	}
	if title == "" {
		return errors.New("Title is empty")
	}
	if username == "" {
		return errors.New("Username is empty")
	}

	var post *entity.Post
	var user *entity.User

	post, err := s.postRepo.FindByTitle(ctx, title)
	if err != nil {
		slog.Error("Failed to find post by title", "title", title, "error", err)
		return fmt.Errorf("failed to find post by title: %s: %w", title, err)
	}

	like.PostID = post.PostID

	user, err = s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		slog.Error("Failed to find user by username", "username", username, "error", err)
		return fmt.Errorf("failed to find user by username: %s: %w", username, err)
	}

	like.UserID = user.UserID

	if err := s.likeRepo.CreateByComment(ctx, like, title, username); err != nil {
		return err
	}
	s.publishLikeEvent("like_createdByPost", like)
	return nil
}

func (s *LikeService) CreateByComment(ctx context.Context, like *entity.Likes, content, username string) error {
	if like == nil {
		return errors.New("Like is nil")
	}
	if content == "" {
		return errors.New("Content is empty")
	}
	if username == "" {
		return errors.New("Username is empty")
	}

	var comment *entity.Comments
	var user *entity.User

	comment, err := s.commentRepo.FindByContent(ctx, content)
	if err != nil {
		slog.Error("Failed to find comment by content", "content", content, "error", err)
		return fmt.Errorf("failed to find comment by content: %s: %w", content, err)
	}

	like.CommentID = comment.CommentID

	user, err = s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		slog.Error("Failed to find user by username", "username", username, "error", err)
		return fmt.Errorf("failed to find user by username: %s: %w", username, err)
	}

	like.UserID = user.UserID

	if err := s.likeRepo.CreateByComment(ctx, like, content, username); err != nil {
		return err
	}
	s.publishLikeEvent("like_createdByComment", like)
	return nil
}

func (s *LikeService) DeleteLike(ctx context.Context, likeID uuid.UUID) error {
	if likeID == uuid.Nil {
		return errors.New("LikeID is nil")
	}
	if err := s.likeRepo.Delete(ctx, likeID); err != nil {
		return err
	}
	s.publishLikeEvent("like_deleted", &entity.Likes{LikeID: likeID})
	return nil
}

func (s *LikeService) publishLikeEvent(eventType string, like *entity.Likes) {
	event := map[string]interface{}{
		"type": eventType,
		"like": like,
	}
	eventJSON, _ := json.Marshal(event)
	err := s.rabbitMQChan.Publish(
		"like_events", // exchange
		"like",        // routing key
		false,         // mandatory
		false,         // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		})
	if err != nil {
		// Log the error, but don't stop the operation
		// TODO: Implement proper error logging
		slog.Debug("Failed to publish like event:", "error", err)
	}
}
