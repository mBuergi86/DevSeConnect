package entity

import (
	"time"

	"github.com/google/uuid"
)

type Network struct {
	NetworkID uuid.UUID `db:"network_id" json:"network_id"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	UserID2   uuid.UUID `db:"user_id2" json:"user_id2"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func newNetwork(userID, userID2 uuid.UUID) *Network {
	return &Network{
		NetworkID: uuid.New(),
		UserID:    userID,
		UserID2:   userID2,
		CreatedAt: time.Now(),
	}
}
