package dtos

type CheckoutOutputDto struct {
	Total float64 `json:"total"`
}

type Item struct {
	IdProduct string `json:"id_product"`
	Quantity  int    `json:"quantity"`
}

type CheckoutInputDto struct {
	Cpf    string `json:"cpf"`
	Items  []Item `json:"items"`
	Coupon string `json:"coupon"`
}
