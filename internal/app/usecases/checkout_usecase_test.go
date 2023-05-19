package usecases

import (
	"clean-sales/internal/infra/repositories"
	testfixture "clean-sales/testFixture"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CheckoutUseCaseTestSuite struct {
	suite.Suite
	checkoutUseCase CheckoutUseCase
}

func (suite *CheckoutUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", "file::memory:")
	if err != nil {
		suite.T().Fatal(err)
	}

	if err := testfixture.PrepDb(db); err != nil {
		suite.T().Fatal(err)
	}

	couponRepository := repositories.NewCouponRepositoryImpl(db)
	productRepository := repositories.NewProductRepositoryImpl(db)
	orderRepository := repositories.NewOrderRepositoryImpl(db)

	suite.checkoutUseCase = NewCheckoutUseCaseImpl(productRepository, couponRepository, orderRepository)
}

func TestCheckoutSuite(t *testing.T) {
	suite.Run(t, new(CheckoutUseCaseTestSuite))
}

func (suite *CheckoutUseCaseTestSuite) TearDown() {
	suite.T().Log("Finished test")
}

func (suite *CheckoutUseCaseTestSuite) TestGivenInvalidCPF_ThenShouldNotCreateOrder() {
	output, err := suite.checkoutUseCase.Execute(&testfixture.InvalidInput)
	assert.EqualError(suite.T(), err, "could not create order: invalid cpf")
	assert.Nil(suite.T(), output)
}

func (suite *CheckoutUseCaseTestSuite) TestGivenValidCPF_ThenShouldCreateOrder() {
	output, err := suite.checkoutUseCase.Execute(&testfixture.ValidInput)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 6030.00, output.Total)

}

func (suite *CheckoutUseCaseTestSuite) TestGivenValidCPFAndCoupon_ThenShouldCreateOrder() {
	output, err := suite.checkoutUseCase.Execute(&testfixture.ValidCouponInput)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 5427.00, output.Total)
}

func (suite *CheckoutUseCaseTestSuite) TestGivenDuplicatedItems_ThenShouldNotCreateOrder() {
	output, err := suite.checkoutUseCase.Execute(&testfixture.DuplicatedItem)
	assert.EqualError(suite.T(), err, "could not create order: duplicated item")
	assert.Nil(suite.T(), output)
}

func (suite *CheckoutUseCaseTestSuite) TestGivenNegativeItemQuantity_ThenShouldNotCreateOrder() {
	output, err := suite.checkoutUseCase.Execute(&testfixture.NegativeItemQuantity)
	assert.EqualError(suite.T(), err, "could not create order: invalid item quantity")
	assert.Nil(suite.T(), output)
}

func (suite *CheckoutUseCaseTestSuite) TestGivenFromAndTo_ThenShouldCalculateFreightAndCreateOrder() {
	output, err := suite.checkoutUseCase.Execute(&testfixture.ValidInputWithFreight)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 6290.00, output.Total)
	assert.Equal(suite.T(), 260.00, output.Freight)
}

func (suite *CheckoutUseCaseTestSuite) TestGivenInvalidProductDimensions_ThenShouldNotCreateOrder() {
	output, err := suite.checkoutUseCase.Execute(&testfixture.NegativeItemDimension)
	assert.EqualError(suite.T(), err, "failed to create product: invalid product dimensions")
	assert.Nil(suite.T(), output)
}
