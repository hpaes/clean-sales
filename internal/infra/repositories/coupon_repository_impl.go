package repositories

import (
	"clean-sales/internal/entities"
	"database/sql"
)

type CouponRepositoryImpl struct {
	db *sql.DB
}

func NewCouponRepositoryImpl(db *sql.DB) CouponRepository {
	return &CouponRepositoryImpl{db: db}
}

// GetCoupon implements CouponRepository
func (c *CouponRepositoryImpl) GetCoupon(code string) (*entities.Coupon, error) {
	stmt, err := c.db.Prepare("SELECT * FROM coupons WHERE code = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var coupon entities.Coupon
	err = stmt.QueryRow(code).Scan(&coupon.Code, &coupon.Discount, &coupon.ExpireAt)
	if err != nil {
		return nil, err
	}

	return &coupon, nil
}
