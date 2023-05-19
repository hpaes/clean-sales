package entities

import (
	"fmt"
	"strconv"
	"time"
)

type Order struct {
	IdOrder string
	Cpf     *CPF
	Coupon  *Coupon
	Items   []*Item
	Freight float64
	Code    string
	// Total   float64
}

func NewOrder(idOrder string, cpf string, sequence int) (*Order, error) {
	validCpf, err := NewCPF(cpf)
	if err != nil {
		return nil, fmt.Errorf("could not create order: %w", err)
	}

	return &Order{
		IdOrder: idOrder,
		Cpf:     validCpf,
		Items:   []*Item{},
		Code:    createCode(sequence),
	}, nil
}

func (o *Order) AddItem(product *Product, quantity int) error {
	for _, item := range o.Items {
		if item.IdProduct == product.IdProduct {
			return fmt.Errorf("could not create order: duplicated item")
		}
	}

	item, err := NewItem(product.IdProduct, product.Price, quantity)
	if err != nil {
		return fmt.Errorf("could not create order: %w", err)
	}

	o.Items = append(o.Items, item)
	return nil
}

func (o *Order) GetTotal() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.CalculateTotal()
	}

	if o.Coupon != nil {
		total -= o.Coupon.CalculateDiscount(total)
	}

	total += o.Freight

	return total
}

func (o *Order) AddCoupon(coupon *Coupon) {
	if coupon.IsValid() {
		o.Coupon = coupon
	}
}

func createCode(sequence int) string {
	year := time.Now().Year()
	sequenceString := fmt.Sprintf("%08s", strconv.Itoa(sequence))
	return fmt.Sprintf("%d%s", year, sequenceString)
}
