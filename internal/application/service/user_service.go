package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/rabbitmq/amqp091-go"
)

type UserService struct {
	userRepo     repository.UserRepository
	rabbitMQChan *amqp091.Channel
}

func NewUserService(userRepo repository.UserRepository, rabbitMQChan *amqp091.Channel) (*UserService, error) {
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

	return &UserService{userRepo: userRepo, rabbitMQChan: rabbitMQChan}, nil
}

func (s *UserService) publishUserEvent(eventType string, user *entity.User) error {
	event := map[string]interface{}{
		"type": eventType,
		"user": user,
	}
	eventJSON, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("Failed to marshal event: %v\n", err)
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

func (s *UserService) Register(username, email, password, firstName, lastName, bio, profilePicture string) (*entity.User, error) {
	existingUser, _ := s.userRepo.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}
	existingUser, _ = s.userRepo.FindByUsername(username)
	if existingUser != nil {
		return nil, errors.New("user with this username already exists")
	}
	user, err := entity.NewUsers(username, email, password, firstName, lastName, bio, profilePicture)
	if err != nil {
		return nil, err
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	if err := s.publishUserEvent("user_registered", user); err != nil {
		// Log the error, but don't fail the registration
		// TODO: Implement proper error logging
		println("Failed to publish user registered event:", err)
	}
	return user, nil
}

func (s *UserService) Login(email, password string) (*entity.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}
	if err := s.publishUserEvent("user_logged_in", user); err != nil {
		// Log the error, but don't fail the login
		// TODO: Implement proper error logging
		println("Failed to publish user logged in event:", err)
	}
	return user, nil
}

func (s *UserService) GetUserByID(id uuid.UUID) (*entity.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) UpdateUser(updateData map[string]interface{}) (*entity.User, error) {
	userID, ok := updateData["user_id"].(uuid.UUID)
	if !ok {
		return nil, errors.New("invalid user ID")
	}

	existingUser, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	// Aktualisieren Sie nur die Felder, die im updateData vorhanden sind
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

	if err := s.userRepo.Update(existingUser); err != nil {
		return nil, err
	}

	if err := s.publishUserEvent("user_updated", existingUser); err != nil {
		// Log the error, but don't fail the update
		// TODO: Implement
		log.Printf("Failed to publish user updated event: %v\n", err)
	}

	return existingUser, nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	if err := s.userRepo.Delete(id); err != nil {
		return err
	}
	if err := s.publishUserEvent("user_deleted", user); err != nil {
		// Log the error, but don't fail the deletion
		// TODO: Implement proper error logging
		println("Failed to publish user deleted event:", err)
	}
	return nil
}

func (s *UserService) GetUsers() ([]*entity.User, error) {
	return s.userRepo.FindAll()
}
