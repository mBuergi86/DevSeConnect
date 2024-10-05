package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindAll(ctx context.Context) ([]*entity.User, error)
}

type PostgresUserRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	if db == nil || redis == nil {
		log.Fatal("Database or Redis is not initialized")
	}

	return &PostgresUserRepository{DB: db, Redis: redis}
}

func (r *PostgresUserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User

	// Try to get from cache
	cachedUser, err := r.getUserFromCache("id", id.String())
	if err == nil && cachedUser != nil {
		return cachedUser, nil
	}

	// If not in cache, get from database
	if err := r.DB.First(&user, "user_id = ?", id).Error; err != nil {
		return nil, err
	}

	// Cache the user for future requests
	r.cacheUser(&user)

	return &user, nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	// Try to get from cache
	cachedUser, err := r.getUserFromCache("email", email)
	if err == nil && cachedUser != nil {
		return cachedUser, nil
	}

	// If not in cache, get from database
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// Cache the user for future requests
	r.cacheUser(&user)

	return &user, nil
}

func (r *PostgresUserRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User

	// Try to get from cache
	cachedUser, err := r.getUserFromCache("username", username)
	if err == nil && cachedUser != nil {
		return cachedUser, nil
	}

	// If not in cache, get from database
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	// Cache the user for future requests
	r.cacheUser(&user)

	return &user, nil
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *entity.User) error {
	tx := r.DB.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return r.cacheUser(user)
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *entity.User) error {
	tx := r.DB.Begin()
	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return r.cacheUser(user)
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tx := r.DB.Begin()
	if err := tx.Delete(&entity.User{}, "user_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return r.removeUserFromCache(id)
}

func (r *PostgresUserRepository) cacheUser(user *entity.User) error {
	ctx := context.Background()
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	pipe := r.Redis.Pipeline()
	pipe.Set(ctx, fmt.Sprintf("user:id:%s", user.UserID), userJSON, time.Hour)
	pipe.Set(ctx, fmt.Sprintf("user:email:%s", user.Email), userJSON, time.Hour)
	pipe.Set(ctx, fmt.Sprintf("user:username:%s", user.Username), userJSON, time.Hour)

	_, err = pipe.Exec(ctx)
	return err
}

func (r *PostgresUserRepository) getUserFromCache(field, value string) (*entity.User, error) {
	ctx := context.Background()
	userJSON, err := r.Redis.Get(ctx, fmt.Sprintf("user:%s:%s", field, value)).Bytes()
	if err == redis.Nil {
		// Cache miss is not an error, return nil to indicate not found in cache
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user entity.User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) removeUserFromCache(id uuid.UUID) error {
	ctx := context.Background()
	user, err := r.FindByID(ctx, id)
	if err != nil {
		return err
	}

	pipe := r.Redis.Pipeline()
	pipe.Del(ctx, fmt.Sprintf("user:id:%s", user.UserID))
	pipe.Del(ctx, fmt.Sprintf("user:email:%s", user.Email))
	pipe.Del(ctx, fmt.Sprintf("user:username:%s", user.Username))

	_, err = pipe.Exec(ctx)
	return err
}
