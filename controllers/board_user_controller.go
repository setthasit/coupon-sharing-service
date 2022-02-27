package controllers

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BoardUserController interface {
	GetUsers(c *gin.Context)
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

func (buCtrl *BoardUserControllerInstance) GetUsers(c *gin.Context) {
	users := buCtrl.BoardUserSv.Find(c)
	responseListHttp(c, http.StatusOK, users, len(users))
}

func (buCtrl *BoardUserControllerInstance) CreateNewUser(c *gin.Context) {
	newUser := new(entities.BoardUserRegister)

	err := c.BindJSON(newUser)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, err.Error())
		return
	}

	_ = buCtrl.BoardUserSv.Register(c, newUser)
	if c.Err() != nil {
		responsMessageHttp(c, http.StatusBadRequest, c.Err().Error())
		return
	}

	responsMessageHttp(c, http.StatusCreated, "user created")
}
