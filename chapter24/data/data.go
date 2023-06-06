package data

import "chapter24/models"

var Milk = models.NewProduct("Milk", "Food", 23.34)

var Products = []*models.Product{
	Milk,
	models.NewProduct("Bread", "Food", 45.56),
	models.NewProduct("Beer", "Food", 67.78),
}
