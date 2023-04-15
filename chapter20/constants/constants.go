package constants

import "chapter20/models"

var Milk = models.NewProduct("Milk", "Food", 23.36)

var Products = []*models.Product{
	Milk,
	models.NewProduct("Bread", "Food", 45.56),
	models.NewProduct("Beer", "Food", 56.67),
	models.NewProduct("Jacket", "Clothes", 156.67),
	models.NewProduct("Jeans", "Clothes", 89.67),
	models.NewProduct("Socks", "Clothes", 39.67),
}
