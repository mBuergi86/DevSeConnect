package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/models"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/mBuergi86/devseconnect/internal/infrastructure/messaging"
	"github.com/mBuergi86/devseconnect/pkg/security"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type UserService struct {
	userRepo     repository.UserRepository
	rabbitMQChan *amqp091.Channel
	producer     *messaging.Producer
	jwtSecret    string
	logger       zerolog.Logger
}

type jwtCustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewUserService(userRepo repository.UserRepository, rabbitMQChan *amqp091.Channel, producer *messaging.Producer) (*UserService, error) {
	return &UserService{
			userRepo:     userRepo,
			rabbitMQChan: rabbitMQChan,
			producer:     producer,
			jwtSecret:    os.Getenv("JWT_SECRET"),
			logger:       zerolog.New(os.Stderr).With().Timestamp().Str("component", "user_service").Logger(),
		},
		nil
}

func (s *UserService) publishUserEvent(ctx context.Context, eventType string, user *entity.User) error {
	userData, err := json.Marshal(user)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to marshal user data")
		return fmt.Errorf("failed to marshal user data: %w", err)
	}

	message := models.EventMessage{
		Action: eventType,
		Data:   json.RawMessage(userData),
	}

	if err = s.producer.PublishMessage("user_events", "user_queue", message); err != nil {
		s.logger.Error().
			Err(err).
			Str("event_type", eventType).
			Msgf("failed to publish %s event: %v", eventType, err)
		return fmt.Errorf("failed to publish %s event: %w", eventType, err)
	}

	s.logger.Info().
		Str("event_type", eventType).
		Msg("successfully published user event")
	return nil
}

func (s *UserService) Register(ctx context.Context, user *entity.User) error {
	hashedPassword, err := security.Hash(user.PasswordHash)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to hash password")
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user.PasswordHash = hashedPassword

	if err := s.userRepo.Create(ctx, user); err != nil {
		if strings.Contains(err.Error(), "duplicate key error") {
			s.logger.Error().Err(err).Msg("Username or email already exists")
			return errors.New("username or email already exists")
		}
		s.logger.Error().Err(err).Msg("Failed to create user")
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.publishUserEvent(ctx, "user_registered", user); err != nil {
		s.logger.Error().Err(err).Msg("Failed to publish user registered event")
	}

	s.logger.Info().Str("user_id", user.UserID.String()).Msg("User registered successfully")
	return nil
}

func (s *UserService) Login(ctx context.Context, username, password string) (*entity.User, string, error) {
	user, err := s.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, "", err
	}

	if ok := security.CheckPasswordHash(user.PasswordHash, password); !ok {
		log.Println("Invalid credentials")
		return nil, "", errors.New("invalid credentials")
	}

	claims := &jwtCustomClaims{
		UserID:   user.UserID.String(),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "devseconnect",
			Subject:   "user",
			ID:        user.UserID.String(),
			Audience:  []string{user.Username},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, "", errors.New("failed to sign token")
	}

	s.logger.Info().
		Str("user_id", user.UserID.String()).
		Msgf("Expired at: %v", claims.ExpiresAt.Time)

	return user, tokenString, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return s.userRepo.FindByUsername(ctx, username)
}

func (s *UserService) UpdateUser(ctx context.Context, updateData map[string]interface{}) (*entity.User, error) {
	userID, ok := updateData["user_id"].(uuid.UUID)
	if !ok {
		return nil, errors.New("invalid user ID")
	}

	existingUser, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to find user by ID")
		return nil, err
	}

	// Update the user fields if they are present
	if username, ok := updateData["username"].(string); ok {
		existingUser.Username = username
	}
	if email, ok := updateData["email"].(string); ok {
		existingUser.Email = email
	}
	if firstName, ok := updateData["first_name"].(string); ok {
		existingUser.FirstName = firstName
	}
	if lastName, ok := updateData["last_name"].(string); ok {
		existingUser.LastName = lastName
	}
	if bio, ok := updateData["bio"].(string); ok {
		existingUser.Bio = bio
	}
	if profilePicture, ok := updateData["profile_picture"].(string); ok {
		existingUser.ProfilePicture = profilePicture
	}
	if isActive, ok := updateData["is_active"].(bool); ok {
		existingUser.IsActive = isActive
	}

	if err := s.userRepo.Update(ctx, existingUser); err != nil {
		s.logger.Error().Err(err).Msg("Failed to update user")
		return nil, err
	}

	s.logger.Info().Str("user_id", userID.String()).Msg("User updated successfully")

	if err := s.publishUserEvent(ctx, "user_updated", existingUser); err != nil {
		s.logger.Error().Err(err).Msg("Failed to publish user updated event")
	}

	return existingUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to find user by ID")
		return err
	}

	if err := s.userRepo.Delete(ctx, id); err != nil {
		s.logger.Error().Err(err).Msg("Failed to delete user")
		return err
	}

	s.logger.Info().Str("user_id", id.String()).Msg("User deleted successfully")

	if err := s.publishUserEvent(ctx, "user_deleted", user); err != nil {
		s.logger.Error().Err(err).Msg("Failed to publish user deleted event")
	}

	return nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return s.userRepo.FindAll(ctx)
}
