package domainerror

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidationError_Error(t *testing.T) {
	builder := NewValidationErrorBuilder()
	builder.AddMessage("key1", "message1")
	builder.AddMessage("key2", "message2")
	err := builder.Build()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "key1")
	assert.Contains(t, err.Error(), "message1")
	assert.Contains(t, err.Error(), "key2")
	assert.Contains(t, err.Error(), "message2")
}

func TestValidationError_Unwrap(t *testing.T) {
	builder := NewValidationErrorBuilder()
	builder.AddMessage("key1", "message1")
	builder.AddMessage("key2", "message2")
	var applicationError ApplicationError
	errors.As(builder.Build(), &applicationError)
	assert.NotNil(t, applicationError)
	err := applicationError.Unwrap()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "key1")
	assert.Contains(t, err.Error(), "message1")
	assert.Contains(t, err.Error(), "key2")
	assert.Contains(t, err.Error(), "message2")
}

func TestValidationError_IsInternalServerError(t *testing.T) {
	err := &ValidationError{}
	assert.False(t, err.IsInternalServerError())
}

func TestValidationError_IsNotFound(t *testing.T) {
	err := &ValidationError{}
	assert.False(t, err.IsNotFound())
}

func TestValidationError_IsValidationError(t *testing.T) {
	err := &ValidationError{}
	assert.True(t, err.IsValidationError())
}

func TestValidationError_GetFields(t *testing.T) {
	builder := NewValidationErrorBuilder()
	builder.AddMessage("key1", "message1")
	builder.AddMessage("key1", "message2")
	builder.AddMessage("key2", "message3")
	var applicationError ApplicationError
	errors.As(builder.Build(), &applicationError)
	assert.NotNil(t, applicationError)
	validationError, ok := applicationError.(*ValidationError)
	assert.True(t, ok)
	fields := validationError.GetFields()
	assert.Len(t, fields["key1"], 2)
	assert.Len(t, fields["key2"], 1)
}
