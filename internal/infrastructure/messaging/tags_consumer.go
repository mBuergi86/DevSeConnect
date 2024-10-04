package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
)

type TagsConsumer struct {
	consumer *Consumer
	repo     repository.TagsRepository
}

func NewTagsConsumer(consumer *Consumer, repo repository.TagsRepository) *TagsConsumer {
	return &TagsConsumer{
		consumer: consumer,
		repo:     repo,
	}
}

func (tc *TagsConsumer) Start() error {
	return tc.consumer.ConsumeMessages(tc.handleMessage)
}

func (tc *TagsConsumer) handleMessage(body []byte) error {
	var message struct {
		Action string      `json:"action"`
		Data   entity.Tags `json:"data"`
	}

	if err := json.Unmarshal(body, &message); err != nil {
		return err
	}

	ctx := context.Background()

	switch message.Action {
	case "create":
		return tc.repo.Create(ctx, &message.Data)
	case "delete":
		return tc.repo.Delete(ctx, message.Data.TagID)
	default:
		return fmt.Errorf("unknow action: %s", message.Action)
	}
}
