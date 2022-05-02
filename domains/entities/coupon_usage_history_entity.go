package entities

import (
	"time"
)

type CouponAction string

const (
	CouponActionCopy   CouponAction = "COPY"
	CouponActionUse    CouponAction = "USE"
	CouponActionCancel CouponAction = "CANCEL"
)

type CouponUsageHistory struct {
	ID uint `json:"id" gorm:"<-:create;primaryKey"`

	BoardID     uint `json:"-"`
	BoardUserID uint `json:"-"`
	CouponID    uint `json:"-"`

	Action   CouponAction `json:"action"`
	ActionAt time.Time    `json:"action_at"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy uint      `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy uint      `json:"updated_by"`
}

type CouponCopyRequest struct {
	BoardUserID uint `json:"-"`
	BoardID     uint `json:"board_id" binding:"required"`
	CouponID    uint `json:"coupon_id" binding:"required"`
}

func (req *CouponCopyRequest) ToCouponUsageHistory() *CouponUsageHistory {
	return &CouponUsageHistory{
		BoardID:     req.BoardID,
		BoardUserID: req.BoardUserID,
		CouponID:    req.CouponID,
		Action:      CouponActionCopy,
		ActionAt:    time.Now(),
		CreatedBy:   req.BoardUserID,
		UpdatedBy:   req.BoardUserID,
	}
}

type CouponUseRequest struct {
	BoardUserID uint `json:"-"`
	BoardID     uint `json:"board_id" binding:"required"`
	CouponID    uint `json:"coupon_id" binding:"required"`
}

func (req *CouponUseRequest) ToCouponUsageHistory() *CouponUsageHistory {
	return &CouponUsageHistory{
		BoardID:     req.BoardID,
		BoardUserID: req.BoardUserID,
		CouponID:    req.CouponID,
		Action:      CouponActionUse,
		ActionAt:    time.Now(),
		CreatedBy:   req.BoardUserID,
		UpdatedBy:   req.BoardUserID,
	}
}

type CouponCancelRequest struct {
	BoardUserID uint `json:"-"`
	BoardID     uint `json:"board_id" binding:"required"`
	CouponID    uint `json:"coupon_id" binding:"required"`
}

func (req *CouponCancelRequest) ToCouponUsageHistory() *CouponUsageHistory {
	return &CouponUsageHistory{
		BoardID:     req.BoardID,
		BoardUserID: req.BoardUserID,
		CouponID:    req.CouponID,
		Action:      CouponActionCancel,
		ActionAt:    time.Now(),
		CreatedBy:   req.BoardUserID,
		UpdatedBy:   req.BoardUserID,
	}
}
