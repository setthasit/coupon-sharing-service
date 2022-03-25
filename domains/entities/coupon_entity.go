package entities

import (
	"time"

	"gorm.io/gorm"
)

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
	CreatedBy      uint           `json:"created_by"`
	UpdatedAt      time.Time      `json:"updated_at"`
	UpdatedBy      uint           `json:"updated_by"`
	UpdatedMessage string         `json:"updated_message"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	DeletedBy      uint           `json:"deleted_by"`
}

type CouponCreateNew struct {
	BoardID         uint            `json:"board_id"`
	Name            string          `json:"name" binding:"required"`
	App             CouponApp       `json:"coupon_app" binding:"required"`
	CouponType      CouponType      `json:"coupon_type" binding:"required"`
	DiscountService DiscountService `json:"discount_service" binding:"required"`
	Quantity        uint            `json:"quantity" binding:"required"`
	DiscountType    DiscountType    `json:"discount_type" binding:"required"`
	DiscountAmount  float64         `json:"discount_amount" binding:"required"`
}

type CouponBulkCreateNew struct {
	BoardID uint              `json:"board_id" binding:"required"`
	Coupons []CouponCreateNew `json:"coupons" binding:"required"`
}

func (newC *CouponCreateNew) ToCouponWithAudit(boardUserID uint) *Coupon {
	return &Coupon{
		BoardID:         newC.BoardID,
		Name:            newC.Name,
		App:             newC.App,
		CouponType:      newC.CouponType,
		DiscountService: newC.DiscountService,
		Quantity:        newC.Quantity,
		DiscountType:    newC.DiscountType,
		DiscountAmount:  newC.DiscountAmount,
		CreatedBy:       boardUserID,
		UpdatedBy:       boardUserID,
	}
}

func (newC *CouponBulkCreateNew) ToCouponsWithAudit(boardUserID uint) []Coupon {
	coupons := make([]Coupon, 0)
	for _, coupon := range newC.Coupons {
		coupons = append(coupons, Coupon{
			BoardID:         newC.BoardID,
			Name:            coupon.Name,
			App:             coupon.App,
			CouponType:      coupon.CouponType,
			DiscountService: coupon.DiscountService,
			Quantity:        coupon.Quantity,
			DiscountType:    coupon.DiscountType,
			DiscountAmount:  coupon.DiscountAmount,
			CreatedBy:       boardUserID,
			UpdatedBy:       boardUserID,
		})
	}

	return coupons
}
