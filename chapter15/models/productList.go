package models

import (
	"errors"
	"fmt"
)

type ProductList []*Product

func (pl ProductList) Subtotal(category string) (subtotal float64, err *CategoryNotExistError) {
	productsCount := 0

	for _, product := range pl {
		if product.Category == category {
			subtotal += product.Price
			productsCount++
		}
	}

	if productsCount == 0 {
		err = &CategoryNotExistError{category}
	}

	return
}

func (pl ProductList) SubtotalAsync(categories []string, output chan<- ProductChannelMessage) {
	for _, category := range categories {
		subtotal, err := pl.Subtotal(category)
		output <- ProductChannelMessage{
			Category:              category,
			Subtotal:              subtotal,
			CategoryNotExistError: err,
		}
	}

	close(output)
}

func (pl ProductList) Subtotal2(category string) (subtotal float64, err error) {
	productsCount := 0

	for _, product := range pl {
		if product.Category == category {
			subtotal += product.Price
			productsCount++
		}
	}

	if productsCount == 0 {
		err = errors.New("Category " + category + " does not exist in product list")
	}

	return
}

func (pl ProductList) SubtotalAsync2(categories []string, output chan<- ProductChannelMessage2) {
	for _, category := range categories {
		subtotal, err := pl.Subtotal2(category)
		output <- ProductChannelMessage2{
			Category: category,
			Subtotal: subtotal,
			Err:      err,
		}
	}

	close(output)
}

func (pl ProductList) Subtotal3(category string) (subtotal float64, err error) {
	productsCount := 0

	for _, product := range pl {
		if product.Category == category {
			subtotal += product.Price
			productsCount++
		}
	}

	if productsCount == 0 {
		err = fmt.Errorf("Category %v does not exist in product list", category)
	}

	return
}
