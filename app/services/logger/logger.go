package logger

import (
	"context"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

type loggerKey struct{}

var loggerNop = &Logger{zap.NewNop()}

func (l *Logger) ToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerKey{}, l)
}

func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(loggerKey{}).(*Logger)
	if !ok {
		return loggerNop
	}
	return logger
}
