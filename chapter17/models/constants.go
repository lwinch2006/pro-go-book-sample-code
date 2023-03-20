package models

var Milk = NewProduct("Milk", "Food", 23.36)

var Products = []*Product{
	Milk,
	NewProduct("Bread", "Food", 45.56),
	NewProduct("Beer", "Food", 78.89),

	NewProduct("Jacket", "Clothes", 123.89),
	NewProduct("Socks", "Clothes", 56.89),
	NewProduct("Shirt", "Clothes", 88.89),
}
