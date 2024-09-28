package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	PostID    uuid.UUID `db:"post_id" json:"post_id"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	MediaType string    `db:"media_type" json:"media_type"`
	MediaURL  string    `db:"media_url" json:"media_url"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewPost(userID uuid.UUID, title, content, mediaType, mediaURL string) *Post {
	return &Post{
		PostID:    uuid.New(),
		UserID:    userID,
		Title:     title,
		Content:   content,
		MediaType: mediaType,
		MediaURL:  mediaURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
