package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comments struct {
	CommentID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"comment_id"`
	PostID    uuid.UUID `gorm:"type:uuid;not null" json:"post_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Post *Post `gorm:"foreignKey:PostID;references:PostID" json:"post"`
	User *User `gorm:"foreignKey:UserID;references:UserID" json:"user"`
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
