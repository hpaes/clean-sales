package usecases

import (
	"clean-sales/internal/app/dtos"
	"clean-sales/internal/infra/repositories"
)

type GetOrderUseCase interface {
	Execute(idOrder string) (*dtos.GetOrderOutputDto, error)
}

type GetOrderUseCaseImpl struct {
	orderRepository repositories.OrderRepository
}

func NewGetOrder(orderRepository repositories.OrderRepository) *GetOrderUseCaseImpl {
	return &GetOrderUseCaseImpl{orderRepository: orderRepository}
}

// Execute implements GetOrderUseCase
func (g *GetOrderUseCaseImpl) Execute(idOrder string) (*dtos.GetOrderOutputDto, error) {
	orderData, err := g.orderRepository.GetOrder(idOrder)
	if err != nil {
		return nil, err
	}

	return &dtos.GetOrderOutputDto{
		Code:  orderData.Code,
		Total: orderData.Total,
	}, nil
}
