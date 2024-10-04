package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/rabbitmq/amqp091-go"
)

type TagsService struct {
	tagsRepo     repository.TagsRepository
	rabbitMQChan *amqp091.Channel
}

func NewTagsService(tagsRepo repository.TagsRepository, rabbitMQChan *amqp091.Channel) *TagsService {
	if tagsRepo == nil {
		panic("tagsRepo is required")
	}

	return &TagsService{
		tagsRepo:     tagsRepo,
		rabbitMQChan: rabbitMQChan,
	}
}

func (s *TagsService) GetTags(ctx context.Context) ([]*entity.Tags, error) {
	return s.tagsRepo.FindAll(ctx)
}

func (s *TagsService) GetTag(ctx context.Context, tagID uuid.UUID) (*entity.Tags, error) {
	return s.tagsRepo.FindByID(ctx, tagID)
}

func (s *TagsService) CreateTag(ctx context.Context, tag *entity.Tags) error {
	return s.tagsRepo.Create(ctx, tag)
}

func (s *TagsService) DeleteTag(ctx context.Context, tagID uuid.UUID) error {
	return s.tagsRepo.Delete(ctx, tagID)
}

func (s *TagsService) publishTagEvent(eventType string, tag *entity.Tags) {
	event := map[string]interface{}{
		"type": eventType,
		"tag":  tag,
	}
	eventJSON, _ := json.Marshal(event)
	err := s.rabbitMQChan.Publish(
		"tag_events", // exchange
		"tag",        // routing key
		false,        // mandatory
		false,        // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		})
	if err != nil {
		// Log the error, but don't stop the operation
		// TODO: Implement proper error logging
		log.Println("Failed to publish tag event:", err)
	}
}
