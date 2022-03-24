package errors

import (
	"fmt"
)

type APIError struct {
	Err        error
	StatusCode int
	InterlCode *string
}

func (err *APIError) Error() (message string) {
	if err.InterlCode != nil {
		message = fmt.Sprintf(
			"error with HTTP Status: %d Internal: %s message: %s",
			err.StatusCode,
			*err.InterlCode,
			err.Err.Error(),
		)
	} else {
		message = fmt.Sprintf(
			"error with HTTP Status: %d message: %s",
			err.StatusCode,
			err.Err.Error(),
		)
	}
	return
}

func NewAPIError(code int, err error) *APIError {
	return &APIError{
		Err:        err,
		StatusCode: code,
	}
}

func NewAPIErrorWithInternal(code int, err error, internal string) *APIError {
	return &APIError{
		Err:        err,
		StatusCode: code,
		InterlCode: &internal,
	}
}
