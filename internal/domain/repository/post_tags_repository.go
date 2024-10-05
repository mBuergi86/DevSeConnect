package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"gorm.io/gorm"
)

type PostTagsRepository interface {
	FindAll(ctx context.Context) ([]*entity.PostTags, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.PostTags, error)
	Create(ctx context.Context, posttags *entity.PostTags, title, tags string) error
}

type PostgresPostTagsRepository struct {
	DB *gorm.DB
}

func NewPostTagsRepository(db *gorm.DB) PostTagsRepository {
	return &PostgresPostTagsRepository{DB: db}
}

func (r *PostgresPostTagsRepository) FindAll(ctx context.Context) ([]*entity.PostTags, error) {
	tx := r.DB.Begin()
	var posttags []*entity.PostTags
	err := tx.Find(&posttags).Error
	if err != nil {
		return nil, err
	}

	return posttags, nil
}

func (r *PostgresPostTagsRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.PostTags, error) {
	var posttags entity.PostTags
	err := r.DB.Preload("Post_Tags").First(&posttags, "post_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &posttags, nil
}

func (r *PostgresPostTagsRepository) Create(ctx context.Context, posttags *entity.PostTags, title, tags string) error {
	tx := r.DB.Begin()
	var post entity.Post
	if err := tx.First(&post, "title = ?", title).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to find post with title %s: %w", title, err)
	}

	var tag entity.Tags
	if err := tx.First(&tag, "tags = ?", tags).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to find tag with tags %s: %w", tags, err)
	}

	posttags.PostID = post.PostID
	posttags.TagID = tag.TagID

	err := tx.Create(posttags).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
