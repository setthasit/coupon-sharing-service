package errors

import (
	"errors"
	"net/http"
)

func ErrCouponHistoyUsageFailedCopyAction() *APIError {
	return NewAPIError(http.StatusNotFound, errors.New("copy failed"))
}

func ErrCouponHistoyUsageFailedUseAction() *APIError {
	return NewAPIError(http.StatusNotFound, errors.New("copy failed"))
}

func ErrCouponHistoyUsageFailedCancelAction() *APIError {
	return NewAPIError(http.StatusNotFound, errors.New("copy failed"))
}
