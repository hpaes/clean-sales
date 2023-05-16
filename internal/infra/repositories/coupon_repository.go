package repositories

import "clean-sales/internal/entities"

type CouponRepository interface {
	GetCoupon(code string) (*entities.Coupon, error)
}
