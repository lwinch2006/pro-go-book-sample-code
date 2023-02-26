package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type product struct {
	name, category string
	price          float64
}

type item struct {
	name, category string
	price          float64
}

type altProduct struct {
	name       string
	categories []string
	price      float64
}

type stockLevel struct {
	product
	count int
}

type supplier struct {
	name, city string
}

type productWithSupplier struct {
	*product
	*supplier
}

func structs1() {
	fmt.Println("structs1()")
	milk := product{
		name:     "Milk",
		category: "Food",
		price:    23.36,
	}

	fmt.Println("Product:", milk)
}

func structs2() {
	fmt.Println("structs2()")
	bread := product{"Bread", "Food", 45.56}
	fmt.Println("Product:", bread)
}

func structs3() {
	fmt.Println("structs3()")
	milkStockLevel := stockLevel{
		product: product{
			name:     "Milk",
			category: "Food",
			price:    23.36,
		},
		count: 10,
	}

	fmt.Println("Product:", milkStockLevel.product)
	fmt.Println("Product stock level:", milkStockLevel.count)
}

func structs4() {
	fmt.Println("structs4()")

	product1 := product{"Bread", "Food", 45.56}
	product2 := product{"Bread", "Food", 45.56}
	product3 := product{"Bread", "Bakery", 45.56}

	fmt.Println("product1 == product2", product1 == product2)
	fmt.Println("product1 == product3", product1 == product3)
}

func structs5() {
	fmt.Println("structs5()")

	product1 := altProduct{"Bread", []string{"Food"}, 45.56}
	product2 := altProduct{"Bread", []string{"Food"}, 45.56}

	// Invalid operation: product1 == product2 (the operator == is not defined on altProduct)
	//fmt.Println("product1 == product2", product1 == product2)

	fmt.Println("Product:", product1)
	fmt.Println("Product:", product2)
}

func structs6() {
	fmt.Println("structs6()")
	bread := product{"Bread", "Food", 45.56}

	item1 := item(bread)

	fmt.Println("Item:", item1)
}

func structs7_1(productParam struct {
	name, category string
	price          float64
}) {
	fmt.Println("Product:", productParam)
}

func structs7() {
	fmt.Println("structs7()")
	bread := product{"Bread", "Food", 45.56}
	structs7_1(bread)
}

func structs8() {
	fmt.Println("structs8()")
	bread := product{"Bread", "Food", 45.56}

	var builder strings.Builder
	_ = json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  bread.name,
		ProductPrice: bread.price,
	})

	fmt.Println(builder.String())
}

func structs9() {
	fmt.Println("structs9()")

	array1 := [1]stockLevel{
		{
			product: product{"Milk", "Food", 23.36},
			count:   10,
		},
	}

	fmt.Println("Stock for product", array1[0].product.name, ":", array1[0].count)

	slice1 := []stockLevel{
		{
			product: product{"Milk", "Food", 23.36},
			count:   10,
		},
	}

	fmt.Println("Stock for product", slice1[0].product.name, ":", slice1[0].count)

	map1 := map[string]stockLevel{
		"milk": {
			product: product{"Milk", "Food", 23.36},
			count:   10,
		},
	}

	fmt.Println("Stock for product", map1["milk"].product.name, ":", map1["milk"].count)
}

func structs10() {
	fmt.Println("structs10()")
	product1 := product{"Bread", "Food", 45.56}

	product2 := product1

	product1.name = "Cake"

	fmt.Println("Product 1:", product1.name)
	fmt.Println("Product 2:", product2.name)
}

func structs11() {
	fmt.Println("structs11()")
	product1 := product{"Bread", "Food", 45.56}

	product2 := &product1

	product1.name = "Cake"

	fmt.Println("Product 1:", product1.name)
	fmt.Println("Product 2:", (*product2).name)
	fmt.Println("Product 2 (without pointer sign):", product2.name)
}

func structs12_1(productToCalc *product) {
	if productToCalc.price > 50 {
		productToCalc.price += productToCalc.price * 0.14
	}
}

func structs12() {
	fmt.Println("structs12()")
	product1 := product{"Bread", "Food", 55.56}

	fmt.Println("Product (before):", product1)
	structs12_1(&product1)
	fmt.Println("Product (after):", product1)
}

func structs13() {
	fmt.Println("structs13()")
	product1 := &product{"Bread", "Food", 55.56}

	fmt.Println("Price for product", product1.name, ":", product1.price, "(before)")
	structs12_1(product1)
	fmt.Println("Price for product", product1.name, ":", product1.price, "(after)")
}

func newProduct(name, category string, price float64) *product {
	return &product{name, category, price}
}

func structs14() {
	fmt.Println("structs14()")

	products := []*product{
		newProduct("Milk", "Food", 23.36),
		newProduct("Bread", "Food", 55.36),
	}

	for _, productItem := range products {
		fmt.Println("Price for product", productItem.name, ":", productItem.price)
	}
}

func newProductWithSupplier(productArg *product, supplier *supplier) *productWithSupplier {
	return &productWithSupplier{productArg, supplier}
}

func newEmptyProductWithSupplier() *productWithSupplier {
	return &productWithSupplier{&product{}, &supplier{}}
}

func newSupplier(name, city string) *supplier {
	return &supplier{name, city}
}

func structs15() {
	fmt.Println("structs15()")

	products := []*productWithSupplier{
		newProductWithSupplier(newProduct("Milk", "Food", 23.36), newSupplier("Tine", "Oslo")),
		newProductWithSupplier(newProduct("Bread", "Food", 55.36), newSupplier("Bakern", "Oslo")),
	}

	for _, productWithSupplierItem := range products {
		fmt.Println("Price for product", productWithSupplierItem.product.name, ":", productWithSupplierItem.price, "by", productWithSupplierItem.supplier.name)
	}
}

func structs16() {
	fmt.Println("structs15()")

	supplier1 := newSupplier("Tine", "Oslo")

	product1 := newProductWithSupplier(newProduct("Milk", "Food", 23.36), supplier1)

	product2 := *product1

	product1.product.name = "New Product Name"
	product2.supplier.name = "New Supplier Name"

	fmt.Println("Price for product", product1.product.name, ":", product1.price, "by", product1.supplier.name)
	fmt.Println("Price for product", product2.product.name, ":", product2.price, "by", product2.supplier.name)
}

func copyProductWithSupplier(source *productWithSupplier) *productWithSupplier {
	copiedProductWithSupplier := *source
	copiedProduct := *source.product
	copiedSupplier := *source.supplier

	copiedProductWithSupplier.product = &copiedProduct
	copiedProductWithSupplier.supplier = &copiedSupplier

	return &copiedProductWithSupplier
}

func structs17() {
	fmt.Println("structs15()")

	supplier1 := newSupplier("Tine", "Oslo")
	product1 := newProductWithSupplier(newProduct("Milk", "Food", 23.36), supplier1)

	product2 := copyProductWithSupplier(product1)

	product1.product.name = "New Product Name"
	product2.supplier.name = "New Supplier Name"

	fmt.Println("Price for product", product1.product.name, ":", product1.price, "by", product1.supplier.name)
	fmt.Println("Price for product", product2.product.name, ":", product2.price, "by", product2.supplier.name)
}

func structs18() {
	//product1 := &productWithSupplier{}
	product2 := newEmptyProductWithSupplier()

	// panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println("Empty product 1 price", product1.product.price)

	fmt.Println("Empty product 2 price", product2.product.price)
}
