package entities

import (
	testfixture "clean-sales/testFixture"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	idOrder  = uuid.NewString()
	prod1, _ = NewProduct("1", "A", 1000, 100, 30, 10, 3)
	prod2, _ = NewProduct("2", "B", 5000, 50, 50, 50, 22)
	prod3, _ = NewProduct("3", "C", 30, 10, 10, 10, 0.9)
)

func TestGivenInvalidCPF_ThenShouldNotCreateOrder(t *testing.T) {
	order, err := NewOrder(idOrder, testfixture.InvalidCPF, 0)
	assert.EqualError(t, err, "could not create order: invalid cpf")
	assert.Nil(t, order)
}

func TestGivenNoItems_ShouldCreateEmptyOrder(t *testing.T) {

	order, err := NewOrder(idOrder, testfixture.ValidCPF, 0)
	assert.NoError(t, err)
	assert.NotNil(t, order)

	assert.Equal(t, 0.0, order.GetTotal())
}

func TestGivenItems_ThenShouldCreateOrder(t *testing.T) {
	order, err := NewOrder(idOrder, testfixture.ValidCPF, 0)
	assert.NoError(t, err)
	assert.NotNil(t, order)

	order.AddItem(prod1, 1)
	order.AddItem(prod2, 1)
	order.AddItem(prod3, 1)
	assert.Equal(t, 6030.0, order.GetTotal())
}

func TestGivenDuplicatedItems_ThenShouldNotCreateOrder(t *testing.T) {
	order, err := NewOrder(idOrder, testfixture.ValidCPF, 0)
	assert.NoError(t, err)
	assert.NotNil(t, order)

	order.AddItem(prod1, 1)
	assert.EqualError(t, order.AddItem(prod1, 1), "could not create order: duplicated item")
}

func TestGivenItemsAndCoupon_ThenShouldCalculateDiscount(t *testing.T) {
	order, err := NewOrder(idOrder, testfixture.ValidCPF, 0)
	assert.NoError(t, err)
	assert.NotNil(t, order)

	order.AddItem(prod1, 1)
	order.AddItem(prod2, 1)
	order.AddItem(prod3, 1)

	coupon := NewCoupon("VALE20", 20.0, testfixture.ValidDate)

	order.AddCoupon(coupon)
	assert.Equal(t, 4824.0, order.GetTotal())
}

func TestGivenValidInput_ShouldCreateOrderAndGenerateCode(t *testing.T) {
	order, err := NewOrder(idOrder, testfixture.ValidCPF, 0)
	assert.NoError(t, err)
	assert.NotNil(t, order)

	assert.Equal(t, testfixture.InvalidSequence, order.Code)
}
