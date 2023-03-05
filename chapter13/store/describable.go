package store

type Describable interface {
	GetName() string
	GetCategory() string
	ItemForSale
}
