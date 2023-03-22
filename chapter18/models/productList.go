package models

import (
	"fmt"
)

type ProductList []*Product

func (pl ProductList) Len() int {
	return len(pl)
}

func (pl ProductList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

func (pl ProductList) String() (output string) {
	for i := 0; i < len(pl); i++ {
		output = fmt.Sprint(output, " ", *pl[i])
	}

	return
}
