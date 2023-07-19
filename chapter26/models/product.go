package models

type Product struct {
	Id         int
	Name       string
	CategoryId int
	Price      float64
	*Category
}

func NewProduct(id int, name string, categoryId int, price float64) *Product {
	return &Product{id, name, categoryId, price, nil}
}
