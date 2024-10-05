package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
)

type PostTagsConsumer struct {
	consumer *Consumer
	repo     repository.PostTagsRepository
}

func NewPostTagsConsumer(consumer *Consumer, repo repository.PostTagsRepository) *PostTagsConsumer {
	return &PostTagsConsumer{
		consumer: consumer,
		repo:     repo,
	}
}

func (tc *PostTagsConsumer) Start() error {
	return tc.consumer.ConsumeMessages(tc.handleMessage)
}

func (tc *PostTagsConsumer) handleMessage(body []byte) error {
	var message struct {
		Action   string          `json:"action"`
		PostTags entity.PostTags `json:"posttags"`
		Post     entity.Post     `json:"post"`
		Tag      entity.Tags     `json:"tag"`
	}

	if err := json.Unmarshal(body, &message); err != nil {
		return err
	}

	ctx := context.Background()

	switch message.Action {
	case "create":
		return tc.repo.Create(ctx, &message.PostTags, message.Post.Title, message.Tag.Name)
	default:
		return fmt.Errorf("unknow action: %s", message.Action)
	}
}
