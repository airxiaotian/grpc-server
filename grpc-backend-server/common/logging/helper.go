package logging

import (
	"context"
	"go.uber.org/zap"
)

type key int

const (
	contextLoggerKey = key(1)
)

func GetLoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(contextLoggerKey).(*zap.SugaredLogger); ok {
		return logger
	}
	return NewLogger()
}

func SetLoggerToContext(parent context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(parent, contextLoggerKey, logger)
}
