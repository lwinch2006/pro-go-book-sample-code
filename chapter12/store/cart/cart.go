package cart

import "chapter12/store"

type Cart struct {
	CustomerName string
	Products     []*store.Product
}

func NewCart(customerName string, products []*store.Product) *Cart {
	return &Cart{customerName, products}
}

func (c *Cart) GetSubtotal() (total float64) {
	for _, product := range c.Products {
		total += product.GetPrice()
	}

	return
}
