package entities

import "errors"

type Product struct {
	IdProduct   string
	Description string
	Price       float64
	Width       float64
	Height      float64
	Length      float64
	Weight      float64
}

func NewProduct(idProduct string, description string, price float64, width float64, height float64, length float64, weight float64) (*Product, error) {
	product := &Product{
		IdProduct:   idProduct,
		Description: description,
		Price:       price,
		Width:       width,
		Height:      height,
		Length:      length,
		Weight:      weight,
	}

	if err := product.isValid(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) isValid() error {
	if p.Width <= 0 || p.Height <= 0 || p.Length <= 0 {
		return errors.New("invalid product dimensions")
	}

	if p.Weight <= 0 {
		return errors.New("invalid product weight")
	}

	if p.Price <= 0 {
		return errors.New("invalid product price")
	}

	return nil
}

func (p *Product) CalculateVolume() float64 {
	return (p.Width * p.Height * p.Length) / 100
}

func (p *Product) CalculateDensity() float64 {
	return p.Weight / p.CalculateVolume()
}
