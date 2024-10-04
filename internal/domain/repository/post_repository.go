package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindAll() ([]*entity.Post, error)
	FindByID(id uuid.UUID) (*entity.Post, error)
	FindByTitle(title string) (*entity.Post, error)
	Create(post *entity.Post, username string) error
	Update(post *entity.Post) error
	Delete(id uuid.UUID) error
}

type PostgresPostRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostgresPostRepository{DB: db}
}

func (r *PostgresPostRepository) FindAll() ([]*entity.Post, error) {
	var posts []*entity.Post
	err := r.DB.Find(&posts).Error
	return posts, err
}

func (r *PostgresPostRepository) FindByID(id uuid.UUID) (*entity.Post, error) {
	var post entity.Post
	err := r.DB.Preload("User").Find(&post, "post_id = ?", id).Error
	return &post, err
}

func (r *PostgresPostRepository) FindByTitle(title string) (*entity.Post, error) {
	var post entity.Post
	err := r.DB.Preload("User").Find(&post, "title = ?", title).Error
	return &post, err
}

func (r *PostgresPostRepository) Create(post *entity.Post, username string) error {
	tx := r.DB.Begin()
	var user entity.User
	if err := tx.First(&user, "username = ?", username).Error; err != nil {
		return fmt.Errorf("Failed to find user wit username %s: %w", username, err)
	}

	post.UserID = user.UserID
	if err := tx.Create(post).Error; err != nil {
		return fmt.Errorf("Failed to create post: %w", err)
	}

	return tx.Commit().Error
}

func (r *PostgresPostRepository) Update(post *entity.Post) error {
	tx := r.DB.Begin()
	if err := tx.Save(post).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *PostgresPostRepository) Delete(id uuid.UUID) error {
	tx := r.DB.Begin()
	if err := tx.Delete(&entity.Post{}, "post_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
