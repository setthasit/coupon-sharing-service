package controllers

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BoardUserController interface {
	CreateNewUser(c *gin.Context)
}

type BoardUserControllerInstance struct {
	BoardUserSv services.BoardUserService
}

func NewBoardUserController(BoardUserSv services.BoardUserService) BoardUserController {
	return &BoardUserControllerInstance{
		BoardUserSv: BoardUserSv,
	}
}

func (buCtrl *BoardUserControllerInstance) CreateNewUser(c *gin.Context) {
	newUser := new(entities.BoardUserRegister)

	err := c.BindJSON(newUser)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, err.Error())
	}

	responsMessageHttp(c, http.StatusCreated, "user created")
}
