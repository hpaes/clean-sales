package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenInvalidItemQuantity_ThenShouldNotCreateItem(t *testing.T) {
	product, _ := NewProduct("1", "A", 1000, 100, 30, 10, 3)
	item, err := NewItem(product.IdProduct, product.Price, 0)
	assert.EqualError(t, err, "invalid item quantity")
	assert.Nil(t, item)
}

func TestShouldCalculateTotal(t *testing.T) {
	product, _ := NewProduct("1", "A", 1000, 100, 30, 10, 3)
	item, _ := NewItem(product.IdProduct, product.Price, 2)
	assert.Equal(t, 2000.0, item.CalculateTotal())
}
