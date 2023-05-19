package entities

import (
	"log"
	"time"
)

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
	expireAt, err := time.Parse("2006-01-02", c.ExpireAt)
	if err != nil {
		return true
	}

	log.Println(time.Now().String())
	return time.Now().After(expireAt)
}

func (c *Coupon) CalculateDiscount(total float64) float64 {
	return total * (c.Discount / 100)
}
