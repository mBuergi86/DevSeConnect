package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
)

type MessageConsumer struct {
	consumer *Consumer
	repo     repository.MessageRepository
}

func NewMessageConsumer(consumer *Consumer, repo repository.MessageRepository) *MessageConsumer {
	return &MessageConsumer{
		consumer: consumer,
		repo:     repo,
	}
}

func (tc *MessageConsumer) Start() error {
	return tc.consumer.ConsumeMessages(tc.handleMessage)
}

func (tc *MessageConsumer) handleMessage(body []byte) error {
	var message struct {
		Action  string          `json:"action"`
		Message entity.Messages `json:"messages"`
		User    entity.User     `json:"user"`
	}

	if err := json.Unmarshal(body, &message); err != nil {
		return err
	}

	ctx := context.Background()

	switch message.Action {
	case "create":
		return tc.repo.Create(ctx, &message.Message, message.User.Username, message.User.Username)
	case "update":
		return tc.repo.Update(ctx, &message.Message, message.Message.MessageID)
	case "delete":
		return tc.repo.Delete(ctx, message.Message.MessageID)
	default:
		return fmt.Errorf("unknow action: %s", message.Action)
	}
}
