package entities

import (
	"time"

	"gorm.io/gorm"
)

type Board struct {
	ID   uint   `json:"id" gorm:"<-:create;primaryKey"`
	Name string `json:"name" gorm:"type:varchar;size:255"`

	BoardUserID uint `json:"-"`

	BoardMember []BoardMember `json:"members" gorm:"foreignKey:BoardID;references:ID"`
	Coupons     []Coupon      `json:"coupons" gorm:"foreignKey:BoardID;references:ID"`

	CreatedAt time.Time      `json:"created_at"`
	CreatedBy uint           `json:"created_by"`
	UpdatedAt time.Time      `json:"updated_at"`
	UpdatedBy uint           `json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	DeletedBy uint           `json:"deleted_by"`
}

type BoardCreateNew struct {
	Name        string        `json:"name" binding:"required"`
	BoardUserID uint          `json:"-"`
	BoardMember []BoardMember `json:"member"`
}

func (bNew *BoardCreateNew) ToBoardWithAudit(boardUserID uint) *Board {
	return &Board{
		Name:        bNew.Name,
		BoardMember: bNew.BoardMember,
		BoardUserID: boardUserID,
		CreatedBy:   boardUserID,
		UpdatedBy:   boardUserID,
	}
}
