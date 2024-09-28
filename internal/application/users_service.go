package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
	"github.com/mBuergi86/devseconnect/pkg/response"
)

type UserService struct {
	repo repository.UsersRepository
}

func NewUserService(repo repository.UsersRepository) *UserService {
	return &UserService{
		repo,
	}
}

func (s *UserService) FindAll(ctx context.Context) ([]entity.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, response.NewError("Failed to fetch users", 500, err)
	}
	return users, nil
}

func (s *UserService) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, response.NewError(fmt.Sprintf("User with id %s not found:", id.String()), 404, err)
	}

	return user, nil
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, response.NewError(fmt.Sprintf("User with email %s not found", email), 404, err)
	}

	return user, nil
}

func (s *UserService) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	// Validate required fields
	if user.Username == "" {
		return nil, response.NewError("Username is required", 400, nil)
	} else if user.Email == "" {
		return nil, response.NewError("Email is required", 400, nil)
	} else if user.Password == "" {
		return nil, response.NewError("Password is required", 400, nil)
	} else if user.FirstName == "" {
		return nil, response.NewError("First name is required", 400, nil)
	} else if user.LastName == "" {
		return nil, response.NewError("Last name is required", 400, nil)
	}

	// Check if email already exists
	existingUser, err := s.repo.FindByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return nil, response.NewError(fmt.Sprintf("User with email %s already exists", existingUser.Email), 409, nil)
	}

	// Generate full name before saving
	user.GenerateFullName()

	// Save the new user
	newUser, err := s.repo.Save(ctx, &user)
	if err != nil {
		return nil, response.NewError("Failed to create user", 500, err)
	}

	return newUser, nil
}

func (s *UserService) Update(ctx context.Context, user entity.User) (*entity.User, error) {
	existingUser, err := s.repo.FindById(ctx, user.UserID)
	if err != nil {
		return nil, response.NewError("User not found", 404, err)
	}

	// Update the fields of the existing user with new values
	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Bio = user.Bio
	existingUser.ProfilePicture = user.ProfilePicture
	existingUser.IsActive = user.IsActive

	// Regenerate full name if the first name or last name has changed
	existingUser.GenerateFullName()

	updatedUser, err := s.repo.Update(ctx, existingUser)
	if err != nil {
		return nil, response.NewError("Failed to update user", 500, err)
	}
	return updatedUser, nil
}

func (s *UserService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.FindById(ctx, id)
	if err != nil {
		return response.NewError("User not found", 404, err)
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return response.NewError("Failed to delete user", 500, err)
	}

	return nil
}
