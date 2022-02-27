package services

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/repositories"
	"coupon-service/infrastructure/security"

	"github.com/gin-gonic/gin"
)

type BoardUserService interface {
	Find(ctx *gin.Context) []entities.BoardUser
	Register(ctx *gin.Context, reqUser *entities.BoardUserRegister) *entities.BoardUser
}

type BoardUserServiceInstance struct {
	boardUserRepo repositories.BoardUserRepository
}

func NewBoardUserService(
	boardUserRepo repositories.BoardUserRepository,
) BoardUserService {
	return &BoardUserServiceInstance{
		boardUserRepo: boardUserRepo,
	}
}

func (sv *BoardUserServiceInstance) Find(ctx *gin.Context) []entities.BoardUser {
	users, err := sv.boardUserRepo.Find(ctx)
	if err != nil {
		ctx.Error(err)
		return nil
	}

	return users
}

func (sv *BoardUserServiceInstance) Register(ctx *gin.Context, reqUser *entities.BoardUserRegister) *entities.BoardUser {
	newUser := reqUser.ToBoardUser()
	encryptPassword, err := security.Encrypt(reqUser.Password)
	if err != nil {
		ctx.Error(err)
		return nil
	}
	newUser.Password = encryptPassword

	createdUser, err := sv.boardUserRepo.Create(ctx, newUser)
	if err != nil {
		ctx.Error(err)
		return nil
	}

	return createdUser
}
