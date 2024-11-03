package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id uuid.UUID) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) Create(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(ctx context.Context, user *entity.User, id uuid.UUID) error {
	args := m.Called(ctx, user, id)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockUserRepository)
	ctx := context.TODO()
	mockUser := &entity.User{
		UserID:    uuid.New(),
		Email:     "test@test.org",
		FirstName: "FirstTest",
		LastName:  "LastTest",
		CreatedAt: time.Now(),
	}

	mockRepo.On("FindAll", mock.Anything).Return([]*entity.User{mockUser}, nil)

	result, err := mockRepo.FindAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, []*entity.User{mockUser}, result)
	mockRepo.AssertExpectations(t)
}

func TestFindByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	ctx := context.TODO()
	mockUUID := uuid.New()

	mockUser := entity.User{
		UserID:    mockUUID,
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@test.com",
		CreatedAt: time.Now(),
	}

	mockRepo.On("FindByID", mock.Anything, mockUUID).Return(mockUser, nil)

	result, err := mockRepo.FindByID(ctx, mockUUID)

	assert.NoError(t, err)
	assert.Equal(t, mockUser, result)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockUserRepository)
	ctx := context.TODO()
	mockUser := &entity.User{
		UserID:       uuid.New(),
		Username:     "test",
		Email:        "test@test.org",
		PasswordHash: "helloTest",
		FirstName:    "FirstTest",
		LastName:     "LastTest",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	mockRepo.On("Create", mock.Anything, mockUser).Return(nil)

	err := mockRepo.Create(ctx, mockUser)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	mockRepo := new(MockUserRepository)
	ctx := context.TODO()
	mockUser := &entity.User{
		UserID:   uuid.New(),
		Username: "test1",
	}

	mockRepo.On("Update", ctx, mockUser, mockUser.UserID).Return(nil)

	err := mockRepo.Update(ctx, mockUser, mockUser.UserID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	mockRepo := new(MockUserRepository)
	ctx := context.TODO()
	mockUUID := uuid.New()

	mockRepo.On("Delete", mock.Anything, mockUUID).Return(nil)

	err := mockRepo.Delete(ctx, mockUUID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
