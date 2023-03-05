package store

type ItemForSale interface {
	PriceWithTax(taxRate float64) float64
}
