package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go-project-template/internal/config"
	"go-project-template/pkg/logger"
	"go.uber.org/zap"
)

// DB wraps the database connection pool
type DB struct {
	Pool *pgxpool.Pool
}

// New creates a new database connection
func New(cfg *config.DatabaseConfig) (*DB, error) {
	// Create connection config
	config, err := pgxpool.ParseConfig(cfg.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	// Set connection pool settings
	config.MaxConns = int32(cfg.MaxOpenConns)
	config.MinConns = int32(cfg.MaxIdleConns)
	config.MaxConnLifetime = cfg.MaxLifetime

	// Create connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("Database connection established successfully",
		zap.String("host", cfg.Host),
		zap.Int("port", cfg.Port),
		zap.String("database", cfg.DBName),
	)

	return &DB{Pool: pool}, nil
}

// Close closes the database connection pool
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
		logger.Info("Database connection closed")
	}
}

// Ping checks if the database connection is alive
func (db *DB) Ping(ctx context.Context) error {
	return db.Pool.Ping(ctx)
}

// Stats returns database connection pool statistics
func (db *DB) Stats() *pgxpool.Stat {
	return db.Pool.Stat()
}

// GetPool returns the underlying connection pool
func (db *DB) GetPool() *pgxpool.Pool {
	return db.Pool
}
