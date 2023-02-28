package models

type ProductModel struct {
	Name, Category string
	Price          float64
}

func NewProduct(name, category string, price float64) *ProductModel {
	return &ProductModel{name, category, price}
}

func (product *ProductModel) GetName() string {
	return product.Name
}

func (product *ProductModel) GetCost(_ bool) float64 {
	return product.Price
}
