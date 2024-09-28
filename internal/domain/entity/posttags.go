package entity

import "github.com/google/uuid"

type PostTags struct {
	PostID uuid.UUID `db:"post_id" json:"post_id"`
	TagID  uuid.UUID `db:"tag_id" json:"tag_id"`
}

func New(postID, tagID uuid.UUID) *PostTags {
	return &PostTags{
		PostID: postID,
		TagID:  tagID,
	}
}
