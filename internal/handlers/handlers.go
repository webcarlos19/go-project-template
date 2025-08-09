package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go-project-template/internal/config"
	"go-project-template/internal/models"
	"go-project-template/internal/service"
	"go-project-template/pkg/database"
	"go-project-template/pkg/logger"
	"go.uber.org/zap"
)

// Handler holds the dependencies for HTTP handlers
type Handler struct {
	userService service.UserService
	db          *database.DB
	config      *config.Config
}

// New creates a new handler instance
func New(userService service.UserService, db *database.DB, config *config.Config) *Handler {
	return &Handler{
		userService: userService,
		db:          db,
		config:      config,
	}
}

// HealthCheck handles health check requests
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Health check requested")

	response := models.HealthResponse{
		Status:    "OK",
		Timestamp: time.Now(),
		Version:   h.config.App.Version,
		Service:   h.config.App.Name,
	}

	// Check database connection
	if h.db != nil {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		if err := h.db.Ping(ctx); err != nil {
			response.Database = "ERROR"
			logger.Error("Database health check failed", zap.Error(err))
		} else {
			response.Database = "OK"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("Failed to encode health check response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// GetUsers handles GET /users requests
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Get users requested")

	// For demonstration, we'll use mock data
	// In a real application, you would parse pagination parameters and use the database
	users, err := h.userService.GetMockUsers(r.Context())
	if err != nil {
		logger.Error("Failed to get users", zap.Error(err))
		h.writeErrorResponse(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		logger.Error("Failed to encode users response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// GetUserByID handles GET /users/{id} requests
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Warn("Invalid user ID provided", zap.String("id", idStr))
		h.writeErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	logger.Debug("Get user by ID requested", zap.Int("user_id", id))

	user, err := h.userService.GetUserByID(r.Context(), id)
	if err != nil {
		logger.Error("Failed to get user by ID", zap.Error(err), zap.Int("user_id", id))
		h.writeErrorResponse(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		logger.Error("Failed to encode user response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// CreateUser handles POST /users requests
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Create user requested")

	var req models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Warn("Invalid JSON in create user request", zap.Error(err))
		h.writeErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Name == "" || req.Email == "" {
		h.writeErrorResponse(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(r.Context(), &req)
	if err != nil {
		logger.Error("Failed to create user", zap.Error(err))
		h.writeErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		logger.Error("Failed to encode create user response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// UpdateUser handles PUT /users/{id} requests
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Warn("Invalid user ID provided", zap.String("id", idStr))
		h.writeErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	logger.Debug("Update user requested", zap.Int("user_id", id))

	var req models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Warn("Invalid JSON in update user request", zap.Error(err))
		h.writeErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Name == "" || req.Email == "" {
		h.writeErrorResponse(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	user, err := h.userService.UpdateUser(r.Context(), id, &req)
	if err != nil {
		logger.Error("Failed to update user", zap.Error(err), zap.Int("user_id", id))
		h.writeErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		logger.Error("Failed to encode update user response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// DeleteUser handles DELETE /users/{id} requests
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Warn("Invalid user ID provided", zap.String("id", idStr))
		h.writeErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	logger.Debug("Delete user requested", zap.Int("user_id", id))

	if err := h.userService.DeleteUser(r.Context(), id); err != nil {
		logger.Error("Failed to delete user", zap.Error(err), zap.Int("user_id", id))
		h.writeErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// writeErrorResponse writes an error response in JSON format
func (h *Handler) writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := models.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
		Code:    statusCode,
	}

	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		logger.Error("Failed to encode error response", zap.Error(err))
	}
}
