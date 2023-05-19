package entities

import (
	testfixture "clean-sales/testFixture"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCheckIfCouponIsValid(t *testing.T) {
	coupon := NewCoupon("VALE20", 20, testfixture.ValidDate)
	assert.Equal(t, true, coupon.IsValid())
}

func TestShouldCalculateDiscount(t *testing.T) {
	coupon := NewCoupon("VALE20", 20, testfixture.ValidDate)
	assert.Equal(t, 200.0, coupon.CalculateDiscount(1000.0))
}
