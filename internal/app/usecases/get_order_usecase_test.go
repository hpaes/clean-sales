package usecases

import (
	"clean-sales/internal/infra/repositories"
	testfixture "clean-sales/testFixture"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type GetOrderUseCaseTestSuite struct {
	suite.Suite
	getOrderUseCase GetOrderUseCase
}

func (suite *GetOrderUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", "file::memory:")
	if err != nil {
		suite.T().Fatal(err)
	}

	if err := testfixture.PrepDb(db); err != nil {
		suite.T().Fatal(err)
	}

	orderRepository := repositories.NewOrderRepositoryImpl(db)

	suite.getOrderUseCase = NewGetOrder(orderRepository)
}

func TestGetOrderSuite(t *testing.T) {
	suite.Run(t, new(GetOrderUseCaseTestSuite))
}

func (suite *GetOrderUseCaseTestSuite) TearDown() {
	suite.T().Log("Finished test")
}

func (suite *GetOrderUseCaseTestSuite) TestGetOrder() {
	order, err := suite.getOrderUseCase.Execute(testfixture.ValidOrderId)
	suite.NoError(err)

	suite.Equal(6030.00, order.Total)
}
