package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"gorm.io/gorm"
)

type LikeRepository interface {
	FindAll(ctx context.Context) ([]*entity.Likes, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Likes, error)
	CreateByPost(ctx context.Context, like *entity.Likes, title, username string) error
	CreateByComment(ctx context.Context, like *entity.Likes, content, username string) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PostgresLikeRepository struct {
	DB *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &PostgresLikeRepository{DB: db}
}

func (r *PostgresLikeRepository) FindAll(ctx context.Context) ([]*entity.Likes, error) {
	var likes []*entity.Likes
	err := r.DB.Find(&likes).Error
	if err != nil {
		return nil, err
	}

	return likes, err
}

func (r *PostgresLikeRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Likes, error) {
	var like entity.Likes
	err := r.DB.Find(&like, "like_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &like, err
}

func (r *PostgresLikeRepository) CreateByPost(ctx context.Context, like *entity.Likes, title, username string) error {
	tx := r.DB.Begin()
	var post entity.Post
	if err := tx.First(&post, "title = ?", title).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to find post with title %s: %w", title, err)
	}

	var user entity.User
	if err := tx.First(&user, "username = ?", username).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to find user with username %s: %w", username, err)
	}

	like.PostID = post.PostID
	like.UserID = user.UserID
	if err := tx.Create(like).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to create by Post likes: %w", err)
	}

	return tx.Commit().Error
}

func (r *PostgresLikeRepository) CreateByComment(ctx context.Context, like *entity.Likes, content, username string) error {
	tx := r.DB.Begin()
	var comment entity.Comments
	if err := tx.First(&comment, "content = ?", content).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to find comment with content %s: %w", content, err)
	}

	var user entity.User
	if err := tx.First(&user, "username = ?", username).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to find user wit username %s: %w", username, err)
	}

	like.CommentID = comment.CommentID
	like.UserID = user.UserID
	if err := tx.Create(like).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to create by Comment likes : %w", err)
	}

	return tx.Commit().Error
}

func (r *PostgresLikeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tx := r.DB.Begin()
	if err := tx.Delete(&entity.Likes{}, "like_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
