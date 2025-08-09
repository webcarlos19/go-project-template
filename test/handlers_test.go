package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-project-template/internal/config"
	"go-project-template/internal/handlers"
	"go-project-template/internal/models"
	"go-project-template/pkg/logger"
)

// MockUserService is a mock implementation of UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, req *models.UserRequest) (*models.UserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

func (m *MockUserService) GetUserByID(ctx context.Context, id int) (*models.UserResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

func (m *MockUserService) GetUserByEmail(ctx context.Context, email string) (*models.UserResponse, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

func (m *MockUserService) GetUsers(ctx context.Context, pagination *models.Pagination) (*models.PaginatedResponse[models.UserResponse], error) {
	args := m.Called(ctx, pagination)
	return args.Get(0).(*models.PaginatedResponse[models.UserResponse]), args.Error(1)
}

func (m *MockUserService) UpdateUser(ctx context.Context, id int, req *models.UserRequest) (*models.UserResponse, error) {
	args := m.Called(ctx, id, req)
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserService) GetMockUsers(ctx context.Context) ([]*models.UserResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*models.UserResponse), args.Error(1)
}

func setupTest() (*handlers.Handler, *MockUserService) {
	// Initialize logger for tests
	log, _ := logger.NewDevelopment()
	logger.SetGlobalLogger(log)

	mockUserService := &MockUserService{}
	cfg := &config.Config{
		App: config.AppConfig{
			Name:    "test-app",
			Version: "1.0.0",
		},
	}
	handler := handlers.New(mockUserService, nil, cfg)
	return handler, mockUserService
}

func TestHealthCheck(t *testing.T) {
	handler, _ := setupTest()

	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.HealthCheck(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.HealthResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "OK", response.Status)
	assert.Equal(t, "test-app", response.Service)
}

func TestGetUsers(t *testing.T) {
	handler, mockService := setupTest()

	mockUsers := []*models.UserResponse{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	}

	mockService.On("GetMockUsers", mock.Anything).Return(mockUsers, nil)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetUsers(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []*models.UserResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, "John Doe", response[0].Name)

	mockService.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	handler, mockService := setupTest()

	userRequest := &models.UserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	expectedResponse := &models.UserResponse{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	mockService.On("CreateUser", mock.Anything, userRequest).Return(expectedResponse, nil)

	jsonBody, _ := json.Marshal(userRequest)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.CreateUser(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response models.UserResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", response.Name)
	assert.Equal(t, "john@example.com", response.Email)

	mockService.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	handler, mockService := setupTest()

	expectedUser := &models.UserResponse{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	mockService.On("GetUserByID", mock.Anything, 1).Return(expectedUser, nil)

	req, err := http.NewRequest("GET", "/users/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", handler.GetUserByID).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.UserResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", response.Name)

	mockService.AssertExpectations(t)
}
