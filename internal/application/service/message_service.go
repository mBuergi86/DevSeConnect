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

type MessageService struct {
	messageRepo  repository.MessageRepository
	userRepo     repository.UserRepository
	rabbitMQChan *amqp091.Channel
}

func NewMessageService(
	messageRepo repository.MessageRepository,
	userRepo repository.UserRepository,
	rabbitMQChan *amqp091.Channel,
) *MessageService {
	if messageRepo == nil || userRepo == nil {
		log.Fatal("messageRepo and userRepo must not be nil")
	}

	return &MessageService{
		messageRepo:  messageRepo,
		userRepo:     userRepo,
		rabbitMQChan: rabbitMQChan,
	}
}

func (s *MessageService) FindAllMessages(ctx context.Context) ([]*entity.Messages, error) {
	return s.messageRepo.FindAll(ctx)
}

func (s *MessageService) FindMessageByID(ctx context.Context, messageID uuid.UUID) (*entity.Messages, error) {
	return s.messageRepo.FindByID(ctx, messageID)
}

func (s *MessageService) CreateMessage(ctx context.Context, message *entity.Messages, username1, username2 string) error {
	if message == nil {
		return errors.New("Message is nil")
	}
	if username1 == "" {
		return errors.New("Sender is empty")
	}
	if username2 == "" {
		return errors.New("Receiver is empty")
	}

	var user1, user2 *entity.User

	user1, err := s.userRepo.FindByUsername(ctx, username1)
	if err != nil {
		slog.Error("Failed to find user by username", "username", username1, "error", err)
		return fmt.Errorf("failed to find user by username: %s: %w", username1, err)
	}

	message.SenderID = user1.UserID

	user2, err = s.userRepo.FindByUsername(ctx, username2)
	if err != nil {
		slog.Error("Failed to find user by username", "username", username2, "error", err)
		return fmt.Errorf("failed to find user by username: %s: %w", username2, err)
	}

	message.ReceiverID = user2.UserID

	if err := s.messageRepo.Create(ctx, message, username1, username2); err != nil {
		return err
	}
	s.publishMessageEvent("message_created", message)
	return nil
}

func (s *MessageService) UpdateMessage(ctx context.Context, message *entity.Messages, messageID uuid.UUID) (*entity.Messages, error) {
	if message == nil {
		return nil, errors.New("Message is nil")
	}
	if messageID == uuid.Nil {
		return nil, errors.New("MessageID is nil")
	}
	if err := s.messageRepo.Update(ctx, message, messageID); err != nil {
		return nil, err
	}
	s.publishMessageEvent("message_updated", message)
	return message, nil
}

func (s *MessageService) DeleteMessage(ctx context.Context, messageID uuid.UUID) error {
	if messageID == uuid.Nil {
		return errors.New("MessageID is nil")
	}
	if err := s.messageRepo.Delete(ctx, messageID); err != nil {
		return err
	}
	s.publishMessageEvent("message_deleted", &entity.Messages{MessageID: messageID})
	return nil
}

func (s *MessageService) publishMessageEvent(eventType string, message *entity.Messages) {
	event := map[string]interface{}{
		"type":    eventType,
		"message": message,
	}
	eventJSON, _ := json.Marshal(event)
	err := s.rabbitMQChan.Publish(
		"message_events", // exchange
		"message",        // routing key
		false,            // mandatory
		false,            // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		})
	if err != nil {
		// Log the error, but don't stop the operation
		// TODO: Implement proper error logging
		slog.Debug("Failed to publish message event:", "error", err)
	}
}
