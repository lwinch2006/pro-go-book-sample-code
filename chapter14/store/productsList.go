package store

type ProductsList []*Product

func NewProductsList(product *Product) ProductsList {
	productsList := make(ProductsList, 1)
	productsList[0] = product
	return productsList
}

func (pl ProductsList) Subtotal() (subtotal float64) {
	for _, product := range pl {
		subtotal += product.Price
	}

	return
}

func (pl ProductsList) SubtotalViaChannel(output chan float64) {
	output <- pl.Subtotal()
}
