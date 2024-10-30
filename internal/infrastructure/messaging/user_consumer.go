package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/models"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/rs/zerolog"
)

type UserConsumer struct {
	consumer *Consumer
	repo     repository.UserRepository
	logger   zerolog.Logger
}

func NewUserConsumer(consumer *Consumer, repo repository.UserRepository) *UserConsumer {
	return &UserConsumer{
		consumer: consumer,
		repo:     repo,
		logger:   zerolog.New(os.Stdout).With().Str("component", "user_consumer").Logger(),
	}
}

func (uc *UserConsumer) Start(ctx context.Context) error {
	return uc.consumer.ConsumeMessages(uc.handleMessage)
}

func (uc *UserConsumer) handleMessage(body []byte) error {
	if len(body) == 0 {
		uc.logger.Warn().Msgf("Received empty message body, sending to DLQ: %s", string(body))
		return fmt.Errorf("received empty message")
	}

	var message models.EventMessage
	if err := json.Unmarshal(body, &message); err != nil {
		uc.logger.Error().Err(err).Msg("Failed to unmarshal user message")
		return err
	}

	if message.Action == "" {
		uc.logger.Warn().Msg("Empty Action in message, cannot proceed")
		return fmt.Errorf("message action is empty")
	}

	var user entity.User
	if err := json.Unmarshal(message.Data, &user); err != nil {
		uc.logger.Error().Err(err).Msg("Failed to unmarshal user data from message")
		return err
	}

	ctx := context.Background()

	switch message.Action {
	case "user_registered":
		return uc.handleUserRegistered(ctx, &user)
	case "user_updated":
		return uc.handleUserUpdated(ctx, &user)
	case "user_deleted":
		return uc.handleUserDeleted(ctx, user)
	default:
		err := fmt.Errorf("unknown action: %s", message.Action)
		uc.logger.Error().Err(err).Msg("Unknown action type received")
		return err
	}
}

func (uc *UserConsumer) handleUserRegistered(ctx context.Context, user *entity.User) error {
	if err := uc.repo.Create(ctx, user); err != nil {
		uc.logger.Error().Err(err).Msg("Failed to create user")
		return err
	}
	uc.logger.Info().Msg("User created successfully")
	return nil
}

func (uc *UserConsumer) handleUserUpdated(ctx context.Context, user *entity.User) error {
	if err := uc.repo.Update(ctx, user); err != nil {
		uc.logger.Error().Err(err).Msg("Failed to update user")
		return err
	}
	uc.logger.Info().Msg("User updated successfully")
	return nil
}

func (uc *UserConsumer) handleUserDeleted(ctx context.Context, user entity.User) error {
	if err := uc.repo.Delete(ctx, user.UserID); err != nil {
		uc.logger.Error().Err(err).Msg("Failed to delete user")
		return err
	}
	uc.logger.Info().Msg("User deleted successfully")
	return nil
}
