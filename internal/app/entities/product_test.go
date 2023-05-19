package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCalculateVolume(t *testing.T) {
	product, err := NewProduct("1", "A", 1000, 100, 30, 10, 3)
	assert.Nil(t, err)
	assert.Equal(t, 300.0, product.CalculateVolume())
}

func TestShouldCalculateDensity(t *testing.T) {
	product, err := NewProduct("1", "A", 1000, 100, 30, 10, 3)
	assert.Nil(t, err)
	assert.Equal(t, 0.01, product.CalculateDensity())
}

func TestGivenInvalidDimensions_ThenShouldNotCreateProduct(t *testing.T) {
	product, err := NewProduct("1", "A", 1000, 0, 30, 10, 3)
	assert.EqualError(t, err, "invalid product dimensions")
	assert.Nil(t, product)
}

func TestGivenInvalidWeight_ThenShouldNotCreateProduct(t *testing.T) {
	product, err := NewProduct("1", "A", 1000, 100, 30, 10, 0)
	assert.EqualError(t, err, "invalid product weight")
	assert.Nil(t, product)
}

func TestGivenInvalidPrice_ThenShouldNotCreateProduct(t *testing.T) {
	product, err := NewProduct("1", "A", 0, 100, 30, 10, 3)
	assert.EqualError(t, err, "invalid product price")
	assert.Nil(t, product)
}
