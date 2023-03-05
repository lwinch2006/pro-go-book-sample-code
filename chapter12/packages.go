package main

import (
	_ "chapter12/data" // alias for packages import that won't be used directly in this go file
	. "chapter12/fmt"  // anonymous alias
	"chapter12/store/cart"
	"github.com/fatih/color"
	"strconv"
	//currencyFmt "chapter12/fmt" // alias defined for the package "chapter12/fmt"
	"chapter12/store"
	"fmt"
)

func packages1() {
	fmt.Println("packages1()")

	product1 := store.NewProduct("Milk", "Food", 23.36)

	fmt.Println("Product:", product1.Name+",", "category:", product1.Category)
}

func packages2() {
	fmt.Println("packages2()")

	product1 := store.NewProduct("Milk", "Food", 23.36)

	//fmt.Println("Product:", product1.Name+",", "price:", currencyFmt.ToCurrency(product1.GetPrice()))
	fmt.Println("Product:", product1.Name+",", "price:", ToCurrency(product1.GetPrice()))
}

func packages3() {
	fmt.Println("packages3()")

	products := []*store.Product{
		store.NewProduct("Milk", "Food", 23.36),
		store.NewProduct("Bread", "Food", 45.56),
	}

	cartItem := cart.NewCart("Daisy", products)

	fmt.Println("Customer:", cartItem.CustomerName)
	fmt.Println("Product list:")
	for index, product := range products {
		fmt.Println(strconv.Itoa(index+1)+":", product.Name)
	}
	fmt.Println("Total:", ToCurrency(cartItem.GetSubtotal()))
}

func packages4() {
	fmt.Println("packages4()")

	product1 := store.NewProduct("Milk", "Food", 33.00)

	//fmt.Println("Product:", product1.Name+",", "price:", currencyFmt.ToCurrency(product1.GetPrice()))
	fmt.Println("Product:", product1.Name+",", "price:", ToCurrency(product1.GetPrice()))
}

func packages5() {
	fmt.Println("packages5()")

	products := []*store.Product{
		store.NewProduct("Milk", "Food", 23.36),
		store.NewProduct("Bread", "Food", 45.56),
	}

	cartItem := cart.NewCart("Daisy", products)

	color.Blue("Customer: " + cartItem.CustomerName)
	color.Green("Product list:")
	for index, product := range products {
		color.Red(strconv.Itoa(index+1) + ": " + product.Name)
	}
	color.Yellow("Total: " + ToCurrency(cartItem.GetSubtotal()))
}
