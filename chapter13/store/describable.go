package store

import "chapter13/store/interfaces"

type Describable interface {
	GetName() string
	GetCategory() string
	interfaces.ItemForSale
}
