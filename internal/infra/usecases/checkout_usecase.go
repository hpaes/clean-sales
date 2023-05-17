package usecases

import (
	"clean-sales/internal/dtos"
)

type CheckoutUseCase interface {
	Execute(input *dtos.CheckoutInputDto) (*dtos.CheckoutOutputDto, error)
}
