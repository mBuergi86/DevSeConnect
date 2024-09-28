package service_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	service "github.com/mBuergi86/devseconnect/internal/application"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_FindAll(t *testing.T) {
	// Mocking the repository
	mockRepo := new(mocks.UsersRepository)
	mockUsers := []entity.User{
		{UserID: uuid.New(), Username: "user1", Email: "user1@example.com"},
		{UserID: uuid.New(), Username: "user2", Email: "user2@example.com"},
	}

	// Set up the mock behavior
	mockRepo.On("FindAll", mock.Anything).Return(mockUsers, nil)

	// Create the service with the mock repository
	userService := service.NewUserService(mockRepo)

	// Run the test
	result, err := userService.FindAll(context.Background())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, mockUsers, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_FindByID_Success(t *testing.T) {
	// Mocking the repository
	mockRepo := new(mocks.UsersRepository)
	userID := uuid.New()
	mockUser := entity.User{UserID: userID, Username: "user1", Email: "user1@example.com"}

	// Set up the mock behavior
	mockRepo.On("FindById", mock.Anything, userID).Return(&mockUser, nil)

	// Create the service with the mock repository
	userService := service.NewUserService(mockRepo)

	// Run the test
	result, err := userService.FindByID(context.Background(), userID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, &mockUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_FindByID_UserNotFound(t *testing.T) {
	// Mocking the repository
	mockRepo := new(mocks.UsersRepository)
	userID := uuid.New()

	// Set up the mock behavior
	mockRepo.On("FindById", mock.Anything, userID).Return(nil, errors.New("not found"))

	// Create the service with the mock repository
	userService := service.NewUserService(mockRepo)

	// Run the test
	result, err := userService.FindByID(context.Background(), userID)

	// Assert
	assert.Nil(t, result)
	assert.EqualError(t, err, "User with id "+userID.String()+" not found: not found")
	mockRepo.AssertExpectations(t)
}

func TestUserService_Create(t *testing.T) {
	// Mocking the repository
	mockRepo := new(mocks.UsersRepository)
	mockUser := entity.User{
		UserID:    uuid.New(),
		Username:  "user1",
		Email:     "user1@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "password",
	}

	// Set up the mock behavior
	mockRepo.On("FindByEmail", mock.Anything, mockUser.Email).Return(nil, nil) // No user found with the email
	mockRepo.On("Save", mock.Anything, &mockUser).Return(&mockUser, nil)

	// Create the service with the mock repository
	userService := service.NewUserService(mockRepo)

	// Run the test
	result, err := userService.Create(context.Background(), mockUser)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, &mockUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Create_EmailExists(t *testing.T) {
	// Mocking the repository
	mockRepo := new(mocks.UsersRepository)
	existingUser := entity.User{
		UserID:    uuid.New(),
		Username:  "user1",
		Email:     "user1@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "password",
	}
	newUser := entity.User{
		UserID:    uuid.New(),
		Username:  "user2",
		Email:     "user1@example.com",
		FirstName: "Jane",
		LastName:  "Doe",
		Password:  "password",
	}

	// Set up the mock behavior
	mockRepo.On("FindByEmail", mock.Anything, newUser.Email).Return(&existingUser, nil) // Email already exists

	// Create the service with the mock repository
	userService := service.NewUserService(mockRepo)

	// Run the test
	result, err := userService.Create(context.Background(), newUser)

	// Assert
	assert.Nil(t, result)
	assert.EqualError(t, err, fmt.Sprintf("User with email %s already exists", newUser.Email))
	mockRepo.AssertExpectations(t)
}
