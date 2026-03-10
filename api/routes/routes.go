package routes

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go-project-template/internal/handlers"
	"go-project-template/pkg/logger"
	"go.uber.org/zap"
)

// SetupRoutes configures all the application routes
func SetupRoutes(handler *handlers.Handler) *mux.Router {
	router := mux.NewRouter()

	// Add middleware
	router.Use(loggingMiddleware)
	router.Use(corsMiddleware)

	// API version 1 routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Health check endpoint
	router.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	// User routes
	api.HandleFunc("/users", handler.GetUsers).Methods("GET")
	api.HandleFunc("/users", handler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id:[0-9]+}", handler.GetUserByID).Methods("GET")
	api.HandleFunc("/users/{id:[0-9]+}", handler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")

	// Root endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message": "Welcome to Go Project Template API", "version": "1.0.0"}`))
	}).Methods("GET")

	return router
}

// loggingMiddleware logs HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a response writer wrapper to capture status code
		wrapper := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Call next handler
		next.ServeHTTP(wrapper, r)

		// Log the request
		duration := time.Since(start)
		logger.Info("HTTP Request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("query", r.URL.RawQuery),
			zap.Int("status", wrapper.statusCode),
			zap.Duration("duration", duration),
			zap.String("user_agent", r.UserAgent()),
			zap.String("remote_addr", r.RemoteAddr),
		)
	})
}

// corsMiddleware adds CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write ensures status code is set
func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.ResponseWriter.Write(b)
}
