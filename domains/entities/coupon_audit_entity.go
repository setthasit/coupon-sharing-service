package entities

import (
	"time"

	"gorm.io/gorm"
)

type CouponAuditLog struct {
	ID       uint   `json:"id" gorm:"<-:create;primaryKey"`
	CouponID uint   `json:"coupon_id"`
	Name     string `json:"name" gorm:"type:varchar;size:255"`

	BoardID uint `json:"-"`

	App             CouponApp       `json:"coupon_app"`
	CouponType      CouponType      `json:"coupon_type"`
	DiscountService DiscountService `json:"discount_service"`
	Quantity        uint            `json:"quantity"`
	DiscountType    DiscountType    `json:"discount_type"`
	DiscountAmount  float64         `json:"discount_amount"`

	CreatedAt      time.Time      `json:"created_at"`
	CreatedBy      uint           `json:"created_by"`
	UpdatedAt      time.Time      `json:"updated_at"`
	UpdatedBy      uint           `json:"updated_by"`
	UpdatedMessage string         `json:"updated_message"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	DeletedBy      uint           `json:"deleted_by"`
}
