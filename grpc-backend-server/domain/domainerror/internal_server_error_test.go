package domainerror

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	cause := errors.New("err")
	err := NewInternalServerError(cause.Error(), cause)
	assert.NotNil(t, err)
	_, ok1 := err.(*InternalServerError)
	assert.False(t, ok1)
	var applicationError ApplicationError
	errors.As(err, &applicationError)
	assert.NotNil(t, applicationError)
	_, ok2 := applicationError.(*InternalServerError)
	assert.True(t, ok2)
}

func TestInternalServerError_Error(t *testing.T) {
	cause := errors.New("err")
	err := &InternalServerError{
		message: cause.Error(),
		cause:   cause,
	}
	assert.Equal(t, err.Error(), cause.Error())
}

func TestInternalServerError_Unwrap(t *testing.T) {
	cause := errors.New("err")
	err := &InternalServerError{
		message: cause.Error(),
		cause:   cause,
	}
	assert.Equal(t, err.Unwrap(), cause)
}

func TestInternalServerError_IsInternalServerError(t *testing.T) {
	err := &InternalServerError{}
	assert.True(t, err.IsInternalServerError())
}

func TestInternalServerError_IsNotFound(t *testing.T) {
	err := &InternalServerError{}
	assert.False(t, err.IsNotFound())
}

func TestInternalServerError_IsValidationError(t *testing.T) {
	err := &InternalServerError{}
	assert.False(t, err.IsValidationError())
}
