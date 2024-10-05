package service

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/rabbitmq/amqp091-go"
)

type PostTagsService struct {
	posttagsRepo repository.PostTagsRepository
	rabbitMQChan *amqp091.Channel
}

func NewPostTagsService(posttagsRepo repository.PostTagsRepository, rabbitMQChan *amqp091.Channel) *PostTagsService {
	if posttagsRepo == nil {
		panic("tagsRepo is required")
	}

	return &PostTagsService{
		posttagsRepo: posttagsRepo,
		rabbitMQChan: rabbitMQChan,
	}
}

func (s *PostTagsService) GetPostTags(ctx context.Context) ([]*entity.PostTags, error) {
	return s.posttagsRepo.FindAll(ctx)
}

func (s *PostTagsService) GetTag(ctx context.Context, tagID uuid.UUID) (*entity.PostTags, error) {
	return s.posttagsRepo.FindByID(ctx, tagID)
}

func (s *PostTagsService) CreateTag(ctx context.Context, tag *entity.PostTags, title, tags string) error {
	return s.posttagsRepo.Create(ctx, tag, title, tags)
}

func (s *PostTagsService) publishTagEvent(eventType string, post_tag *entity.PostTags) {
	event := map[string]interface{}{
		"type":     eventType,
		"post_tag": post_tag,
	}
	eventJSON, _ := json.Marshal(event)
	err := s.rabbitMQChan.Publish(
		"post_tag_events", // exchange
		"post_tag",        // routing key
		false,             // mandatory
		false,             // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		})
	if err != nil {
		// Log the error, but don't stop the operation
		// TODO: Implement proper error logging
		slog.Debug("Failed to publish tag event:", err)
	}
}
