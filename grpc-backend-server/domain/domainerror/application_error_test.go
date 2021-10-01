package domainerror

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInternalServerError(t *testing.T) {
	t.Run("standard error", func(t *testing.T) {
		err := errors.New("err")
		assert.False(t, IsInternalServerError(err))
	})
	t.Run("NotFoundError", func(t *testing.T) {
		err := &NotFoundError{}
		assert.False(t, IsInternalServerError(err))
	})
	t.Run("ValidationError", func(t *testing.T) {
		err := &ValidationError{}
		assert.False(t, IsInternalServerError(err))
	})
	t.Run("InternalServerError", func(t *testing.T) {
		err := &InternalServerError{}
		assert.True(t, IsInternalServerError(err))
	})
}

func TestIsNotFoundError(t *testing.T) {
	t.Run("standard error", func(t *testing.T) {
		err := errors.New("err")
		assert.False(t, IsNotFoundError(err))
	})
	t.Run("NotFoundError", func(t *testing.T) {
		err := &NotFoundError{}
		assert.True(t, IsNotFoundError(err))
	})
	t.Run("ValidationError", func(t *testing.T) {
		err := &ValidationError{}
		assert.False(t, IsNotFoundError(err))
	})
	t.Run("InternalServerError", func(t *testing.T) {
		err := &InternalServerError{}
		assert.False(t, IsNotFoundError(err))
	})
}

func TestIsValidationError(t *testing.T) {
	t.Run("standard error", func(t *testing.T) {
		err := errors.New("err")
		assert.False(t, IsValidationError(err))
	})
	t.Run("NotFoundError", func(t *testing.T) {
		err := &NotFoundError{}
		assert.False(t, IsValidationError(err))
	})
	t.Run("ValidationError", func(t *testing.T) {
		err := &ValidationError{}
		assert.True(t, IsValidationError(err))
	})
	t.Run("InternalServerError", func(t *testing.T) {
		err := &InternalServerError{}
		assert.False(t, IsValidationError(err))
	})
}

func TestIsApplicationError(t *testing.T) {
	t.Run("standard error", func(t *testing.T) {
		err := errors.New("err")
		assert.False(t, IsValidationError(err))
	})
	t.Run("NotFoundError", func(t *testing.T) {
		err := &NotFoundError{}
		assert.True(t, IsApplicationError(err))
	})
	t.Run("ValidationError", func(t *testing.T) {
		err := &ValidationError{}
		assert.True(t, IsApplicationError(err))
	})
	t.Run("InternalServerError", func(t *testing.T) {
		err := &InternalServerError{}
		assert.True(t, IsApplicationError(err))
	})
}
