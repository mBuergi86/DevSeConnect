package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
)

type CommentConsumer struct {
	consumer *Consumer
	repo     repository.CommentRepository
}

func NewCommentConsumer(consumer *Consumer, repo repository.CommentRepository) *CommentConsumer {
	return &CommentConsumer{
		consumer: consumer,
		repo:     repo,
	}
}

func (cc *CommentConsumer) Start() error {
	return cc.consumer.ConsumeMessages(cc.handleMessage)
}

func (cc *CommentConsumer) handleMessage(body []byte) error {
	var message struct {
		Action string          `json:"action"`
		Data   entity.Comments `json:"data"`
	}
	if err := json.Unmarshal(body, &message); err != nil {
		return err
	}

	ctx := context.Background()

	switch message.Action {
	case "create":
		return cc.repo.Create(ctx, &message.Data, message.Data.Post.Title, message.Data.User.Username)
	case "update":
		return cc.repo.Update(ctx, &message.Data, message.Data.CommentID)
	case "delete":
		return cc.repo.Delete(ctx, message.Data.CommentID)
	default:
		return fmt.Errorf("unknown action: %s", message.Action)
	}
}
