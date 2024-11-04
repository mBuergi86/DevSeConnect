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
	maxRetries := 3
	retries := 0

	for retries < maxRetries {
		var message models.EventMessage
		err := json.Unmarshal(body, &message)
		if err != nil {
			retries++
			uc.logger.Error().Err(err).Msgf("Failed to unmarshal message, attempt %d", retries)
			if retries == maxRetries {
				uc.logger.Error().Msg("Max retries reached, moving to DLQ")
				return fmt.Errorf("failed to process message after %d attempts", maxRetries)
			}
			continue
		}

		err = uc.processAction(message)
		if err == nil {
			return nil
		}

		retries++
		if retries == maxRetries {
			uc.logger.Error().Err(err).Msg("Moving to DLQ after max retries")
			return fmt.Errorf("message failed after max retries: %w", err)
		}
	}
	return nil
}

func (uc *UserConsumer) processAction(message models.EventMessage) error {
	var user entity.User
	err := json.Unmarshal(message.Data, &user)
	if err != nil {
		return err
	}

	switch message.Action {
	case "user_registered":
		return uc.handleUserRegistered(context.Background(), &user)
	case "user_updated":
		return uc.handleUserUpdated(context.Background(), &user)
	case "user_deleted":
		return uc.handleUserDeleted(context.Background(), user)
	default:
		return fmt.Errorf("unknown action: %s", message.Action)
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
