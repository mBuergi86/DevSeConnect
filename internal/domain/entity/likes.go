package entity

import (
	"time"

	"github.com/google/uuid"
)

type Likes struct {
	LikeID    uuid.UUID `db:"like_id" json:"like_id"`
	PostID    uuid.UUID `db:"post_id" json:"post_id"`
	CommentID uuid.UUID `db:"comment_id" json:"comment_id"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func newLikes(postID, commentID, userID uuid.UUID) *Likes {
	return &Likes{
		LikeID:    uuid.New(),
		PostID:    postID,
		CommentID: commentID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
}
