package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/mBuergi86/devseconnect/pkg/security"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type UserService struct {
	userRepo     repository.UserRepository
	rabbitMQChan *amqp091.Channel
	jwtSecret    string
}

type jwtCustomClaims struct {
	UserID   string          `json:"user_id"`
	Username string          `json:"username"`
	ExpireAt jwt.NumericDate `json:"exp"`
	jwt.RegisteredClaims
}

func NewUserService(userRepo repository.UserRepository, rabbitMQChan *amqp091.Channel) (*UserService, error) {
	if userRepo == nil {
		return nil, errors.New("user repository is required")
	}

	_, err := rabbitMQChan.QueueDeclare(
		"user_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, err
	}

	return &UserService{userRepo: userRepo, rabbitMQChan: rabbitMQChan, jwtSecret: os.Getenv("JWT_SECRET")}, nil
}

func (s *UserService) publishUserEvent(eventType string, user *entity.User) error {
	event := map[string]interface{}{
		"type": eventType,
		"user": user,
	}
	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %v\n", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = s.rabbitMQChan.PublishWithContext(ctx,
		"user_events", // exchange
		"user_queue",  // routing key
		false,         // mandatory
		false,         // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		})
	if err != nil {
		log.Printf("Failed to publish event: %v\n", err)
		return err
	}

	log.Printf("Published %s event for user %s\n", eventType, user.UserID)
	return nil
}

func (s *UserService) Register(ctx context.Context, username, email, password, firstName, lastName, bio, profilePicture string) (*entity.User, error) {
	existingUser, _ := s.userRepo.FindByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}
	existingUser, _ = s.userRepo.FindByUsername(ctx, username)
	if existingUser != nil {
		return nil, errors.New("user with this username already exists")
	}
	user, err := entity.NewUsers(username, email, password, firstName, lastName, bio, profilePicture)
	if err != nil {
		return nil, err
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	if err := s.publishUserEvent("user_registered", user); err != nil {
		// Log the error, but don't fail the registration
		// TODO: Implement proper error logging
		println("Failed to publish user registered event:", err)
	}
	return user, nil
}

func (s *UserService) Login(ctx context.Context, username, password string) (*entity.User, string, error) {
	logger := zerolog.New(os.Stdout)
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
		ExpireAt: *jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "devseconnect",
			Subject:   "user_token",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	logger.Log().Msgf("Claims: %+v\n", claims)
	logger.Log().Msgf("JWT Secret: %s\n", s.jwtSecret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	logger.Log().Msgf("Token: %+v\n", token)

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, "", errors.New("failed to sign token")
	}

	logger.Log().Msgf("TokenString: %s\n", tokenString)

	if err := s.publishUserEvent("user_logged_in", user); err != nil {
		log.Printf("Failed to publish user logged in event: %v", err)
	}

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
		return nil, err
	}

	if err := s.publishUserEvent("user_updated", existingUser); err != nil {
		// Log the error, but don't fail the update
		// TODO: Implement
		log.Printf("Failed to publish user updated event: %v\n", err)
	}

	return existingUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if err := s.userRepo.Delete(ctx, id); err != nil {
		return err
	}
	if err := s.publishUserEvent("user_deleted", user); err != nil {
		// Log the error, but don't fail the deletion
		// TODO: Implement proper error logging
		println("Failed to publish user deleted event:", err)
	}
	return nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return s.userRepo.FindAll(ctx)
}
