package main

import (
	"fmt"
	"strconv"
)

func functions1() {
	fmt.Println("functions1()")
	price := 12.36
	tax := 25
	totalPrice := price + (price * float64(tax) / 100.0)

	fmt.Println("Price (excl. MVA):", strconv.FormatFloat(price, 'f', 2, 64), "MVA (%):", tax, "Price (incl. MVA):", strconv.FormatFloat(totalPrice, 'f', 2, 64))
}

func functions2(product string, price float64, tax int) {
	fmt.Println("functions2()")
	totalPrice := price + (price * float64(tax) / 100.0)
	fmt.Println("Product", product, "Price (excl. MVA):", strconv.FormatFloat(price, 'f', 2, 64), "MVA (%):", tax, "Price (incl. MVA):", strconv.FormatFloat(totalPrice, 'f', 2, 64))
}

func functions3(product string, price, taxRate float64) {
	fmt.Println("functions3()")
	totalPrice := price + (price * taxRate)
	fmt.Println("Product", product, "Price (excl. MVA):", strconv.FormatFloat(price, 'f', 2, 64), "MVA (%):", int(taxRate*100), "Price (incl. MVA):", strconv.FormatFloat(totalPrice, 'f', 2, 64))
}

func functions4(product string, _ float64) {
	fmt.Println("functions4()")
	fmt.Println("Function with one not used parameter")
}

func functions5(string, float64) {
	fmt.Println("functions5()")
	fmt.Println("Functions with no parameters used at all")
}

func functions6(product string, suppliers []string) {
	fmt.Println("functions6()")
	fmt.Println("There are following suppliers for the product:", product)
	for _, supplier := range suppliers {
		fmt.Println("Supplier:", supplier)
	}
}

func functions7(product string, suppliers ...string) {
	fmt.Println("functions7()")
	fmt.Println("There are following suppliers for the product:", product)

	if len(suppliers) == 0 {
		fmt.Println("Suppliers: None")
		return
	}

	for _, supplier := range suppliers {
		fmt.Println("Supplier:", supplier)
	}
}

func functions8(first, second string, firstP, secondP *string) {
	fmt.Println("functions8()")
	temp1 := first
	first = second
	second = temp1

	if firstP == nil || secondP == nil {
		return
	}

	temp2 := *firstP
	*firstP = *secondP
	*secondP = temp2
}

func functions9(price, taxRate float64) string {
	fmt.Println("functions9()")
	tax := price * taxRate
	return strconv.FormatFloat(tax, 'f', 2, 64)
}

func functions10(first, second string) (string, string) {
	fmt.Println("functions10()")
	return second, first
}

func functions11_1(price float64) float64 {
	if price <= 100.0 {
		return -1
	}

	return price * 0.25
}

func functions11() {
	fmt.Println("functions11()")
	products := map[string]float64{
		"Milk":    12.36,
		"Bugatti": 2123500,
	}

	for product, price := range products {
		tax := functions11_1(price)

		if tax == -1 {
			fmt.Println("No tax for product:", product)
		} else {
			fmt.Println("Tax for product:", product, ":", tax)
		}
	}
}

func functions12_1(price float64) (float64, bool) {
	if price <= 100.0 {
		return 0, false
	}

	return price * 0.25, true
}

func functions12() {
	fmt.Println("functions12()")
	products := map[string]float64{
		"Milk":    12.36,
		"Bugatti": 2123500,
	}

	for product, price := range products {
		if tax, isTaxable := functions12_1(price); !isTaxable {
			fmt.Println("No tax for product:", product)
		} else {
			fmt.Println("Tax for product:", product, ":", tax)
		}
	}
}

func functions13() (total, totalTax float64) {
	fmt.Println("functions13() started")
	defer fmt.Println("First defer")

	products := map[string]float64{
		"Milk":    12.36,
		"Bugatti": 2123500,
	}

	for _, price := range products {
		total += price

		if tax, isTaxable := functions12_1(price); !isTaxable {

		} else {
			total += tax
			totalTax += tax
		}
	}

	defer fmt.Println("Second defer")
	fmt.Println("functions13() finished")
	return
}
