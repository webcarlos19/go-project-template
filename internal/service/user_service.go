package service

import (
	"context"
	"fmt"
	"time"

	"go-project-template/internal/models"
	"go-project-template/internal/repository"
	"go-project-template/pkg/logger"
	"go.uber.org/zap"
)

// UserService defines the interface for user service operations
type UserService interface {
	CreateUser(ctx context.Context, req *models.UserRequest) (*models.UserResponse, error)
	GetUserByID(ctx context.Context, id int) (*models.UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*models.UserResponse, error)
	GetUsers(ctx context.Context, pagination *models.Pagination) (*models.PaginatedResponse[models.UserResponse], error)
	UpdateUser(ctx context.Context, id int, req *models.UserRequest) (*models.UserResponse, error)
	DeleteUser(ctx context.Context, id int) error
	GetMockUsers(ctx context.Context) ([]*models.UserResponse, error)
}

// userService implements UserService interface
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (s *userService) CreateUser(ctx context.Context, req *models.UserRequest) (*models.UserResponse, error) {
	logger.Info("Creating new user", zap.String("email", req.Email))

	// Check if user already exists
	existingUser, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	response := user.ToResponse()
	return &response, nil
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(ctx context.Context, id int) (*models.UserResponse, error) {
	logger.Debug("Getting user by ID", zap.Int("user_id", id))

	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	response := user.ToResponse()
	return &response, nil
}

// GetUserByEmail retrieves a user by email
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.UserResponse, error) {
	logger.Debug("Getting user by email", zap.String("email", email))

	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	response := user.ToResponse()
	return &response, nil
}

// GetUsers retrieves all users with pagination
func (s *userService) GetUsers(ctx context.Context, pagination *models.Pagination) (*models.PaginatedResponse[models.UserResponse], error) {
	logger.Debug("Getting users", zap.Int("page", pagination.Page), zap.Int("per_page", pagination.PerPage))

	users, total, err := s.userRepo.GetAll(ctx, pagination)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	// Convert to response format
	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToResponse()
	}

	totalPages := (total + pagination.PerPage - 1) / pagination.PerPage

	return &models.PaginatedResponse[models.UserResponse]{
		Data:       userResponses,
		Total:      total,
		Page:       pagination.Page,
		PerPage:    pagination.PerPage,
		TotalPages: totalPages,
	}, nil
}

// UpdateUser updates a user
func (s *userService) UpdateUser(ctx context.Context, id int, req *models.UserRequest) (*models.UserResponse, error) {
	logger.Info("Updating user", zap.Int("user_id", id))

	// Check if user exists
	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Check if email is being changed and if it already exists
	if existingUser.Email != req.Email {
		emailUser, _ := s.userRepo.GetByEmail(ctx, req.Email)
		if emailUser != nil && emailUser.ID != id {
			return nil, fmt.Errorf("user with email %s already exists", req.Email)
		}
	}

	// Update user fields
	existingUser.Name = req.Name
	existingUser.Email = req.Email

	if err := s.userRepo.Update(ctx, existingUser); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	response := existingUser.ToResponse()
	return &response, nil
}

// DeleteUser deletes a user
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	logger.Info("Deleting user", zap.Int("user_id", id))

	if err := s.userRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// GetMockUsers returns mock users for demonstration
func (s *userService) GetMockUsers(ctx context.Context) ([]*models.UserResponse, error) {
	logger.Debug("Getting mock users")

	mockUsers := []*models.UserResponse{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: time.Now().Add(-24 * time.Hour),
			UpdatedAt: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:        2,
			Name:      "Jane Smith",
			Email:     "jane.smith@example.com",
			CreatedAt: time.Now().Add(-48 * time.Hour),
			UpdatedAt: time.Now().Add(-48 * time.Hour),
		},
		{
			ID:        3,
			Name:      "Bob Johnson",
			Email:     "bob.johnson@example.com",
			CreatedAt: time.Now().Add(-72 * time.Hour),
			UpdatedAt: time.Now().Add(-72 * time.Hour),
		},
	}

	return mockUsers, nil
}
