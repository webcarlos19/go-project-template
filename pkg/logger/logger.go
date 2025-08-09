package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps zap logger
type Logger struct {
	*zap.Logger
}

// New creates a new logger instance
func New(level, format, output string) (*Logger, error) {
	// Parse log level
	logLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		logLevel = zapcore.InfoLevel
	}

	// Create encoder config
	var encoderConfig zapcore.EncoderConfig
	if format == "json" {
		encoderConfig = zap.NewProductionEncoderConfig()
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Create encoder
	var encoder zapcore.Encoder
	if format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Create writer syncer
	var writeSyncer zapcore.WriteSyncer
	switch output {
	case "stdout":
		writeSyncer = zapcore.AddSync(os.Stdout)
	case "stderr":
		writeSyncer = zapcore.AddSync(os.Stderr)
	default:
		file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			writeSyncer = zapcore.AddSync(os.Stdout)
		} else {
			writeSyncer = zapcore.AddSync(file)
		}
	}

	// Create core
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	// Create logger
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{zapLogger}, nil
}

// NewDevelopment creates a development logger
func NewDevelopment() (*Logger, error) {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &Logger{zapLogger}, nil
}

// NewProduction creates a production logger
func NewProduction() (*Logger, error) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &Logger{zapLogger}, nil
}

// WithField adds a single field to the logger
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{l.Logger.With(zap.Any(key, value))}
}

// WithFields adds multiple fields to the logger
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}
	return &Logger{l.Logger.With(zapFields...)}
}

// WithError adds an error field to the logger
func (l *Logger) WithError(err error) *Logger {
	return &Logger{l.Logger.With(zap.Error(err))}
}

// Global logger instance
var globalLogger *Logger

// SetGlobalLogger sets the global logger instance
func SetGlobalLogger(logger *Logger) {
	globalLogger = logger
}

// GetGlobalLogger returns the global logger instance
func GetGlobalLogger() *Logger {
	if globalLogger == nil {
		// Create a default logger if none is set
		logger, _ := NewDevelopment()
		globalLogger = logger
	}
	return globalLogger
}

// Info logs at info level using the global logger
func Info(msg string, fields ...zap.Field) {
	GetGlobalLogger().Info(msg, fields...)
}

// Error logs at error level using the global logger
func Error(msg string, fields ...zap.Field) {
	GetGlobalLogger().Error(msg, fields...)
}

// Debug logs at debug level using the global logger
func Debug(msg string, fields ...zap.Field) {
	GetGlobalLogger().Debug(msg, fields...)
}

// Warn logs at warn level using the global logger
func Warn(msg string, fields ...zap.Field) {
	GetGlobalLogger().Warn(msg, fields...)
}

// Fatal logs at fatal level using the global logger and exits
func Fatal(msg string, fields ...zap.Field) {
	GetGlobalLogger().Fatal(msg, fields...)
}
