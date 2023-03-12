package seed

import "chapter15/models"

var Products = models.ProductList{
	models.NewProduct("Milk", "Food", 23.36),
	models.NewProduct("Bread", "Food", 45.56),
	models.NewProduct("Beer", "Food", 60.34),
	models.NewProduct("Jacket", "Clothes", 160.34),
	models.NewProduct("Pants", "Clothes", 100.34),
	models.NewProduct("Hat", "Clothes", 55.34),
	models.NewProduct("Football ball", "Sports", 35.34),
	models.NewProduct("Boxing gloves", "Sports", 76.34),
	models.NewProduct("Tennis racquet", "Sports", 123.34),
}
