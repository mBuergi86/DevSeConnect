package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"gorm.io/gorm"
)

type MessageRepository interface {
	FindAll() ([]*entity.Messages, error)
	FindByID(ctx context.Context, messageID uuid.UUID) (*entity.Messages, error)
	Create(ctx context.Context, message *entity.Messages, username1, username2 string) error
	Update(ctx context.Context, message *entity.Messages, messageID uuid.UUID) error
	Delete(ctx context.Context, messageID uuid.UUID) error
}

type PostgresMessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &PostgresMessageRepository{DB: db}
}

func (r *PostgresMessageRepository) FindAll() ([]*entity.Messages, error) {
	var messages []*entity.Messages
	if err := r.DB.Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *PostgresMessageRepository) FindByID(ctx context.Context, messageID uuid.UUID) (*entity.Messages, error) {
	var message entity.Messages
	if err := r.DB.First(&message, "message_id = ?", messageID).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *PostgresMessageRepository) Create(ctx context.Context, messages *entity.Messages, username1, username2 string) error {
	tx := r.DB.Begin()

	var user1 entity.User
	if err := tx.Find(&user1, "username = ?", username1).Error; err != nil {
		tx.Rollback()
		return err
	}

	var user2 entity.User
	if err := tx.Find(&user2, "username = ?", username2).Error; err != nil {
		tx.Rollback()
		return err
	}

	messages.SenderID = user1.UserID
	messages.ReceiverID = user2.UserID

	if err := tx.Create(messages).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *PostgresMessageRepository) Update(ctx context.Context, messages *entity.Messages, messageID uuid.UUID) error {
	tx := r.DB.Begin()

	if err := tx.Model(&entity.Messages{}).
		Where("message_id = ?", messageID).
		Updates(map[string]interface{}{
			"content": messages.Content,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *PostgresMessageRepository) Delete(ctx context.Context, messageID uuid.UUID) error {
	tx := r.DB.Begin()
	if err := tx.Delete(&entity.Messages{}, "message_id = ?", messageID).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
