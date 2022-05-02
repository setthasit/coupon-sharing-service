package errors

import (
	"errors"
	"net/http"
)

func ErrBoardMemberNotMember() *APIError {
	return NewAPIError(http.StatusForbidden, errors.New("you are not member of this board"))
}
