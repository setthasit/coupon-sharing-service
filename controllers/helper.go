package controllers

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/services"
	"coupon-service/infrastructure/auth"
	apiError "coupon-service/infrastructure/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	if apiErr, ok := err.(*apiError.APIError); ok {
		responsMessageHttp(c, apiErr.StatusCode, apiErr.Err.Error())
		return
	}
	responsMessageHttp(c, http.StatusInternalServerError, err.Error())
	return
}

func getUserInfoFromGinContext(c *gin.Context, sv services.BoardUserService) (*entities.BoardUser, error) {
	user, err := sv.FindByGoogleUserID(c, c.GetString(auth.AuthGUserIDContextKey))
	if err != nil {
		return nil, err
	}

	return user, nil
}
