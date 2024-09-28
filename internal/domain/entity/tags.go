package entity

import (
	"time"

	"github.com/google/uuid"
)

type Tags struct {
	TagID     uuid.UUID `db:"tag_id" json:"tag_id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func newTags(name string) *Tags {
	return &Tags{
		TagID:     uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
