package models

type Purchase struct {
	*Customer
	*Product
	Total   float64
	taxRate float64
}

func NewPurchase(product *Product, customer *Customer) *Purchase {
	purchase := &Purchase{
		Customer: customer,
		Product:  product,
		taxRate:  25.0,
	}

	purchase.GetTotal()
	return purchase
}

func (p *Purchase) calcTax() float64 {
	return p.Price * p.taxRate / 100.00
}

func (p *Purchase) GetTotal() {
	p.Total = p.Price + p.calcTax()
}
