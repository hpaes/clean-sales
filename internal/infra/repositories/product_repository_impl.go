package repositories

import (
	"clean-sales/internal/entities"
	"database/sql"
)

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepositoryImpl(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

// GetProduct implements ProductRepository
func (p *ProductRepositoryImpl) GetProduct(id string) (*entities.Product, error) {
	stmt, err := p.db.Prepare("SELECT * FROM products WHERE id_product = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product entities.Product

	err = stmt.QueryRow(id).Scan(&product.IdProduct, &product.Description, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
