package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
)

type PostConsumer struct {
	consumer *Consumer
	repo     repository.PostRepository
}

func NewPostConsumer(consumer *Consumer, repo repository.PostRepository) *PostConsumer {
	return &PostConsumer{
		consumer: consumer,
		repo:     repo,
	}
}

func (pc *PostConsumer) Start() error {
	return pc.consumer.ConsumeMessages(pc.handleMessage)
}

func (pc *PostConsumer) handleMessage(body []byte) error {
	ctx := context.Background()
	var message struct {
		Action string      `json:"action"`
		Data   entity.Post `json:"data"`
	}
	if err := json.Unmarshal(body, &message); err != nil {
		return err
	}

	switch message.Action {
	case "create":
		return pc.repo.Create(ctx, &message.Data, message.Data.User.Username)
	case "update":
		return pc.repo.Update(ctx, &message.Data)
	case "delete":
		return pc.repo.Delete(ctx, message.Data.PostID)
	default:
		return fmt.Errorf("unknown action: %s", message.Action)
	}
}
