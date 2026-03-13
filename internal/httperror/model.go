package httperror

import (
	"errors"
)

type HttpError struct {
	InternalError error  `json:"-"`
	PublicError   string `json:"errorMessage"`
}

// Implements the error interface
func (err *HttpError) Error() string {
	return err.InternalError.Error()
}

// Implements the pubinterr interface
func (err *HttpError) Public() string {
	return err.PublicError
}

// New creates a new HttpError, setting the internal error to the public error if
// err param is nil
func New(err error, message string) *HttpError {
	if err == nil {
		return &HttpError{InternalError: errors.New(message), PublicError: message}
	}
	return &HttpError{InternalError: err, PublicError: message}
}

// OptNew is used when you don't want to check if the error exists before creating an HTTPError,
// returning nil if the internal error is nil
func OptNew(err error, message string) *HttpError {
	if err == nil {
		return nil
	}
	return &HttpError{InternalError: err, PublicError: message}
}

// NewStr sets both internal and public to the same str value
func NewStr(message string) *HttpError {
	return &HttpError{InternalError: errors.New(message), PublicError: message}
}
