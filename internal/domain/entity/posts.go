package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	PostID    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"post_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	MediaType string    `gorm:"size:50" json:"media_type"`
	MediaURL  string    `gorm:"size:255" json:"media_url"`
	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	// Remove the User relation if it's not needed
	// User      User      `gorm:"foreignKey:UserID" json:"-"`
	User *User `gorm:"foreignKey:UserID;references:UserID" json:"user"`
}

func NewPost(userID uuid.UUID, title, content, mediaType, mediaURL string) *Post {
	return &Post{
		Title:     title,
		Content:   content,
		MediaType: mediaType,
		MediaURL:  mediaURL,
		IsDeleted: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
