package entities

import (
	"time"

	"gorm.io/gorm"
)

type BoardMember struct {
	ID uint `json:"id" gorm:"<-:create;primaryKey"`

	BoardUserID uint      `json:"-" gorm:"NOTNULL"`
	BoardUser   BoardUser `json:"-" gorm:"NOTNULL"`
	BoardID     uint      `json:"-" gorm:"NOTNULL"`
	Board       Board     `json:"-" gorm:"NOTNULL"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
