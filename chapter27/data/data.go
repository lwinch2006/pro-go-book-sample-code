package data

import "chapter27/models"

var Milk = models.NewProduct("Milk", "Food", 23.34)

var Products = []*models.Product{
	Milk,
	models.NewProduct("Bread", "Food", 45.56),
	models.NewProduct("Beer", "Food", 56.67),
}

var JohnDoe = models.NewCustomer("John Doe", "New York")
