package testfixture

import (
	"clean-sales/internal/dtos"
	"time"
)

var (
	ValidInput = dtos.CheckoutInputDto{
		Cpf: "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: 1},
			{IdProduct: "2", Quantity: 1},
			{IdProduct: "3", Quantity: 1},
		},
	}

	ValidInputWithFreight = dtos.CheckoutInputDto{
		Cpf: "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: 1},
			{IdProduct: "2", Quantity: 1},
			{IdProduct: "3", Quantity: 1},
		},
		From: "01001-000",
		To:   "04321-000",
	}

	ValidCouponInput = dtos.CheckoutInputDto{
		Cpf: "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: 1},
			{IdProduct: "2", Quantity: 1},
			{IdProduct: "3", Quantity: 1},
		},
		Coupon: "CUPOM10",
	}

	NegativeItemQuantity = dtos.CheckoutInputDto{
		Cpf: "029.496.970-54",
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

	ValidDate = time.Now().AddDate(0, 1, 0).Format("2006-01-02")
)
