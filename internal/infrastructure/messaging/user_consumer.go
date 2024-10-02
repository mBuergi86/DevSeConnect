package messaging

import (
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

func (uc *UserConsumer) handleMessage(body []byte) error {
	var message struct {
		Action string      `json:"action"`
		Data   entity.User `json:"data"`
	}

	if err := json.Unmarshal(body, &message); err != nil {
		return err
	}

	switch message.Action {
	case "create":
		return uc.repo.Create(&message.Data)
	case "update":
		return uc.repo.Update(&message.Data)
	case "delete":
		return uc.repo.Delete(message.Data.UserID)
	default:
		return fmt.Errorf("unknown action: %s", message.Action)
	}
}
