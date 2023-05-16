package repositories

import (
	"clean-sales/internal/infra/repositories"
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CouponRepositoryTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (suite *CouponRepositoryTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", "file::memory:")
	if err != nil {
		suite.T().Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS coupons (code text PRIMARY KEY NOT NULL, discount DECIMAL(10,2), expire_at text);")
	if err != nil {
		suite.T().Fatal(err)
	}

	validDate := time.Now().AddDate(1, 0, 0).Format("2006-01-02")
	invalidDate := time.Now().AddDate(-1, 0, 0).Format("2006-01-02")

	_, err = db.Exec("INSERT INTO coupons (code, discount, expire_at) VALUES ('CUPOM10',10.00, '" + validDate + "');")
	if err != nil {
		suite.T().Fatal(err)
	}

	_, err = db.Exec("INSERT INTO coupons (code, discount, expire_at) VALUES ('EXPIRED',10.00, '" + invalidDate + "');")
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.db = db
}

func (suite *CouponRepositoryTestSuite) TearDownTest() {
	suite.NoError(suite.db.Close())
}

func TestCouponRepositorySuite(t *testing.T) {
	suite.Run(t, new(CouponRepositoryTestSuite))
}

func (suite *CouponRepositoryTestSuite) TestGetCoupon() {
	repo := repositories.NewCouponRepositoryImpl(suite.db)
	coupon, err := repo.GetCoupon("CUPOM10")
	suite.NoError(err)

	suite.Equal("CUPOM10", coupon.Code)
	suite.Equal(10.0, coupon.Discount)
}
