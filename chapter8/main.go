package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Chapter 8")

	fmt.Println()
	functions1()

	fmt.Println()
	functions2("Milk", 21.36, 14)
	fmt.Println()
	functions2("Bugatti", 2123500, 25)

	fmt.Println()
	functions3("Milk", 21.36, 0.14)
	fmt.Println()
	functions3("Bugatti", 2123500, 0.25)

	fmt.Println()
	functions4("Milk", 21.36)

	fmt.Println()
	functions5("Milk", 21.36)

	fmt.Println()
	functions6("Laptop", []string{"Apple", "Lenovo", "Samsung"})

	fmt.Println()
	functions7("Laptop", "Apple", "Lenovo", "Samsung")
	fmt.Println()
	functions7("PC")
	fmt.Println()
	suppliers := []string{"Apple", "Lenovo", "Samsung"}
	functions7("Laptop", suppliers...)

	fmt.Println()
	first := "First 1"
	second := "Second 1"
	fmt.Println("Before swap:", first, second)
	functions8(first, second, nil, nil)
	fmt.Println("After swap:", first, second)

	fmt.Println()
	first = "First 2"
	second = "Second 2"
	fmt.Println("Before swap:", first, second)
	functions8("", "", &first, &second)
	fmt.Println("After swap:", first, second)

	fmt.Println()
	fmt.Println("Tax from 1 milk carton:", functions9(21.36, 0.14))
	fmt.Println()
	fmt.Println("Tax from 1 Bugatti car:", functions9(2123500, 0.25))

	fmt.Println()
	first = "First 1"
	second = "Second 1"
	fmt.Println("Before swap:", first, second)
	first, second = functions10(first, second)
	fmt.Println("After swap:", first, second)

	fmt.Println()
	functions11()

	fmt.Println()
	functions12()

	fmt.Println()
	total, totalTax := functions13()
	fmt.Println("Total price: ", strconv.FormatFloat(total, 'f', 2, 64), "(including total tax: "+strconv.FormatFloat(totalTax, 'f', 2, 64)+")")
}
