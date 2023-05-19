package repositories

import (
	testfixture "clean-sales/testFixture"
	"database/sql"
	"testing"

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

	if err := testfixture.PrepDb(db); err != nil {
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
	repo := NewCouponRepositoryImpl(suite.db)
	coupon, err := repo.GetCoupon("CUPOM10")
	suite.NoError(err)

	suite.Equal("CUPOM10", coupon.Code)
	suite.Equal(10.0, coupon.Discount)
}
