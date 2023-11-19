package errors

import (
	"errors"
	"fmt"
)

var (
	ErrBadRequest    = errors.New("bad request")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrInternalError = errors.New("internal error")
)

type ServiceError struct {
	ogError  error
	svcError error
}

func (e *ServiceError) GetOgError() error {
	return e.ogError
}

func (e *ServiceError) GetSvcError() error {
	return e.svcError
}

func NewServiceError(ogError, svcError error) *ServiceError {
	return &ServiceError{
		ogError:  ogError,
		svcError: svcError,
	}
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("original error: %v, service error: %v", e.ogError, e.svcError)
}
