package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/pkg/security"
)

type User struct {
	UserID   uuid.UUID `db:"user_id" json:"user_id"`
	Username string    `db:"username" json:"username"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password_hash" json:"-"`

	FirstName      string    `db:"first_name" json:"first_name"`
	LastName       string    `db:"last_name" json:"last_name"`
	Bio            string    `db:"bio" json:"bio"`
	ProfilePicture string    `db:"profile_picture" json:"profile_picture"`
	IsActive       bool      `db:"is_active" json:"is_active"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

func (u *User) GenerateFullName() string {
	return u.FirstName + " " + u.LastName
}

func NewUsers(username, email, password, firstName, lastName, bio, profilePicture string) (*User, error) {
	if username == "" || email == "" || password == "" {
		return nil, errors.New("username, email, and password are required")
	}

	passwordHash, err := security.Hash(password)
	if err != nil {
		return nil, err
	}

	return &User{
		UserID:         uuid.New(),
		Username:       username,
		Email:          email,
		Password:       string(passwordHash),
		FirstName:      firstName,
		LastName:       lastName,
		Bio:            bio,
		ProfilePicture: profilePicture,
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}
