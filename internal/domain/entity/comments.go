package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comments struct {
	CommentID uuid.UUID `db:"comment_id" json:"comment_id"`
	PostID    uuid.UUID `db:"post_id" json:"post_id"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	Content   string    `db:"content" json:"content"`
	IsDeleted bool      `db:"is_deleted" json:"is_deleted"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewComment(postID, userID uuid.UUID, content string) *Comments {
	return &Comments{
		CommentID: uuid.New(),
		PostID:    postID,
		UserID:    userID,
		Content:   content,
		IsDeleted: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
