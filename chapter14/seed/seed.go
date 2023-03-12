package seed

import (
	"chapter14/store"
	"fmt"
	"math/rand"
	"time"
)

var Products = []*store.Product{
	store.NewProduct("Milk", "Food", 23.36),
	store.NewProduct("Bread", "Food", 45.56),
	store.NewProduct("Beer", "Food", 60.34),
	store.NewProduct("Jacket", "Clothes", 160.34),
	store.NewProduct("Pants", "Clothes", 100.34),
	store.NewProduct("Hat", "Clothes", 55.34),
	store.NewProduct("Football ball", "Sports", 35.34),
	store.NewProduct("Boxing gloves", "Sports", 76.34),
	store.NewProduct("Tennis racquet", "Sports", 123.34),
}

var ProductsByCategory = make(store.ProductsByCategory)

func init() {
	for _, product := range Products {
		if _, ok := ProductsByCategory[product.Category]; ok {
			ProductsByCategory[product.Category] = append(ProductsByCategory[product.Category], product)
		} else {
			ProductsByCategory[product.Category] = store.NewProductsList(product)
		}
	}
}

var Customers = []string{"Alice", "Bob", "Charlie", "Dora"}

// DispatchOrders function sends order details through channel
// Symbol "<-" means that channel only for writing
func DispatchOrders(output chan<- store.DispatchNotification) {
	randGen := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	ordersNumber := randGen.Intn(3) + 2
	fmt.Println("Order count:", ordersNumber)

	for i := 0; i < ordersNumber; i++ {
		output <- store.DispatchNotification{
			Customer: Customers[randGen.Intn(len(Customers)-1)],
			Product:  Products[randGen.Intn(len(Products)-1)],
			Quantity: randGen.Intn(10) + 1,
		}
	}

	close(output)
}

// DispatchOrders2 function sends order details through channel
// Symbol "<-" means that channel only for writing
func DispatchOrders2(output chan<- store.DispatchNotification) {
	randGen := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	ordersNumber := randGen.Intn(5) + 5
	fmt.Println("Order count:", ordersNumber)

	for i := 0; i < ordersNumber; i++ {
		output <- store.DispatchNotification{
			Customer: Customers[randGen.Intn(len(Customers)-1)],
			Product:  Products[randGen.Intn(len(Products)-1)],
			Quantity: randGen.Intn(10) + 1,
		}
		time.Sleep(time.Millisecond * 750)
	}

	close(output)
}

func EnumerateProducts(output chan<- *store.Product) {
	for _, product := range Products[:3] {
		output <- product
	}

	close(output)
}

func EnumerateProducts2(output chan<- *store.Product) {
	for _, product := range Products {
		select {
		case output <- product:
			fmt.Println("Sending product:", product.Name)
		default:
			fmt.Println("Discarding product:", product.Name)
		}
	}

	close(output)
}

func EnumerateProducts3(output1, output2 chan<- *store.Product) {
	for _, product := range Products {
		select {
		case output1 <- product:
			fmt.Println("Sending product (via output1):", product.Name)
		case output2 <- product:
			fmt.Println("Sending product (via output2):", product.Name)
		default:
			fmt.Println("Discarding product:", product.Name)
		}
	}

	close(output1)
	close(output2)
}
