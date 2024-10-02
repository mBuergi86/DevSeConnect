package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/pkg/security"
)

type User struct {
	UserID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"user_id"`
	Username       string    `gorm:"uniqueIndex;not null" json:"username"`
	Email          string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash   string    `gorm:"column:password_hash;not null" json:"-"` // Geändert zu PasswordHash
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Bio            string    `json:"bio"`
	ProfilePicture string    `json:"profile_picture"`
	IsActive       bool      `gorm:"default:true" json:"is_active"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func NewUsers(username, email, password, firstName, lastName, bio, profilePicture string) (*User, error) {
	hashedPasswort, err := security.Hash(password)
	if err != nil {
		return nil, err
	}

	return &User{
		UserID:         uuid.New(),
		Username:       username,
		Email:          email,
		PasswordHash:   string(hashedPasswort),
		FirstName:      firstName,
		LastName:       lastName,
		Bio:            bio,
		ProfilePicture: profilePicture,
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := security.Hash(password)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	return security.CheckPasswordHash(u.PasswordHash, password)
}
