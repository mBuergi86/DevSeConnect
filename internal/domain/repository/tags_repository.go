package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"gorm.io/gorm"
)

type TagsRepository interface {
	FindAll(ctx context.Context) ([]*entity.Tags, error)
	FindByID(ctx context.Context, tagID uuid.UUID) (*entity.Tags, error)
	Create(ctx context.Context, tag *entity.Tags) error
	Delete(ctx context.Context, tagID uuid.UUID) error
}

type PostgresTagsRepository struct {
	DB *gorm.DB
}

func NewTagsRepository(db *gorm.DB) *PostgresTagsRepository {
	if db == nil {
		panic("nil db")
	}

	return &PostgresTagsRepository{
		DB: db,
	}
}

func (r *PostgresTagsRepository) FindAll(ctx context.Context) ([]*entity.Tags, error) {
	var tags []*entity.Tags
	if err := r.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *PostgresTagsRepository) FindByID(ctx context.Context, tagID uuid.UUID) (*entity.Tags, error) {
	var tag entity.Tags
	tx := r.DB.Begin()
	if err := tx.Find(&tag, "tag_id = ?", tagID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &tag, nil
}

func (r *PostgresTagsRepository) Create(ctx context.Context, tag *entity.Tags) error {
	tx := r.DB.Begin()
	return tx.Create(tag).Error
}

func (r *PostgresTagsRepository) Delete(ctx context.Context, tagID uuid.UUID) error {
	tx := r.DB.Begin()
	return tx.Delete(&entity.Tags{}, tagID).Error
}
