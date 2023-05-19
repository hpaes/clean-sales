package repositories

import (
	"clean-sales/internal/app/entities"
	"database/sql"
	"fmt"
)

type OrderModel struct {
	IdOrder string
	Cpf     string
	Code    string
	Total   float64
	Freight float64
}

type OrderRepository interface {
	GetOrder(id string) (OrderModel, error)
	CountOrder() (int, error)
	SaveOrder(order *entities.Order) error
}

type OrderRepositoryImpl struct {
	database *sql.DB
}

func NewOrderRepositoryImpl(db *sql.DB) OrderRepository {
	return &OrderRepositoryImpl{database: db}
}

// GetOrder implements OrderRepository
func (o *OrderRepositoryImpl) GetOrder(id string) (OrderModel, error) {
	query := "SELECT * FROM orders WHERE id_order = ?"

	row := o.database.QueryRow(query, id)

	var orderData OrderModel

	err := row.Scan(&orderData.IdOrder, &orderData.Cpf, &orderData.Code, &orderData.Total, &orderData.Freight)
	if err != nil {
		return OrderModel{}, fmt.Errorf("failed to scan order data: %w", err)
	}

	return orderData, nil
}

// CountOrder implements OrderRepository
func (o *OrderRepositoryImpl) CountOrder() (int, error) {
	query := "SELECT COUNT(*) FROM orders"

	row := o.database.QueryRow(query)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count order: %w", err)
	}

	return count, nil
}

// SaveOrder implements OrderRepository
func (o *OrderRepositoryImpl) SaveOrder(order *entities.Order) error {
	query := "INSERT INTO orders (id_order, cpf, code, total, freight) VALUES (?, ?, ?, ?, ?)"

	_, err := o.database.Exec(query, order.IdOrder, order.Cpf.Value, order.Code, order.GetTotal(), order.Freight)
	if err != nil {
		return fmt.Errorf("failed to save order: %w", err)
	}

	return nil
}
