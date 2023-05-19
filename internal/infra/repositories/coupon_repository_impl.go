package repositories

import (
	"clean-sales/internal/app/entities"
	"database/sql"
	"fmt"
)

type CouponRepository interface {
	GetCoupon(code string) (*entities.Coupon, error)
}

type CouponRepositoryImpl struct {
	database *sql.DB
}

func NewCouponRepositoryImpl(db *sql.DB) CouponRepository {
	return &CouponRepositoryImpl{database: db}
}

// GetCoupon implements CouponRepository
func (c *CouponRepositoryImpl) GetCoupon(code string) (*entities.Coupon, error) {
	query := "SELECT * FROM coupons WHERE code = ?"

	row := c.database.QueryRow(query, code)

	var coupon entities.Coupon

	err := row.Scan(&coupon.Code, &coupon.Discount, &coupon.ExpireAt)
	if err != nil {
		return nil, fmt.Errorf("failed to scan coupon data: %w", err)
	}

	return &coupon, nil
}
