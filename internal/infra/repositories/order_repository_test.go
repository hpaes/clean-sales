package repositories

import (
	"clean-sales/internal/app/entities"
	testfixture "clean-sales/testFixture"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	db    *sql.DB
	order *entities.Order
}

func (suite *OrderRepositoryTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", "file::memory:")
	if err != nil {
		suite.T().Fatal(err)
	}

	if err := testfixture.PrepDb(db); err != nil {
		suite.T().Fatal(err)
	}

	order, err := entities.NewOrder(uuid.NewString(), testfixture.ValidCPF, 0)
	suite.NoError(err)
	p1, _ := entities.NewProduct("1", "A", 1000, 100, 30, 10, 3)
	p2, _ := entities.NewProduct("2", "B", 5000, 50, 50, 50, 22)
	p3, _ := entities.NewProduct("3", "C", 30, 10, 10, 10, 0.9)
	order.AddItem(p1, 1)
	order.AddItem(p2, 1)
	order.AddItem(p3, 1)

	suite.order = order
	suite.db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.NoError(suite.db.Close())
}

func TestOrderRepositorySuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGetOrder() {
	repo := NewOrderRepositoryImpl(suite.db)

	order, err := repo.GetOrder(testfixture.ValidOrderId)
	suite.NoError(err)
	suite.Assert().Equal(testfixture.ValidOrderId, order.IdOrder)
}

func (suite *OrderRepositoryTestSuite) TestShouldCountOrder() {
	repo := NewOrderRepositoryImpl(suite.db)

	count, err := repo.CountOrder()
	suite.NoError(err)
	suite.Assert().Equal(1, count)
}

func (suite *OrderRepositoryTestSuite) TestShouldSaveOrder() {
	repo := NewOrderRepositoryImpl(suite.db)

	err := repo.SaveOrder(suite.order)
	suite.NoError(err)

	count, err := repo.CountOrder()
	suite.NoError(err)
	suite.Assert().Equal(2, count)
}
