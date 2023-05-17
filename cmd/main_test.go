package main

import (
	"bytes"
	"clean-sales/internal/dtos"
	testfixture "clean-sales/testFixture"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenDuplicatedItems_ThenShouldNotCreateOrder(t *testing.T) {
	t.SkipNow()
	bodyReq, err := json.Marshal(testfixture.DuplicatedItem)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:9090/checkout", bytes.NewBuffer(bodyReq))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	bodyRes, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	stringRes := strings.TrimSuffix(string(bodyRes), "\n")
	body, _ := strconv.Unquote(stringRes)
	assert.Equal(t, "duplicated product", body)
}

func TestGivenNegativeItemQuantity_ThenShouldNotCreateOrder(t *testing.T) {
	t.SkipNow()
	bodyReq, err := json.Marshal(testfixture.NegativeItemQuantity)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:9090/checkout", bytes.NewBuffer(bodyReq))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	bodyRes, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	stringRes := strings.TrimSuffix(string(bodyRes), "\n")
	body, _ := strconv.Unquote(stringRes)
	assert.Equal(t, "invalid quantity", body)
}

func TestGivenInvalidCPF_ThenShouldNotCreateOrder(t *testing.T) {
	t.SkipNow()
	bodyReq, err := json.Marshal(testfixture.InvalidInput)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:9090/checkout", bytes.NewBuffer(bodyReq))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	bodyRes, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	stringRes := strings.TrimSuffix(string(bodyRes), "\n")
	body, _ := strconv.Unquote(stringRes)
	assert.Equal(t, "invalid cpf", body)
}

func TestGivenValidInput_ThenShouldCreateOrder(t *testing.T) {
	t.SkipNow()
	bodyReq, err := json.Marshal(testfixture.ValidInput)
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
	t.SkipNow()
	bodyReq, err := json.Marshal(testfixture.ValidCouponInput)
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

func TestGivenFromAndTo_ThenShouldCalculateFreightAndCreateOrder(t *testing.T) {
	t.SkipNow()
	bodyReq, err := json.Marshal(testfixture.ValidInputWithFreight)
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
	assert.Equal(t, 6290.00, orderOutput.Total)
}

func TestGivenNegativeItemDimension_ThenShouldNotCreateOrder(t *testing.T) {
	t.SkipNow()
	bodyReq, err := json.Marshal(testfixture.NegativeItemDimension)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "http://localhost:9090/checkout", bytes.NewBuffer(bodyReq))
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	bodyRes, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	stringRes := strings.TrimSuffix(string(bodyRes), "\n")
	body, _ := strconv.Unquote(stringRes)
	assert.Equal(t, "invalid product dimensions", body)
}
