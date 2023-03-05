package store

type Product struct {
	Name, Category string
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) PriceWithTax(taxRate float64) float64 {
	return p.price + p.price*taxRate
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetCategory() string {
	return p.Category
}
