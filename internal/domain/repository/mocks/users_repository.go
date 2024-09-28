package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

// UsersRepository is a mock implementation of the repository.UsersRepository interface.
type UsersRepository struct {
	mock.Mock
}

// FindAll is a mock method for retrieving all users.
func (m *UsersRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.User), args.Error(1)
}

// FindById is a mock method for finding a user by ID.
func (m *UsersRepository) FindById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

// FindByEmail is a mock method for finding a user by email.
func (m *UsersRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

// Save is a mock method for saving a new user.
func (m *UsersRepository) Save(ctx context.Context, user *entity.User) (*entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*entity.User), args.Error(1)
}

// Update is a mock method for updating an existing user.
func (m *UsersRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*entity.User), args.Error(1)
}

// Delete is a mock method for deleting a user by ID.
func (m *UsersRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
