package testfixture

import (
	"clean-sales/internal/app/dtos"
	"fmt"
	"strconv"
	"time"
)

var (
	ValidInput = dtos.CheckoutInputDto{
		IdOrder: ValidOrderId,
		Cpf:     "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: 1},
			{IdProduct: "2", Quantity: 1},
			{IdProduct: "3", Quantity: 1},
		},
	}

	ValidInputWithFreight = dtos.CheckoutInputDto{
		IdOrder: ValidOrderId,
		Cpf:     "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: 1},
			{IdProduct: "2", Quantity: 1},
			{IdProduct: "3", Quantity: 1},
		},
		From: "01001-000",
		To:   "04321-000",
	}

	ValidCouponInput = dtos.CheckoutInputDto{
		IdOrder: ValidOrderId,
		Cpf:     "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: 1},
			{IdProduct: "2", Quantity: 1},
			{IdProduct: "3", Quantity: 1},
		},
		Coupon: "CUPOM10",
	}

	NegativeItemQuantity = dtos.CheckoutInputDto{
		IdOrder: ValidOrderId,
		Cpf:     "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: -1},
			{IdProduct: "2", Quantity: 1},
			{IdProduct: "3", Quantity: 1},
		},
	}

	DuplicatedItem = dtos.CheckoutInputDto{
		Cpf: "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: 1},
			{IdProduct: "1", Quantity: 1},
		},
	}

	NegativeItemDimension = dtos.CheckoutInputDto{
		Cpf: "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "4", Quantity: 1},
		},
	}

	InvalidInput = dtos.CheckoutInputDto{
		Cpf:   "029.496.970-53",
		Items: []dtos.Item{},
	}
	ValidOrderId    = "7d22ab11-30a1-4d3c-8cf4-e70c1c9"
	ValidDate       = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	ExpiredDate     = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	ValidCPF        = "029.496.970-54"
	InvalidCPF      = "111.111.111-11"
	ValidSequence   = fmt.Sprintf("%d%s", time.Now().Year(), fmt.Sprintf("%08s", strconv.Itoa(1)))
	InvalidSequence = fmt.Sprintf("%d%s", time.Now().Year(), fmt.Sprintf("%08s", strconv.Itoa(0)))
)
