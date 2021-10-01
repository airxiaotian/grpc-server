package logging

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLoggerToContext(t *testing.T) {
	t.Run("set and get", func(t *testing.T) {
		logger := NewLogger()
		ctx1 := context.Background()
		ctx2 := SetLoggerToContext(ctx1, logger)
		assert.Equal(t, GetLoggerFromContext(ctx2), logger)
	})
	t.Run("default logger", func(t *testing.T) {
		ctx := context.Background()
		assert.NotNil(t, GetLoggerFromContext(ctx))
	})
}
