package domainerror

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNotFoundError(t *testing.T) {
	cause := errors.New("err")
	err := NewNotFoundError(cause.Error(), cause)
	assert.NotNil(t, err)
	_, ok1 := err.(*NotFoundError)
	assert.False(t, ok1)
	var applicationError ApplicationError
	errors.As(err, &applicationError)
	assert.NotNil(t, applicationError)
	_, ok2 := applicationError.(*NotFoundError)
	assert.True(t, ok2)
}

func TestNotFoundError_Error(t *testing.T) {
	cause := errors.New("err")
	err := &NotFoundError{
		message: cause.Error(),
		cause:   cause,
	}
	assert.Equal(t, err.Error(), cause.Error())
}

func TestNotFoundError_Unwrap(t *testing.T) {
	cause := errors.New("err")
	err := &NotFoundError{
		message: cause.Error(),
		cause:   cause,
	}
	assert.Equal(t, err.Unwrap(), cause)
}

func TestNotFoundError_IsInternalServerError(t *testing.T) {
	err := &NotFoundError{}
	assert.False(t, err.IsInternalServerError())
}

func TestNotFoundError_IsNotFound(t *testing.T) {
	err := &NotFoundError{}
	assert.True(t, err.IsNotFound())
}

func TestNotFoundError_IsValidationError(t *testing.T) {
	err := &NotFoundError{}
	assert.False(t, err.IsValidationError())
}
