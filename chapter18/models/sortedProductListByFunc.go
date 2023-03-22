package models

import "sort"

type SortProductFunc func(p1, p2 *Product) bool

type SortedProductListByFunc struct {
	ProductList
	SortProductFunc
}

func (pl SortedProductListByFunc) ProductListIsSorted() bool {
	return sort.IsSorted(pl)
}

func (pl SortedProductListByFunc) Less(i, j int) bool {
	return pl.SortProductFunc(pl.ProductList[i], pl.ProductList[j])
}
