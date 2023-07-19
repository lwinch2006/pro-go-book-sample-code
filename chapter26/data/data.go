package data

import "chapter26/models"

var Milk = models.NewProduct(1, "Milk", 1, 23.34)

var Products = []*models.Product{
	Milk,
	models.NewProduct(2, "Bread", 1, 45.56),
	models.NewProduct(3, "Beer", 1, 67.78),
}
