package store

var standardTaxRate = newTaxRate(0.14, 20)

type Product struct {
	Name, Category string
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) GetPrice() float64 {
	return standardTaxRate.calcTax(p)
}

func (p *Product) SetPrice(value float64) {
	p.price = value
}
