package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
)

type LikeConsumer struct {
	consumer *Consumer
	repo     repository.LikeRepository
}

func NewLikeConsumer(consumer *Consumer, repo repository.LikeRepository) *LikeConsumer {
	return &LikeConsumer{
		consumer: consumer,
		repo:     repo,
	}
}

func (tc *LikeConsumer) Start() error {
	return tc.consumer.ConsumeMessages(tc.handleLike)
}

func (tc *LikeConsumer) handleLike(body []byte) error {
	var like struct {
		Action string       `json:"action"`
		Like   entity.Likes `json:"like"`
		Post   entity.Post  `json:"post"`
		User   entity.User  `json:"user"`
	}

	if err := json.Unmarshal(body, &like); err != nil {
		return err
	}

	ctx := context.Background()

	switch like.Action {
	case "createByPost":
		return tc.repo.CreateByPost(ctx, &like.Like, like.Post.Title, like.User.Username)
	case "createByComment":
		return tc.repo.CreateByComment(ctx, &like.Like, like.Post.Content, like.User.Username)
	case "delete":
		return tc.repo.Delete(ctx, like.Like.LikeID)
	default:
		return fmt.Errorf("unknow action: %s", like.Action)
	}
}
