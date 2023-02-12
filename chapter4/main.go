package main

import (
	"fmt"
	"math/rand"
)

func main() {

	const price, tax float32 = 275.00, 27.5
	const quantity, inStock = 2, true

	test := 300

	fmt.Println("Value:", rand.Int())

	fmt.Println("Total cost:", quantity*(price+tax))
	fmt.Println("In stock:", inStock)

	fmt.Println(test)
}
