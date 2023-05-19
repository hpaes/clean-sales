package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCalculateFreight(t *testing.T) {
	product, _ := NewProduct("1", "A", 1000, 100, 30, 10, 3)
	freight := CalculateFreight(product)
	assert.Equal(t, 30.0, freight)
}

func TestShouldCalculateMinimumFreight(t *testing.T) {
	product, _ := NewProduct("3", "C", 30, 10, 10, 10, 0.9)
	freight := CalculateFreight(product)
	assert.Equal(t, 10.0, freight)
}
