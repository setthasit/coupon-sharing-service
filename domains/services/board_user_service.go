package services

import (
	"coupon-service/domains/entities"
	"coupon-service/domains/repositories"
	"coupon-service/infrastructure/auth"
	"coupon-service/infrastructure/errors"

	"github.com/gin-gonic/gin"
)

type BoardUserService interface {
	Find(ctx *gin.Context) ([]entities.BoardUser, error)
	FindByGoogleUserID(ctx *gin.Context, googleUserID string) (*entities.BoardUser, error)
	SignInGoogle(ctx *gin.Context, tokenReq *entities.BoardUserSignInGoogle) (*entities.BoardUser, error)
}

type BoardUserServiceInstance struct {
	boardUserRepo repositories.BoardUserRepository
	googleOAuth   *auth.GoogleOAuth
}

func NewBoardUserService(
	boardUserRepo repositories.BoardUserRepository,
	googleOAuth *auth.GoogleOAuth,
) BoardUserService {
	return &BoardUserServiceInstance{
		boardUserRepo: boardUserRepo,
		googleOAuth:   googleOAuth,
	}
}

func (sv *BoardUserServiceInstance) Find(ctx *gin.Context) ([]entities.BoardUser, error) {
	users, err := sv.boardUserRepo.Find(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (sv *BoardUserServiceInstance) SignInGoogle(ctx *gin.Context, tokenReq *entities.BoardUserSignInGoogle) (*entities.BoardUser, error) {
	userInfo, err := sv.googleOAuth.GetUserData(ctx, tokenReq.TokenID)
	if err != nil {
		return nil, err
	}

	user, err := sv.boardUserRepo.FindByGoogleUserID(ctx, userInfo.Id)
	if err != nil {
		if ok, _ := errors.CompareError(err, errors.ErrBoardUserUserNotFound()); ok {
			err = nil
			user, err = sv.boardUserRepo.Create(ctx, &entities.BoardUser{
				Name:         userInfo.Name,
				Email:        userInfo.Email,
				GoogleUserID: userInfo.Id,
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (sv *BoardUserServiceInstance) FindByGoogleUserID(ctx *gin.Context, googleUserID string) (*entities.BoardUser, error) {
	user, err := sv.boardUserRepo.FindByGoogleUserID(ctx, googleUserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
