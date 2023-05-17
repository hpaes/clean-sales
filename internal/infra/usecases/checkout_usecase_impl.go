package usecases

import (
	"clean-sales/internal/dtos"
	"clean-sales/internal/entities"
	"clean-sales/internal/infra/repositories"
	"errors"
	"math"
)

type CheckoutUseCaseImpl struct {
	productRepository repositories.ProductRepository
	couponRepository  repositories.CouponRepository
}

func NewCheckoutUseCaseImpl(productRepository repositories.ProductRepository, couponRepository repositories.CouponRepository) *CheckoutUseCaseImpl {
	return &CheckoutUseCaseImpl{
		productRepository: productRepository,
		couponRepository:  couponRepository,
	}
}

// Execute implements CheckoutUseCase
func (c *CheckoutUseCaseImpl) Execute(input *dtos.CheckoutInputDto) (*dtos.CheckoutOutputDto, error) {
	var checkDuplicate []string
	var product *entities.Product
	var output dtos.CheckoutOutputDto

	err := entities.Validate(input.Cpf)
	if err != nil {
		return nil, err
	}

	if (len(input.Items)) > 0 {
		for _, item := range input.Items {
			if item.Quantity < 0 {
				return nil, errors.New("invalid quantity")
			}

			product, err = c.productRepository.GetProduct(item.IdProduct)
			if err != nil {
				return nil, err
			}

			for _, v := range checkDuplicate {
				if v == product.IdProduct {
					return nil, errors.New("duplicated product")
				}
			}

			checkDuplicate = append(checkDuplicate, product.IdProduct)
			output.Total += product.Price * float64(item.Quantity)

			if input.From != "" || input.To != "" {
				volume := (product.Width / 100) * (product.Height / 100) * (product.Length / 100)
				density := product.Weight / volume
				freight := volume * (density / 100) * 1000
				freight = math.Max(10, freight)
				output.Freight += freight * float64(item.Quantity)
			}
		}
	}

	if input.Coupon != "" {
		coupon, err := c.couponRepository.GetCoupon(input.Coupon)
		if err != nil {
			return nil, err
		}
		output.Total -= output.Total * (coupon.Discount / 100)
	}

	output.Total += output.Freight
	return &output, nil
}
