package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-project-template/api/routes"
	"go-project-template/internal/config"
	"go-project-template/internal/handlers"
	"go-project-template/internal/repository"
	"go-project-template/internal/service"
	"go-project-template/pkg/database"
	"go-project-template/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := loadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	log, err := initLogger(cfg.Logger)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	logger.SetGlobalLogger(log)

	logger.Info("Starting application",
		zap.String("app", cfg.App.Name),
		zap.String("version", cfg.App.Version),
		zap.String("environment", cfg.App.Environment),
	)

	// Initialize database (optional for this template)
	var db *database.DB
	if cfg.App.Environment != "development" {
		db, err = initDatabase(cfg.Database)
		if err != nil {
			logger.Fatal("Failed to initialize database", zap.Error(err))
		}
		defer db.Close()
	}

	// Initialize dependencies
	userRepo := repository.NewUserRepository(nil) // Pass db.Pool if database is available
	userService := service.NewUserService(userRepo)
	handler := handlers.New(userService, db, cfg)

	// Setup routes
	router := routes.SetupRoutes(handler)

	// Create HTTP server
	server := &http.Server{
		Addr:         cfg.Server.GetServerAddress(),
		Handler:      router,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Server starting", zap.String("address", cfg.Server.GetServerAddress()))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	logger.Info("Server started successfully", zap.String("address", cfg.Server.GetServerAddress()))

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Create a deadline for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exited")
}

// loadConfig loads the application configuration
func loadConfig() (*config.Config, error) {
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "development"
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./configs"
	}

	return config.LoadWithEnvironment(configPath, environment)
}

// initLogger initializes the logger based on configuration
func initLogger(cfg config.LoggerConfig) (*logger.Logger, error) {
	return logger.New(cfg.Level, cfg.Format, cfg.Output)
}

// initDatabase initializes the database connection
func initDatabase(cfg config.DatabaseConfig) (*database.DB, error) {
	return database.New(&cfg)
}
