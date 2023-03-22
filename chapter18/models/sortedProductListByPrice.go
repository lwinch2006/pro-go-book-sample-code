package models

import "sort"

type SortedProductListByPrice struct{ ProductList }

func (pl SortedProductListByPrice) ProductListIsSorted() bool {
	return sort.IsSorted(pl)
}

func (pl SortedProductListByPrice) Less(i, j int) bool {
	return pl.ProductList[i].Price < pl.ProductList[j].Price
}
