package errors

import (
	"errors"
	"net/http"
)

func ErrCouponNotFound() *APIError {
	return NewAPIError(http.StatusNotFound, errors.New("coupon not found"))
}

func ErrCouponCannotCreate() *APIError {
	return NewAPIError(http.StatusNotFound, errors.New("create coupon failed"))
}
