package entity

import (
	"time"

	"github.com/google/uuid"
)

type Messages struct {
	MessageID  uuid.UUID `db:"message_id" json:"message_id"`
	SenderID   uuid.UUID `db:"sender_id" json:"sender_id"`
	ReceiverID uuid.UUID `db:"receiver_id" json:"receiver_id"`
	Content    string    `db:"content" json:"content"`
	IsRead     bool      `db:"is_read" json:"is_read"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

func newMessages(senderID, receiverID uuid.UUID, content string, isRead bool) *Messages {
	return &Messages{
		MessageID:  uuid.New(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		IsRead:     isRead,
		CreatedAt:  time.Now(),
	}
}
