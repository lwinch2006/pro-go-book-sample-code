package models

type Product struct {
	Name, Category string
	Price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) ApplyTax() float64 {
	return p.Price * 1.2
}

func (p *Product) ApplyDiscount(discount float64) float64 {
	return p.Price - (p.Price * discount)
}

func GetName(p *Product) string {
	return p.Name
}
