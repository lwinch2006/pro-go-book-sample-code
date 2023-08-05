package main

import (
	"chapter13/store"
	"chapter13/store/interfaces"
	"fmt"
)

func composition1() {
	fmt.Println("composition1()")

	products := []*store.Product{
		{
			Name:     "Milk",
			Category: "Food",
		},
		store.NewProduct("Bread", "Food", 45.56),
	}

	for _, product := range products {
		fmt.Println("Product:", product.Name+",", "price:", product.PriceWithTax(0.14))
	}
}

func composition2() {
	fmt.Println("composition2()")

	boats := []*store.Boat{
		store.NewBoat("Kayak", "Watersport", 100, 2, false),
		store.NewBoat("Speedboat", "Leisure", 1000, 4, true),
	}

	for _, boat := range boats {
		fmt.Println("Product:", boat.Name+",", "price:", boat.PriceWithTax(0.14))
	}
}

func composition3() {
	fmt.Println("composition3()")

	boats := []*store.RentalBoat{
		store.NewRentalBoat("Yacht", "Leisure", 10000, 20, true, true, "Zigzag McCrack", "Donald Duck"),
	}

	for _, boat := range boats {
		fmt.Println("Product:", boat.Name+",", "captain:", boat.Captain+",", "price:", boat.PriceWithTax(0.14))
	}
}

func composition4() {
	fmt.Println("composition4()")

	deals := []*store.SpecialDeal{
		store.NewSpeciaDeal("Milk på dato", 15, "Milk", "Food", 23.36),
	}

	for _, deal := range deals {
		fmt.Println("Deal:", deal.Name+",", "price:", deal.PriceWithTax())
	}
}

func composition5() {
	fmt.Println("compostion5()")

	bundleOffer := store.NewBundleOffer("Milk på dato", 15, "Milk", "Food", 23.36)

	//fmt.Println("Bundle offer:", bundleOffer.Name+",", "price:", bundleOffer.PriceWithTax()) // Ambiguous references
	fmt.Println("Bundle offer:", bundleOffer.SpecialDeal.Name+",", "price:", bundleOffer.SpecialDeal.PriceWithTax())
}

func composition6() {
	fmt.Println("composition6()")

	products := []interfaces.ItemForSale{
		store.NewProduct("Milk", "Food", 23.36),
		store.NewBoat("Yacht", "Leisure", 1000, 20, true),
	}

	for _, product := range products {
		fmt.Println("Product price:", product.PriceWithTax(0.25))
	}
}

func composition7() {
	fmt.Println("composition7()")

	products := []interfaces.ItemForSale{
		store.NewProduct("Milk", "Food", 23.36),
		store.NewBoat("Yacht", "Leisure", 1000, 20, true),
	}

	for _, product := range products {
		switch item := product.(type) {
		case *store.Product:
			fmt.Println("Product name:", item.Name+",", "price:", item.PriceWithTax(0.14))

		case *store.Boat:
			fmt.Println("Boat name:", item.Name+",", "price:", item.PriceWithTax(0.25))

		default:
			fmt.Println("Product price:", product.PriceWithTax(0.25))
		}
	}
}

func composition8() {
	fmt.Println("composition8()")

	products := []interfaces.ItemForSale{
		store.NewProduct("Milk", "Food", 23.36),
		store.NewBoat("Yacht", "Leisure", 1000, 20, true),
	}

	for _, product := range products {
		switch item := product.(type) {
		case store.Describable:
			fmt.Println("Product name:", item.GetName()+",", "price:", item.(interfaces.ItemForSale).PriceWithTax(0.25))

		default:
			fmt.Println("Product price:", product.PriceWithTax(0.25))
		}
	}
}

func composition9() {
	fmt.Println("composition9()")

	products := []interfaces.ItemForSale{
		store.NewProduct("Milk", "Food", 23.36),
		store.NewBoat("Yacht", "Leisure", 1000, 20, true),
	}

	for _, product := range products {
		switch item := product.(type) {
		case store.Describable:
			fmt.Println("Product name:", item.GetName()+",", "price:", item.PriceWithTax(0.25))

		default:
			fmt.Println("Product price:", product.PriceWithTax(0.25))
		}
	}
}
