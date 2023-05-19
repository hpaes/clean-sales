package entities

import "errors"

type Item struct {
	IdProduct string
	Price     float64
	Quantity  int
}

func NewItem(idProduct string, price float64, quantity int) (*Item, error) {
	if quantity <= 0 {
		return nil, errors.New("invalid item quantity")
	}

	return &Item{
		IdProduct: idProduct,
		Price:     price,
		Quantity:  quantity,
	}, nil
}

func (i *Item) CalculateTotal() float64 {
	return i.Price * float64(i.Quantity)
}
