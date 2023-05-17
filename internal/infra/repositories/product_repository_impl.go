package repositories

import (
	"clean-sales/internal/entities"
	"database/sql"
	"fmt"
)

type ProductRepositoryImpl struct {
	database *sql.DB
}

func NewProductRepositoryImpl(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{database: db}
}

// GetProduct implements ProductRepository
func (p *ProductRepositoryImpl) GetProduct(id string) (*entities.Product, error) {
	query := "SELECT * FROM products WHERE id_product = ?"

	row := p.database.QueryRow(query, id)

	var productData entities.Product
	err := row.Scan(&productData.IdProduct, &productData.Description, &productData.Price, &productData.Width, &productData.Height, &productData.Length, &productData.Weight)
	if err != nil {
		return nil, fmt.Errorf("failed to scan product data: %w", err)
	}

	product, err := entities.NewProduct(productData.IdProduct, productData.Description, productData.Price, productData.Width, productData.Height, productData.Length, productData.Weight)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}
