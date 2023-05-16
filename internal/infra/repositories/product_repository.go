package repositories

import "clean-sales/internal/entities"

type ProductRepository interface {
	GetProduct(id string) (*entities.Product, error)
}
