package store

type SpecialDeal struct {
	Name string
	*Product
	price float64
}

func NewSpeciaDeal(name string, price float64, productName, productCategory string, productPrice float64) *SpecialDeal {
	return &SpecialDeal{name, NewProduct(productName, productCategory, productPrice), price}
}

func (sd *SpecialDeal) PriceWithTax() float64 {
	return sd.price
}
