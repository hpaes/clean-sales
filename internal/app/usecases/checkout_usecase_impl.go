package usecases

import (
	"clean-sales/internal/app/dtos"
	"clean-sales/internal/app/entities"
	"clean-sales/internal/infra/repositories"
)

type CheckoutUseCase interface {
	Execute(input *dtos.CheckoutInputDto) (*dtos.CheckoutOutputDto, error)
}

type CheckoutUseCaseImpl struct {
	productRepository repositories.ProductRepository
	couponRepository  repositories.CouponRepository
	orderRepository   repositories.OrderRepository
}

func NewCheckoutUseCaseImpl(
	productRepository repositories.ProductRepository,
	couponRepository repositories.CouponRepository,
	orderRepository repositories.OrderRepository) *CheckoutUseCaseImpl {
	return &CheckoutUseCaseImpl{
		productRepository: productRepository,
		couponRepository:  couponRepository,
		orderRepository:   orderRepository,
	}
}

// Execute implements CheckoutUseCase
func (c *CheckoutUseCaseImpl) Execute(input *dtos.CheckoutInputDto) (*dtos.CheckoutOutputDto, error) {
	var output dtos.CheckoutOutputDto

	sequence, err := c.orderRepository.CountOrder()
	if err != nil {
		return nil, err
	}

	order, err := entities.NewOrder(input.IdOrder, input.Cpf, sequence)
	if err != nil {
		return nil, err
	}

	if (len(input.Items)) > 0 {
		for _, item := range input.Items {
			product, err := c.productRepository.GetProduct(item.IdProduct)
			if err != nil {
				return nil, err
			}

			err = order.AddItem(product, item.Quantity)
			if err != nil {
				return nil, err
			}

			if input.From != "" || input.To != "" {
				order.Freight += entities.CalculateFreight(product) * float64(item.Quantity)
			}
		}
	}

	if input.Coupon != "" {
		coupon, err := c.couponRepository.GetCoupon(input.Coupon)
		if err != nil {
			return nil, err
		}

		order.AddCoupon(coupon)
	}

	c.orderRepository.SaveOrder(order)

	output.Freight = order.Freight
	output.Total = order.GetTotal()
	return &output, nil
}
