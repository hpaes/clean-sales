package entities

import (
	"math"
)

func CalculateFreight(product *Product) float64 {
	freight := product.CalculateVolume() * 1000 * (product.CalculateDensity() / 100)
	return math.Max(10, freight)
}
