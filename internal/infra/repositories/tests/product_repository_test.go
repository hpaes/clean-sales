package repositories

import (
	"clean-sales/internal/infra/repositories"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ProductRepositoryTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (suite *ProductRepositoryTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", "file::memory:")
	if err != nil {
		suite.T().Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id_product text PRIMARY KEY NOT NULL,description text, price DECIMAL(10,2));")
	if err != nil {
		suite.T().Fatal(err)
	}

	_, err = db.Exec("INSERT INTO products (id_product, description, price) VALUES ('1', 'A', 1000);")
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.db = db
}

func (suite *ProductRepositoryTestSuite) TearDownTest() {
	suite.NoError(suite.db.Close())
}

func TestProductRepositorySuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}

func (suite *ProductRepositoryTestSuite) TestGetProduct() {
	repo := repositories.NewProductRepositoryImpl(suite.db)
	product, err := repo.GetProduct("1")
	suite.NoError(err)

	suite.Equal("1", product.IdProduct)
	suite.Equal("A", product.Description)
	suite.Equal(1000.0, product.Price)
}
