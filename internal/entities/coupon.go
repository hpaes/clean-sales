package entities

import "time"

type Coupon struct {
	Code     string
	Discount float64
	ExpireAt string
}

func NewCoupon(code string, discount float64, expire_at string) *Coupon {
	return &Coupon{
		Code:     code,
		Discount: discount,
		ExpireAt: expire_at,
	}
}

func (c *Coupon) IsValid() bool {
	expireAt, err := time.Parse(time.RFC3339, c.ExpireAt)
	if err != nil {
		return true
	}

	return time.Now().After(expireAt)
}

func (c *Coupon) CalculateDiscount(total float64) float64 {
	return total * (c.Discount / 100)
}
