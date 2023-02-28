package main

import (
	"fmt"
	"interfaces/models"
	"strconv"
)

type product struct {
	name, category string
	price          float64
}

func newProduct(name, category string, price float64) *product {
	return &product{name, category, price}
}

type supplier struct {
	name, city string
}

func newSupplier(name, city string) *supplier {
	return &supplier{name, city}
}

type productWithSupplier struct {
	*product
	*supplier
}

func newProductWithSupplier(productName, category string, price float64, supplierName, city string) *productWithSupplier {
	return &productWithSupplier{newProduct(productName, category, price), newSupplier(supplierName, city)}
}

func interfaces1() {
	fmt.Println("interfaces1()")
	products := []*product{
		newProduct("Milk", "Food", 23.36),
		newProduct("Bread", "Food", 45.36),
	}

	for _, productItem := range products {
		fmt.Println(*productItem)
	}
}

func printProductDetails(productItem *product) {
	fmt.Println("Name:", productItem.name+",", "Category:", productItem.category+",", "Price:", productItem.price)
}

func interfaces2() {
	fmt.Println("interfaces2()")
	products := []*product{
		newProduct("Milk", "Food", 23.36),
		newProduct("Bread", "Food", 45.36),
	}

	for _, productItem := range products {
		printProductDetails(productItem)
	}
}

func (productItem *product) printProductDetails() {
	fmt.Println("Name:", productItem.name+",", "Category:", productItem.category+",", "Price:", productItem.price)
}

func interfaces3() {
	fmt.Println("interfaces3()")
	products := []*product{
		newProduct("Milk", "Food", 23.36),
		newProduct("Bread", "Food", 45.36),
	}

	for _, productItem := range products {
		productItem.printProductDetails()
	}
}

func (productItem *product) calculateTax(taxRate, threshold float64) float64 {
	if productItem.price > threshold {
		return productItem.price + (productItem.price * taxRate)
	}

	return productItem.price
}

func (productItem *product) printProductDetailsWithTax() {
	fmt.Println("Name:", productItem.name+",", "Category:", productItem.category+",", "Price:", productItem.calculateTax(0.25, 40))
}

func interfaces4() {
	fmt.Println("interfaces4()")
	products := []*product{
		newProduct("Milk", "Food", 23.36),
		newProduct("Bread", "Food", 45.36),
	}

	for _, productItem := range products {
		productItem.printProductDetailsWithTax()
	}
}

func (productItem *product) printDetails() {
	fmt.Println("Name:", productItem.name+",", "Category:", productItem.category+",", "Price:", productItem.price)
}

func (supplierItem *supplier) printDetails() {
	fmt.Println("Name:", supplierItem.name+",", "City:", supplierItem.city)
}

func interfaces5() {
	fmt.Println("interfaces5()")
	products := []*productWithSupplier{
		newProductWithSupplier("Milk", "Food", 23.36, "Tine", "Oslo"),
		newProductWithSupplier("Bread", "Food", 45.36, "Bakkeren", "Stavanger"),
	}

	for _, productWithSupplierItem := range products {
		productWithSupplierItem.product.printDetails()
		productWithSupplierItem.supplier.printDetails()
		fmt.Println()
	}
}

func interfaces6() {
	fmt.Println("interfaces6()")

	productValue := product{
		name:     "Milk",
		category: "Food",
		price:    23.36,
	}

	productPointer := &product{
		name:     "Bread",
		category: "Food",
		price:    56.60,
	}

	// Method can be invoked on both values and pointers (if receiver of the method is pointer)
	productValue.printDetails()
	productPointer.printDetails()
}

func (productItem product) printDetails2() {
	fmt.Println("Name:", productItem.name+",", "Category:", productItem.category+",", "Price:", productItem.price)
}

func interfaces7() {
	fmt.Println("interfaces7()")

	productValue := product{
		name:     "Milk",
		category: "Food",
		price:    23.36,
	}

	productPointer := &product{
		name:     "Bread",
		category: "Food",
		price:    56.60,
	}

	// Method can be invoked on both values and pointers (also if receiver of the method is value)
	productValue.printDetails2()
	productPointer.printDetails2()

	fmt.Println()

	// Such invocation can be done with values, but not pointers, following method signature
	product.printDetails2(productValue)

	// Such invocation can be done with pointers, but not values, following method signature
	(*product).printDetails(productPointer)
}

type productList []*product

func (productListItem *productList) calculateTotalPriceByCategory() map[string]float64 {
	total := make(map[string]float64, 0)

	for _, productItem := range *productListItem {
		total[productItem.category] += productItem.price
	}

	return total
}

func interfaces8() {
	fmt.Println("interfaces8()")
	products := productList{
		newProduct("Milk", "Food", 23.36),
		newProduct("Bread", "Food", 45.36),
	}

	for _, productItem := range products {
		productItem.printDetails()
	}

	fmt.Println()

	for category, total := range products.calculateTotalPriceByCategory() {
		fmt.Println("Category:", category+",", "Total:", total)
	}
}

func interfaces9() {
	fmt.Println("interfaces9()")
	products := []*product{
		newProduct("Milk", "Food", 23.36),
		newProduct("Bread", "Food", 45.36),
	}

	for _, productItem := range products {
		productItem.printDetails()
	}

	fmt.Println()

	productListConverted := productList(products)

	for category, total := range productListConverted.calculateTotalPriceByCategory() {
		fmt.Println("Category:", category+",", "Total:", total)
	}
}

func interfaces10() {
	fmt.Println("interfaces10()")

	productItem := models.NewProduct("iPhone", "Electronics", 1232)
	serviceItem := models.NewService("Phone coverage", 12, 123)

	fmt.Println("Product:", productItem.Name+",", "Category:", productItem.Category+",", "Price:", productItem.Price)
	fmt.Println("Service:", serviceItem.Description+",", "Total price:", strconv.FormatFloat(float64(serviceItem.DurationMonths)*serviceItem.MonthlyFee, 'f', 2, 64)+",")
}

func interfaces11() {
	fmt.Println("interfaces11()")

	expenses := []models.Expense{
		models.NewProduct("iPhone", "Electronics", 1232),
		models.NewService("Phone coverage", 12, 123),
	}

	for _, expense := range expenses {
		fmt.Println("Expense name: ", expense.GetName()+",", "total cost:", expense.GetCost(true))
	}
}

func calculateTotalExpense(expenses []models.Expense) (total float64) {
	for _, expense := range expenses {
		total += expense.GetCost(true)
	}

	return
}

func interfaces12() {
	fmt.Println("interfaces12()")

	expenses := []models.Expense{
		models.NewProduct("iPhone", "Electronics", 1232),
		models.NewService("Phone coverage", 12, 123),
	}

	for _, expense := range expenses {
		fmt.Println("Expense name: ", expense.GetName()+",", "total cost:", expense.GetCost(true))
	}

	fmt.Println("Total expenses cost:", calculateTotalExpense(expenses))
}

func interfaces13() {
	fmt.Println("interfaces13()")

	account := models.NewAccount(123)

	account.AddExpense(models.NewProduct("iPhone", "Electronics", 1232))
	account.AddExpense(models.NewService("Phone coverage", 12, 123))

	fmt.Println("Account number:", account.Number)
	for _, expense := range account.Expenses {
		fmt.Println("Expense name: ", expense.GetName()+",", "total cost:", expense.GetCost(true))
	}

	fmt.Println("Total expenses cost:", calculateTotalExpense(account.Expenses))
}

func interfaces14() {
	fmt.Println("interfaces14()")

	var p1 models.Expense = models.NewProduct("iPhone", "Electronics", 1232)
	var p2 models.Expense = models.NewProduct("iPhone", "Electronics", 1232)
	var s1 models.Expense = models.NewService("Phone coverage", 12, 123)
	var s2 models.Expense = models.NewService("Phone coverage", 12, 123)

	// There always be false since receiver type of method implementation is pointers, aka *ProductModel, *ServiceModel
	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("s1 == s2:", s1 == s2)
}

func interfaces15() {
	fmt.Println("interfaces15()")

	expenses := []models.Expense{
		models.NewService("Phone coverage", 12, 123),
		models.NewService("Laptop coverage", 12, 1230),
		models.NewProduct("iPhone", "Electronics", 1232),
	}

	for _, expense := range expenses {
		if service, isService := expense.(*models.ServiceModel); isService {
			fmt.Println("Service:", service.Description+",", "Total price:", strconv.FormatFloat(float64(service.DurationMonths)*service.MonthlyFee, 'f', 2, 64)+",")
		} else {
			fmt.Println("Expense name: ", expense.GetName()+",", "total cost:", expense.GetCost(true))
		}
	}

	fmt.Println("Total expenses cost:", calculateTotalExpense(expenses))
}

func interfaces16() {
	fmt.Println("interfaces16()")

	expenses := []models.Expense{
		models.NewService("Phone coverage", 12, 123),
		models.NewService("Laptop coverage", 12, 1230),
		models.NewProduct("iPhone", "Electronics", 1232),
	}

	for _, expense := range expenses {
		switch value := expense.(type) {
		case *models.ServiceModel:
			fmt.Println("Service:", value.Description+",", "Total price:", strconv.FormatFloat(float64(value.DurationMonths)*value.MonthlyFee, 'f', 2, 64)+",")
		case *models.ProductModel:
			fmt.Println("Product:", value.Name+",", "Category:", value.Category+",", "Price:", value.Price)
		default:
			fmt.Println("Expense name:", value.GetName()+",", "total cost:", value.GetCost(true))
		}
	}

	fmt.Println("Total expenses cost:", calculateTotalExpense(expenses))
}

func interfaces17() {
	fmt.Println("interfaces17()")

	emptyInterfaces := []interface{}{
		models.NewService("Phone coverage", 12, 123),
		models.NewService("Laptop coverage", 12, 1230),
		models.NewProduct("iPhone", "Electronics", 1232),
		"Hello World",
		123456,
	}

	for _, emptyInterface := range emptyInterfaces {
		switch value := emptyInterface.(type) {
		case *models.ServiceModel:
			fmt.Println("Service:", value.Description+",", "Total price:", strconv.FormatFloat(float64(value.DurationMonths)*value.MonthlyFee, 'f', 2, 64)+",")
		case *models.ProductModel:
			fmt.Println("Product:", value.Name+",", "Category:", value.Category+",", "Price:", value.Price)
		case models.Expense:
			fmt.Println("Expense name: ", value.GetName()+",", "total cost:", value.GetCost(true))
		default:
			fmt.Println("Value:", value)
		}
	}
}

func processEmptyInterface(item interface{}) {
	switch value := item.(type) {
	case *models.ServiceModel:
		fmt.Println("Service:", value.Description+",", "Total price:", strconv.FormatFloat(float64(value.DurationMonths)*value.MonthlyFee, 'f', 2, 64)+",")
	case *models.ProductModel:
		fmt.Println("Product:", value.Name+",", "Category:", value.Category+",", "Price:", value.Price)
	case models.Expense:
		fmt.Println("Expense name: ", value.GetName()+",", "total cost:", value.GetCost(true))
	default:
		fmt.Println("Value:", value)
	}
}

func interfaces18() {
	fmt.Println("interfaces18()")

	emptyInterfaces := []interface{}{
		models.NewService("Phone coverage", 12, 123),
		models.NewService("Laptop coverage", 12, 1230),
		models.NewProduct("iPhone", "Electronics", 1232),
		"Hello World",
		123456,
	}

	for _, emptyInterface := range emptyInterfaces {
		processEmptyInterface(emptyInterface)
	}
}

func processEmptyInterfaces(items ...interface{}) {
	for _, emptyInterface := range items {
		switch value := emptyInterface.(type) {
		case *models.ServiceModel:
			fmt.Println("Service:", value.Description+",", "Total price:", strconv.FormatFloat(float64(value.DurationMonths)*value.MonthlyFee, 'f', 2, 64)+",")
		case *models.ProductModel:
			fmt.Println("Product:", value.Name+",", "Category:", value.Category+",", "Price:", value.Price)
		case models.Expense:
			fmt.Println("Expense name: ", value.GetName()+",", "total cost:", value.GetCost(true))
		default:
			fmt.Println("Value:", value)
		}
	}
}

func interfaces19() {
	fmt.Println("interfaces19()")

	emptyInterfaces := []interface{}{
		models.NewService("Phone coverage", 12, 123),
		models.NewService("Laptop coverage", 12, 1230),
		models.NewProduct("iPhone", "Electronics", 1232),
		"Hello World",
		123456,
	}

	processEmptyInterfaces(emptyInterfaces...)
}
