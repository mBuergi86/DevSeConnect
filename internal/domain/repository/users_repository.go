package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
)

type UsersRepository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Save(ctx context.Context, saved *entity.User) (*entity.User, error)
	Update(ctx context.Context, updated *entity.User) (*entity.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
