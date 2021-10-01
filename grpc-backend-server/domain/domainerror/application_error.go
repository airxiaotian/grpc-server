package domainerror

import "errors"

type ApplicationError interface {
	Error() string
	Unwrap() error
	IsNotFound() bool
	IsValidationError() bool
	IsInternalServerError() bool
}

func IsNotFoundError(err error) bool {
	var applicationError ApplicationError
	return errors.As(err, &applicationError) && applicationError.IsNotFound()
}

func IsValidationError(err error) bool {
	var applicationError ApplicationError
	return errors.As(err, &applicationError) && applicationError.IsValidationError()
}

func IsInternalServerError(err error) bool {
	var applicationError ApplicationError
	return errors.As(err, &applicationError) && applicationError.IsInternalServerError()
}

func IsApplicationError(err error) bool {
	var applicationError ApplicationError
	return errors.As(err, &applicationError)
}
