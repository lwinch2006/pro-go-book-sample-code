package main

import (
	"chapter18/models"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Sort1() {
	fmt.Println("Sort1()")
	values := []int{3, 2, 6, 5, 9, 8}

	Printfln("Original values: %v", values)
	Printfln("Are values sorted: %v", sort.IntsAreSorted(values))
	sort.Ints(values)
	Printfln("Sorted values: %v", values)
	Printfln("Are values sorted: %v", sort.IntsAreSorted(values))

	valueToSearch := 7
	index := sort.SearchInts(values, valueToSearch)
	Printfln("Returned index: %v, value (%v) is found: %v", index, valueToSearch, index < len(values) && values[index] == valueToSearch)

	valueToSearch = 10
	index = sort.SearchInts(values, valueToSearch)
	Printfln("Returned index: %v, value (%v) is found: %v", index, valueToSearch, index < len(values) && values[index] == valueToSearch)
}

func Sort2() {
	fmt.Println("Sort2()")

	rand.Seed(time.Now().UnixNano())

	values := make([]float64, 5)

	for i := 0; i < len(values); i++ {
		values[i] = float64(rand.Intn(20)) + rand.Float64()
	}

	Printfln("Original values: %v", values)
	Printfln("Are values sorted: %v", sort.Float64sAreSorted(values))
	sort.Float64s(values)
	Printfln("Sorted values: %v", values)
	Printfln("Are values sorted: %v", sort.Float64sAreSorted(values))

	valueToSearch := float64(rand.Intn(20)) + rand.Float64()
	index := sort.SearchFloat64s(values, valueToSearch)
	Printfln("Returned index: %v, value (%v) is found: %v", index, valueToSearch, index < len(values) && values[index] == valueToSearch)

	valueToSearch = 30.0 + float64(rand.Intn(20)) + rand.Float64()
	index = sort.SearchFloat64s(values, valueToSearch)
	Printfln("Returned index: %v, value (%v) is found: %v", index, valueToSearch, index < len(values) && values[index] == valueToSearch)
}

func Sort3() {
	fmt.Println("Sort3()")

	values := []string{"Dave", "Charlie", "Alice", "Ester", "Bob"}

	Printfln("Original values: %v", values)
	Printfln("Are values sorted: %v", sort.StringsAreSorted(values))
	sort.Strings(values)
	Printfln("Sorted values: %v", values)
	Printfln("Are values sorted: %v", sort.StringsAreSorted(values))

	valueToSearch := "Dmitry"
	index := sort.SearchStrings(values, valueToSearch)
	Printfln("Returned index: %v, value (%v) is found: %v", index, valueToSearch, index < len(values) && values[index] == valueToSearch)

	valueToSearch = "Michael"
	index = sort.SearchStrings(values, valueToSearch)
	Printfln("Returned index: %v, value (%v) is found: %v", index, valueToSearch, index < len(values) && values[index] == valueToSearch)
}

func Sort4() {
	fmt.Println("Sort4()")

	values1 := models.SortedProductListByPrice{
		ProductList: models.ProductList{
			models.NewProduct("Bread", "Food", 45.56),
			models.NewProduct("Beer", "Food", 67.23),
			models.NewProduct("Milk", "Food", 23.36),
		},
	}

	Printfln("Original values: %v", values1)
	Printfln("Are values sorted: %v", values1.ProductListIsSorted())
	sort.Sort(values1)
	Printfln("Sorted values: %v", values1)
	Printfln("Are values sorted: %v", values1.ProductListIsSorted())

	fmt.Println()

	values2 := models.SortedProductListByName{
		ProductList: models.ProductList{
			models.NewProduct("Bread", "Food", 45.56),
			models.NewProduct("Beer", "Food", 67.23),
			models.NewProduct("Milk", "Food", 23.36),
		},
	}

	Printfln("Original values: %v", values2)
	Printfln("Are values sorted: %v", values2.ProductListIsSorted())
	sort.Sort(values2)
	Printfln("Sorted values: %v", values2)
	Printfln("Are values sorted: %v", values2.ProductListIsSorted())

	fmt.Println()

	values3 := models.SortedProductListByFunc{
		ProductList: models.ProductList{
			models.NewProduct("Bread", "Food", 45.56),
			models.NewProduct("Beer", "Food", 67.23),
			models.NewProduct("Milk", "Food", 23.36),
		},
		SortProductFunc: func(p1, p2 *models.Product) bool {
			return p1.Price > p2.Price
		},
	}

	Printfln("Original values: %v", values3)
	Printfln("Are values sorted: %v", values3.ProductListIsSorted())
	sort.Sort(values3)
	Printfln("Sorted values: %v", values3)
	Printfln("Are values sorted: %v", values3.ProductListIsSorted())
}
