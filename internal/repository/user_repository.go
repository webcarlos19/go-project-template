package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-project-template/internal/models"
	"go-project-template/pkg/logger"
	"go.uber.org/zap"
)

// UserRepository defines the interface for user repository operations
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetAll(ctx context.Context, pagination *models.Pagination) ([]*models.User, int, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
}

// userRepository implements UserRepository interface
type userRepository struct {
	db *pgxpool.Pool
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		logger.Error("Failed to create user", zap.Error(err), zap.String("email", user.Email))
		return fmt.Errorf("failed to create user: %w", err)
	}

	logger.Info("User created successfully", zap.Int("user_id", user.ID), zap.String("email", user.Email))
	return nil
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE id = $1`

	user := &models.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found with id %d", id)
		}
		logger.Error("Failed to get user by ID", zap.Error(err), zap.Int("user_id", id))
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE email = $1`

	user := &models.User{}
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found with email %s", email)
		}
		logger.Error("Failed to get user by email", zap.Error(err), zap.String("email", email))
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// GetAll retrieves all users with pagination
func (r *userRepository) GetAll(ctx context.Context, pagination *models.Pagination) ([]*models.User, int, error) {
	// Get total count
	countQuery := `SELECT COUNT(*) FROM users`
	var total int
	err := r.db.QueryRow(ctx, countQuery).Scan(&total)
	if err != nil {
		logger.Error("Failed to get users count", zap.Error(err))
		return nil, 0, fmt.Errorf("failed to get users count: %w", err)
	}

	// Get users with pagination
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(ctx, query, pagination.PerPage, pagination.Offset)
	if err != nil {
		logger.Error("Failed to get users", zap.Error(err))
		return nil, 0, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			logger.Error("Failed to scan user", zap.Error(err))
			return nil, 0, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		logger.Error("Error iterating users", zap.Error(err))
		return nil, 0, fmt.Errorf("error iterating users: %w", err)
	}

	return users, total, nil
}

// Update updates a user
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, updated_at = $3
		WHERE id = $4`

	user.UpdatedAt = time.Now()

	result, err := r.db.Exec(ctx, query, user.Name, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		logger.Error("Failed to update user", zap.Error(err), zap.Int("user_id", user.ID))
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found with id %d", user.ID)
	}

	logger.Info("User updated successfully", zap.Int("user_id", user.ID))
	return nil
}

// Delete deletes a user
func (r *userRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Failed to delete user", zap.Error(err), zap.Int("user_id", id))
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user not found with id %d", id)
	}

	logger.Info("User deleted successfully", zap.Int("user_id", id))
	return nil
}
