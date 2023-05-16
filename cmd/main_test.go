package main

import (
	"bytes"
	"clean-sales/internal/dtos"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	validInput = dtos.CheckoutInputDto{
		Cpf: "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: 1},
			{IdProduct: "2", Quantity: 1},
			{IdProduct: "3", Quantity: 1},
		},
	}

	validCouponInput = dtos.CheckoutInputDto{
		Cpf: "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: 1},
			{IdProduct: "2", Quantity: 1},
			{IdProduct: "3", Quantity: 1},
		},
		Coupon: "CUPOM10",
	}

	negativeItemQuantity = dtos.CheckoutInputDto{
		Cpf: "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: -1},
			{IdProduct: "2", Quantity: 1},
			{IdProduct: "3", Quantity: 1},
		},
	}

	duplicatedItem = dtos.CheckoutInputDto{
		Cpf: "029.496.970-54",
		Items: []dtos.Item{
			{IdProduct: "1", Quantity: 1},
			{IdProduct: "1", Quantity: 1},
		},
	}

	invalidInput = dtos.CheckoutInputDto{
		Cpf:   "029.496.970-53",
		Items: []dtos.Item{},
	}
)

func TestGivenDuplicatedItems_ThenShouldNotCreateOrder(t *testing.T) {
	bodyReq, err := json.Marshal(duplicatedItem)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:9090/checkout", bytes.NewBuffer(bodyReq))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	bodyRes, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	body := string(bodyRes)
	assert.Equal(t, "product duplicated", strings.TrimSuffix(body, "\n"))
}

func TestGivenNegativeItemQuantity_ThenShouldNotCreateOrder(t *testing.T) {
	bodyReq, err := json.Marshal(negativeItemQuantity)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:9090/checkout", bytes.NewBuffer(bodyReq))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	bodyRes, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	body := string(bodyRes)
	assert.Equal(t, "invalid quantity", strings.TrimSuffix(body, "\n"))
}

func TestGivenInvalidCPF_ThenShouldNotCreateOrder(t *testing.T) {

	bodyReq, err := json.Marshal(invalidInput)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:9090/checkout", bytes.NewBuffer(bodyReq))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	bodyRes, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	body := string(bodyRes)
	assert.Equal(t, "invalid cpf", strings.TrimSuffix(body, "\n"))
}

func TestGivenValidInput_ThenShouldCreateOrder(t *testing.T) {
	bodyReq, err := json.Marshal(validInput)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:9090/checkout", bytes.NewBuffer(bodyReq))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	bodyRes, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var orderOutput dtos.CheckoutOutputDto
	err = json.Unmarshal(bodyRes, &orderOutput)
	assert.NoError(t, err)
	assert.Equal(t, 6030.00, orderOutput.Total)
}

func TestGivenValidCoupon_ThenShouldCreateOrderWithDiscount(t *testing.T) {
	bodyReq, err := json.Marshal(validCouponInput)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:9090/checkout", bytes.NewBuffer(bodyReq))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	bodyRes, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var orderOutput dtos.CheckoutOutputDto
	err = json.Unmarshal(bodyRes, &orderOutput)
	assert.NoError(t, err)
	assert.Equal(t, 5427.00, orderOutput.Total)
}
