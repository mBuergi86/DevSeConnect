package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserConnections struct {
	FollowerID uuid.UUID `db:"follower_id" json:"follower_id"`
	FollowedID uuid.UUID `db:"followed_id" json:"followed_id"`
	IsActive   bool      `db:"is_active" json:"is_active"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

func newUserConnections(followedID uuid.UUID, isActive bool) *UserConnections {
	return &UserConnections{
		FollowerID: uuid.New(),
		FollowedID: followedID,
		IsActive:   isActive,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
