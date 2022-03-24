package auth

import (
	"coupon-service/infrastructure/errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthGUserIDContextKey = "g_auth_sub"
)

type AuthMiddleware struct {
	googleOAuth *GoogleOAuth
}

func NewAuthMiddleware(
	googleOAuth *GoogleOAuth,
) *AuthMiddleware {
	return &AuthMiddleware{
		googleOAuth: googleOAuth,
	}
}

func (am *AuthMiddleware) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		if len(bearerToken) < 7 {
			log.Print("token length if not sufficient")
			err := errors.ErrAuthInvalidToken()

			c.AbortWithStatusJSON(err.StatusCode, gin.H{
				"message": err.Err.Error(),
			})
			return
		}

		tokenType := bearerToken[:6]
		token := bearerToken[7:]

		if strings.ToLower(tokenType) != "bearer" {
			log.Print("token is not bearer")
			err := errors.ErrAuthInvalidToken()
			c.AbortWithStatusJSON(err.StatusCode, gin.H{
				"message": err.Err.Error(),
			})
			return
		}

		if token == "" {
			err := errors.ErrAuthInvalidToken()
			c.AbortWithStatusJSON(err.StatusCode, gin.H{
				"message": err.Err.Error(),
			})
			return
		}

		userInfo, err := am.googleOAuth.GetUserData(c, token)
		if err != nil {
			err := errors.ErrAuthInvalidToken()
			c.AbortWithStatusJSON(err.StatusCode, gin.H{
				"message": err.Err.Error(),
			})
			return
		}

		c.Set(AuthGUserIDContextKey, userInfo.Id)

		c.Next()
	}
}
