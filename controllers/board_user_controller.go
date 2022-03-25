package controllers

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/services"
	"net/http"

	"coupon-service/infrastructure/auth"
	apiError "coupon-service/infrastructure/errors"

	"github.com/gin-gonic/gin"
)

type BoardUserController interface {
	GetUsers(c *gin.Context)
	GetSelfInfo(c *gin.Context)
	SignInGoogle(c *gin.Context)
}

type BoardUserControllerInstance struct {
	boardUserSv services.BoardUserService
}

func NewBoardUserController(
	boardUserSv services.BoardUserService,
) BoardUserController {
	return &BoardUserControllerInstance{
		boardUserSv: boardUserSv,
	}
}

func (buCtrl *BoardUserControllerInstance) SignInGoogle(c *gin.Context) {
	tokenReq := new(entities.BoardUserSignInGoogle)

	err := c.BindJSON(tokenReq)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := buCtrl.boardUserSv.SignInGoogle(c, tokenReq)
	if err != nil {
		if apiErr, ok := err.(*apiError.APIError); ok {
			responsMessageHttp(c, apiErr.StatusCode, apiErr.Err.Error())
			return
		}
		responsMessageHttp(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseItemHttp(c, http.StatusOK, user)
}

func (buCtrl *BoardUserControllerInstance) GetUsers(c *gin.Context) {
	users, err := buCtrl.boardUserSv.Find(c)
	if err != nil {
		if apiErr, ok := err.(*apiError.APIError); ok {
			responsMessageHttp(c, apiErr.StatusCode, apiErr.Err.Error())
			return
		}
		responsMessageHttp(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseListHttp(c, http.StatusOK, users, len(users))
}

func (buCtrl *BoardUserControllerInstance) GetSelfInfo(c *gin.Context) {
	user, err := buCtrl.boardUserSv.FindByGoogleUserID(c, c.GetString(auth.AuthGUserIDContextKey))
	if err != nil {
		if apiErr, ok := err.(*apiError.APIError); ok {
			responsMessageHttp(c, apiErr.StatusCode, apiErr.Err.Error())
			return
		}
		responsMessageHttp(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseItemHttp(c, http.StatusOK, user)
}
