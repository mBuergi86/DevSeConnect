package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/mBuergi86/devseconnect/internal/domain/entity"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("record does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type PostgresUsersRepository struct {
	db *sql.DB
}

func NewPostgresUsersRepository(db *sql.DB) *PostgresUsersRepository {
	return &PostgresUsersRepository{db: db}
}

// Ensure PostgresUsersRepository implements UsersRepository interface
var _ UsersRepository = &PostgresUsersRepository{}

// FindAll fetches all users
func (r *PostgresUsersRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Bio, &user.ProfilePicture, &user.IsActive, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// FindById fetches a user by UUID
func (r *PostgresUsersRepository) FindById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE user_id = $1", id).Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Bio, &user.ProfilePicture, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrNotExist
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail fetches a user by email
func (r *PostgresUsersRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE email = $1", email).Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Bio, &user.ProfilePicture, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrNotExist
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

// Save inserts a new user
func (r *PostgresUsersRepository) Save(ctx context.Context, user *entity.User) (*entity.User, error) {
	_, err := r.db.ExecContext(ctx, `INSERT INTO users (user_id, username, email, password_hash, first_name, last_name, bio, profile_picture, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		user.UserID, user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.Bio, user.ProfilePicture, user.IsActive, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, ErrDuplicate
	}
	return user, nil
}

// Update modifies an existing user
func (r *PostgresUsersRepository) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	_, err := r.db.ExecContext(ctx, `UPDATE users SET username=$2, email=$3, password_hash=$4, first_name=$5, last_name=$6, bio=$7, profile_picture=$8, is_active=$9, updated_at=$10
		WHERE user_id=$1`, user.UserID, user.Username, user.Email, user.Password, user.FirstName, user.LastName, user.Bio, user.ProfilePicture, user.IsActive, user.UpdatedAt)
	if err != nil {
		return nil, ErrUpdateFailed
	}
	return user, nil
}

// Delete removes a user by UUID
func (r *PostgresUsersRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE user_id = $1", id)
	if err != nil {
		return ErrDeleteFailed
	}
	return nil
}
