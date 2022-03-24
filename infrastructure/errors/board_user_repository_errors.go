package errors

import (
	"errors"
	"net/http"
)

func ErrBoardUserIncorrectAuthen() *APIError {
	return NewAPIError(http.StatusBadRequest, errors.New("incorrect email/password"))
}

func ErrBoardUserUserNotFound() *APIError {
	return NewAPIError(http.StatusNotFound, errors.New("user not found"))
}

func ErrBoardUserCreateFailed() *APIError {
	return NewAPIError(http.StatusNotFound, errors.New("create user failed"))
}
