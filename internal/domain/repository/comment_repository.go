package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"gorm.io/gorm"
)

type CommentRepository interface {
	FindAll() ([]*entity.Comments, error)
	FindByID(ctx context.Context, commentID uuid.UUID) (*entity.Comments, error)
	Create(ctx context.Context, comment *entity.Comments, title, username string) error
	Update(ctx context.Context, comment *entity.Comments, commentID uuid.UUID) error
	Delete(ctx context.Context, commentID uuid.UUID) error
}

type PostgresCommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &PostgresCommentRepository{DB: db}
}

func (r *PostgresCommentRepository) FindAll() ([]*entity.Comments, error) {
	var comments []*entity.Comments
	tx := r.DB.Begin()
	if err := tx.Find(&comments).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return comments, nil
}

func (r *PostgresCommentRepository) FindByID(ctx context.Context, commentID uuid.UUID) (*entity.Comments, error) {
	var comment entity.Comments
	tx := r.DB.Begin()
	if err := tx.Find(&comment, "comment_id = ?", commentID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &comment, nil
}

func (r *PostgresCommentRepository) Create(ctx context.Context, comments *entity.Comments, title, username string) error {
	tx := r.DB.Begin()

	var post entity.Post
	if err := tx.Find(&post, "title = ?", title).Error; err != nil {
		tx.Rollback()
		return err
	}

	var user entity.User
	if err := tx.Find(&user, "username = ?", username).Error; err != nil {
		tx.Rollback()
		return err
	}

	comments.PostID = post.PostID
	comments.UserID = user.UserID

	if err := tx.Create(comments).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *PostgresCommentRepository) Update(ctx context.Context, comments *entity.Comments, commentID uuid.UUID) error {
	tx := r.DB.Begin()

	if err := tx.Model(&entity.Comments{}).
		Where("comment_id = ?", commentID).
		Updates(map[string]interface{}{
			"content": comments.Content,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *PostgresCommentRepository) Delete(ctx context.Context, commentID uuid.UUID) error {
	tx := r.DB.Begin()
	if err := tx.Delete(&entity.Comments{}, "comment_id = ?", commentID).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
