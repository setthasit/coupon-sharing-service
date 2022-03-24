package controllers

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/services"
	"net/http"

	"coupon-service/infrastructure/auth"
	apiError "coupon-service/infrastructure/errors"

	"github.com/gin-gonic/gin"
)

type BoardController interface {
	GetBoardByUser(c *gin.Context)
	CreateNewBoard(c *gin.Context)
}

type BoardControllerInstance struct {
	BoardSv     services.BoardService
	BoardUserSv services.BoardUserService
}

func NewBoardController(
	boardSv services.BoardService,
	BoardUserSv services.BoardUserService,
) BoardController {
	return &BoardControllerInstance{
		BoardSv:     boardSv,
		BoardUserSv: BoardUserSv,
	}
}

func (buCtrl *BoardControllerInstance) GetBoardByUser(c *gin.Context) {
	userInfo, err := buCtrl.getUserInfo(c)
	if err != nil {
		if apiErr, ok := err.(*apiError.APIError); ok {
			responsMessageHttp(c, apiErr.StatusCode, apiErr.Err.Error())
			return
		}
		responsMessageHttp(c, http.StatusInternalServerError, err.Error())
		return
	}

	boards, err := buCtrl.BoardSv.Find(c, userInfo.ID)
	if err != nil {
		if apiErr, ok := err.(*apiError.APIError); ok {
			responsMessageHttp(c, apiErr.StatusCode, apiErr.Err.Error())
			return
		}
		responsMessageHttp(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseListHttp(c, http.StatusOK, boards, len(boards))
}

func (buCtrl *BoardControllerInstance) CreateNewBoard(c *gin.Context) {
	newBoardReq := new(entities.BoardCreateNew)
	err := c.BindJSON(newBoardReq)
	if err != nil {
		responsMessageHttp(c, http.StatusBadRequest, err.Error())
		return
	}

	userInfo, err := buCtrl.getUserInfo(c)
	if err != nil {
		if apiErr, ok := err.(*apiError.APIError); ok {
			responsMessageHttp(c, apiErr.StatusCode, apiErr.Err.Error())
			return
		}
		responsMessageHttp(c, http.StatusInternalServerError, err.Error())
		return
	}

	newBoardReq.BoardUserID = userInfo.ID
	board, err := buCtrl.BoardSv.CreateNewBoard(c, newBoardReq)
	if err != nil {
		if apiErr, ok := err.(*apiError.APIError); ok {
			responsMessageHttp(c, apiErr.StatusCode, apiErr.Err.Error())
			return
		}
		responsMessageHttp(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseItemHttp(c, http.StatusOK, board)
}

func (buCtrl *BoardControllerInstance) getUserInfo(c *gin.Context) (*entities.BoardUser, error) {
	user, err := buCtrl.BoardUserSv.FindByGoogleUserID(c, c.GetString(auth.AuthGUserIDContextKey))
	if err != nil {
		return nil, err
	}

	return user, nil
}
