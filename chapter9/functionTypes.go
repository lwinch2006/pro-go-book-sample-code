package main

import (
	"fmt"
	"strconv"
)

type TotalPriceCalculator func(float64) float64

func funcTypes1_1(price float64) float64 {
	return price + (price * 0.25)
}

func funcTypes1_2(price float64) float64 {
	return price
}

func funcTypes1_3(price float64) func(float64) float64 {
	if price > 100 {
		return funcTypes1_1
	}

	return funcTypes1_2
}

func funcTypes1_4(price float64) TotalPriceCalculator {
	if price > 100 {
		return funcTypes1_1
	}

	return funcTypes1_2
}

func funcTypes1_5(price float64) TotalPriceCalculator {
	if price > 100 {
		var result TotalPriceCalculator = func(price float64) float64 {
			return price + (price * 0.25)
		}

		return result
	}

	result := func(price float64) float64 {
		return price
	}

	return result
}

func funcTypes1() {
	fmt.Println("funcTypes1()")

	products := map[string]float64{
		"Milk":    23.36,
		"Bugatti": 2134765,
	}

	fmt.Println("Original:", products)

	for product, price := range products {
		var calcFunc func(float64) float64
		fmt.Println("Function is assigned 1: ", calcFunc != nil)
		if price > 100 {
			calcFunc = funcTypes1_1
			fmt.Println("Function is assigned 2: ", calcFunc != nil)
			fmt.Println("Total price (incl taxes) for product", product, ":", strconv.FormatFloat(calcFunc(price), 'f', 2, 64))
			continue
		}

		calcFunc = funcTypes1_2
		fmt.Println("Function is assigned 2: ", calcFunc != nil)
		fmt.Println("Total price (no taxes) for product", product, ":", calcFunc(price))
	}
}

func funcTypes2(message string, printFunc func(a ...any) (n int, err error)) {
	_, _ = printFunc("funcTypes2()")
	_, _ = printFunc(message)
}

func funcTypes3() {
	fmt.Println("funcTypes3()")

	products := map[string]float64{
		"Milk":    23.36,
		"Bugatti": 2134765,
	}

	fmt.Println("Original:", products)

	for product, price := range products {
		calcFunc := funcTypes1_3(price)

		if price > 100 {
			fmt.Println("Total price (incl taxes) for product", product, ":", strconv.FormatFloat(calcFunc(price), 'f', 2, 64))
			continue
		}

		fmt.Println("Total price (no taxes) for product", product, ":", calcFunc(price))
	}
}

func funcTypes4() {
	fmt.Println("funcTypes4()")

	products := map[string]float64{
		"Milk":    23.36,
		"Bugatti": 2134765,
	}

	fmt.Println("Original:", products)

	for product, price := range products {
		calcFunc := funcTypes1_4(price)

		if price > 100 {
			fmt.Println("Total price (incl taxes) for product", product, ":", strconv.FormatFloat(calcFunc(price), 'f', 2, 64))
			continue
		}

		fmt.Println("Total price (no taxes) for product", product, ":", calcFunc(price))
	}
}

func funcTypes5() {
	fmt.Println("funcTypes5()")

	products := map[string]float64{
		"Milk":    23.36,
		"Bugatti": 2134765,
	}

	fmt.Println("Original:", products)

	for product, price := range products {
		var calcFunc TotalPriceCalculator

		if price > 100 {
			calcFunc = funcTypes1_1
			fmt.Println("Total price (incl taxes) for product", product, ":", strconv.FormatFloat(calcFunc(price), 'f', 2, 64))
			continue
		}

		calcFunc = funcTypes1_2
		fmt.Println("Total price (no taxes) for product", product, ":", calcFunc(price))
	}
}

func funcTypes6() {
	fmt.Println("funcTypes5()")

	products := map[string]float64{
		"Milk":    23.36,
		"Bugatti": 2134765,
	}

	fmt.Println("Original:", products)

	for product, price := range products {
		calcFunc := funcTypes1_5(price)

		if price > 100 {
			fmt.Println("Total price (incl taxes) for product", product, ":", strconv.FormatFloat(calcFunc(price), 'f', 2, 64))
			continue
		}

		fmt.Println("Total price (no taxes) for product", product, ":", calcFunc(price))
	}
}

var globalVariable = false

func funcTypes7_1(threshold float64, taxRate float64) TotalPriceCalculator {
	return func(price float64) float64 {
		if globalVariable {
			return 0.00
		} else if price > threshold {
			return price + (price * taxRate)
		}

		return price
	}
}

func funcTypes7() {
	foodTaxCalculator := funcTypes7_1(50, 0.14)
	carsTaxCalculator := funcTypes7_1(100, 0.25)

	foods := map[string]float64{
		"Bread": 55.0,
		"Milk":  20.0,
	}

	cars := map[string]float64{
		"Ferrari": 1000,
		"Bugatti": 2000,
	}

	fmt.Println("Original foods:", foods)
	for product, price := range foods {
		fmt.Println("Total price for", product, ":", foodTaxCalculator(price))
	}

	fmt.Println("Original cars:", cars)
	for product, price := range cars {
		fmt.Println("Total price for", product, ":", carsTaxCalculator(price))
	}
}

func funcTypes8() {

	globalVariable = false
	foodTaxCalculator := funcTypes7_1(50, 0.14)

	globalVariable = true
	carsTaxCalculator := funcTypes7_1(100, 0.25)

	foods := map[string]float64{
		"Bread": 55.0,
		"Milk":  20.0,
	}

	cars := map[string]float64{
		"Ferrari": 1000,
		"Bugatti": 2000,
	}

	fmt.Println("Original foods:", foods)
	for product, price := range foods {
		fmt.Println("Total price for", product, ":", foodTaxCalculator(price))
	}

	fmt.Println("Original cars:", cars)
	for product, price := range cars {
		fmt.Println("Total price for", product, ":", carsTaxCalculator(price))
	}
}

func funcTypes9_1(threshold float64, taxRate float64) TotalPriceCalculator {
	fixedGlobalVariable := globalVariable
	return func(price float64) float64 {
		if fixedGlobalVariable {
			return 0.00
		} else if price > threshold {
			return price + (price * taxRate)
		}

		return price
	}
}

func funcTypes9() {

	globalVariable = false
	foodTaxCalculator := funcTypes9_1(50, 0.14)

	globalVariable = true
	carsTaxCalculator := funcTypes9_1(100, 0.25)

	foods := map[string]float64{
		"Bread": 55.0,
		"Milk":  20.0,
	}

	cars := map[string]float64{
		"Ferrari": 1000,
		"Bugatti": 2000,
	}

	fmt.Println("Original foods:", foods)
	for product, price := range foods {
		fmt.Println("Total price for", product, ":", foodTaxCalculator(price))
	}

	fmt.Println("Original cars:", cars)
	for product, price := range cars {
		fmt.Println("Total price for", product, ":", carsTaxCalculator(price))
	}
}

func funcTypes10_1(threshold float64, taxRate float64, zeroPrice *bool) TotalPriceCalculator {
	return func(price float64) float64 {
		if *zeroPrice {
			return 0.00
		} else if price > threshold {
			return price + (price * taxRate)
		}

		return price
	}
}

func funcTypes10() {

	globalVariable = false
	foodTaxCalculator := funcTypes10_1(50, 0.14, &globalVariable)

	globalVariable = true
	carsTaxCalculator := funcTypes10_1(100, 0.25, &globalVariable)

	foods := map[string]float64{
		"Bread": 55.0,
		"Milk":  20.0,
	}

	cars := map[string]float64{
		"Ferrari": 1000,
		"Bugatti": 2000,
	}

	fmt.Println("Original foods:", foods)
	for product, price := range foods {
		fmt.Println("Total price for", product, ":", foodTaxCalculator(price))
	}

	fmt.Println("Original cars:", cars)
	for product, price := range cars {
		fmt.Println("Total price for", product, ":", carsTaxCalculator(price))
	}
}
