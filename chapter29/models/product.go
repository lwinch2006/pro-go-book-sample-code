package models

import "fmt"

type Product struct {
	Name, Category string
	Price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) unexportedMethod() {}

func (p *Product) GetAmount() string {
	return fmt.Sprintf("%.2f", p.Price)
}

func (p *Product) currencyName() string {
	return "USD"
}
