package constants

import "chapter22/models"

var Products = []*models.Product{
	models.NewProduct("Milk", "Food", 23.34),
	models.NewProduct("Bread", "Food", 45.56),
	models.NewProduct("Beer", "Food", 67.78),
}
