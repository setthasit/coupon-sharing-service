package entities

import (
	"coupon-service/infrastructure/security"
	"time"

	"gorm.io/gorm"
)

type BoardUser struct {
	ID        uint                `json:"id" gorm:"<-:create;primaryKey"`
	Name      string              `json:"name" gorm:"type:varchar;size:255"`
	Email     string              `json:"email" gorm:"uniqueIndex;type:varchar;size:100"`
	Password  security.EncryptVal `json:"-"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	DeletedAt gorm.DeletedAt      `json:"deleted_at"`
}

type BoardUserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BoardUserSignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *BoardUserRegister) ToBoardUser() *BoardUser {
	return &BoardUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: security.EncryptVal(user.Password),
	}
}
