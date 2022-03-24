package entities

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	ID   uint   `json:"id" gorm:"<-:create;primaryKey"`
	Name string `json:"name" gorm:"type:varchar;size:255"`

	BoardID uint `json:"-"`

	App             CouponApp       `json:"coupon_app"`
	CouponType      CouponType      `json:"coupon_type"`
	DiscountService DiscountService `json:"discount_service"`
	Quantity        uint            `json:"quantity"`
	DiscountType    DiscountType    `json:"discount_type"`
	DiscountAmount  float64         `json:"discount_amount"`

	CreatedAt      time.Time      `json:"created_at"`
	CreatedBy      int            `json:"created_by"`
	UpdatedAt      time.Time      `json:"updated_at"`
	UpdatedBy      int            `json:"updated_by"`
	UpdatedMessage string         `json:"updated_message"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	DeletedBy      int            `json:"deleted_by"`
}

type CouponApp string
type CouponType string
type DiscountService string
type DiscountType string

const (
	DiscountTypePercent DiscountType = "PERCENT"
	DiscountTypeFix     DiscountType = "FIX"

	DiscountServiceFoodDiscount DiscountService = "FOOD_DISCOUNT"
	DiscountServiceDeliveryFee  DiscountService = "DELIVERY_FEE"

	CouponTypeSingle   CouponType = "SINGLE_USE"
	CouponTypeMultiple CouponType = "MULTIPLE_USE"

	CouponAppGrab    CouponApp = "GRAB"
	CouponAppLineMan CouponApp = "LINE_MAN"
)
