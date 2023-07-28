package models

type Purchase struct {
	*Customer
	*Product
	Total   float64
	taxRate float64
}

func NewPurchase(customer *Customer, product *Product, taxRate float64) (result *Purchase) {
	result = &Purchase{
		Customer: customer,
		Product:  product,
		Total:    product.Price * (100.0 + taxRate),
		taxRate:  taxRate,
	}

	return
}
