package models

type Product struct {
	Name, Category string
	Price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}
