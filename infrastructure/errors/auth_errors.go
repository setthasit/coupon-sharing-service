package errors

import (
	"errors"
	"net/http"
)

func ErrAuthInvalidToken() *APIError {
	return NewAPIError(http.StatusForbidden, errors.New("invalid token"))
}
