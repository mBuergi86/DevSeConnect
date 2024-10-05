package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
)

type UserConsumer struct {
	consumer *Consumer
	repo     repository.UserRepository
}

func NewUserConsumer(consumer *Consumer, repo repository.UserRepository) *UserConsumer {
	return &UserConsumer{
		consumer: consumer,
		repo:     repo,
	}
}

func (uc *UserConsumer) Start() error {
	return uc.consumer.ConsumeMessages(uc.handleMessage)
}

func (uc *UserConsumer) handleMessage(body []byte) error {
	ctx := context.Background()
	var message struct {
		Action string      `json:"action"`
		Data   entity.User `json:"data"`
	}

	if err := json.Unmarshal(body, &message); err != nil {
		return err
	}

	switch message.Action {
	case "create":
		return uc.repo.Create(ctx, &message.Data)
	case "update":
		return uc.repo.Update(ctx, &message.Data)
	case "delete":
		return uc.repo.Delete(ctx, message.Data.UserID)
	default:
		return fmt.Errorf("unknown action: %s", message.Action)
	}
}
