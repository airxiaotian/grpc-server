package domainerror

import "github.com/pkg/errors"

type NotFoundError struct {
	message string
	cause   error
}

func NewNotFoundError(message string, cause error) error {
	return errors.WithStack(&NotFoundError{message, cause})
}

func (n *NotFoundError) Error() string {
	return n.message
}

func (n *NotFoundError) Unwrap() error {
	return n.cause
}

func (n *NotFoundError) IsNotFound() bool {
	return true
}

func (n *NotFoundError) IsValidationError() bool {
	return false
}

func (n *NotFoundError) IsInternalServerError() bool {
	return false
}
