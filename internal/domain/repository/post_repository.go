package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindAll(ctx context.Context) ([]*entity.Post, error)
	FindByID(ctx context.Context, id uuid.UUID) ([]*entity.Post, error)
	FindByTitle(ctx context.Context, title string) (*entity.Post, error)
	Create(ctx context.Context, post *entity.Post) error
	Update(ctx context.Context, post *entity.Post, userID uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PostgresPostRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostgresPostRepository{DB: db}
}

func (r *PostgresPostRepository) FindAll(ctx context.Context) ([]*entity.Post, error) {
	var posts []*entity.Post
	err := r.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, err
}

func (r *PostgresPostRepository) FindByID(ctx context.Context, id uuid.UUID) ([]*entity.Post, error) {
	var post []*entity.Post
	err := r.DB.Preload("User").Find(&post, "user_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return post, err
}

func (r *PostgresPostRepository) FindByTitle(ctx context.Context, title string) (*entity.Post, error) {
	var post entity.Post
	err := r.DB.Preload("User").First(&post, "title = ?", title).Error
	if err != nil {
		return nil, err
	}

	return &post, err
}

func (r *PostgresPostRepository) Create(ctx context.Context, post *entity.Post) error {
	tx := r.DB.Begin()

	if err := tx.Create(post).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to create post: %w", err)
	}

	return tx.Commit().Error
}

func (r *PostgresPostRepository) Update(ctx context.Context, post *entity.Post, userID uuid.UUID) error {
	tx := r.DB.Begin()
	if err := tx.Model(&entity.Post{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"title":   post.Title,
			"content": post.Content,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *PostgresPostRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tx := r.DB.Begin()
	if err := tx.Delete(&entity.Post{}, "post_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
