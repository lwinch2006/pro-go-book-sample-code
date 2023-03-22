package models

import "sort"

type SortedProductListByName struct{ ProductList }

func (pl SortedProductListByName) ProductListIsSorted() bool {
	return sort.IsSorted(pl)
}

func (pl SortedProductListByName) Less(i, j int) bool {
	return pl.ProductList[i].Name < pl.ProductList[j].Name
}
