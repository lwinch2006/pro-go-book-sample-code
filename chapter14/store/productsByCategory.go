package store

import (
	"chapter14/utils"
	"fmt"
	"time"
)

type ProductsByCategory map[string]ProductsList

func (pc ProductsByCategory) PrintTotal2() {
	total := 0.00

	for category, products := range pc {
		subtotal := products.Subtotal()
		total += subtotal

		fmt.Println(category+":", utils.ToCurrency(subtotal))
	}
	fmt.Println("Total:", utils.ToCurrency(total))
}

func (pc ProductsByCategory) PrintTotal3() {
	total := 0.00

	for category := range pc {
		go pc.PrintSubtotal3(category)
	}
	fmt.Println("Total:", utils.ToCurrency(total))
}

func (pc ProductsByCategory) PrintSubtotal3(category string) {
	fmt.Println(category+":", utils.ToCurrency(pc[category].Subtotal()))
	time.Sleep(time.Millisecond * 100)
}

func (pc ProductsByCategory) PrintTotal4() {
	var channel chan float64 = make(chan float64)
	total := 0.00

	for category := range pc {
		go pc.PrintSubtotal4(category, channel)
	}

	for i := 0; i < len(pc); i++ {
		total += <-channel
	}

	fmt.Println("Total:", utils.ToCurrency(total))
}

func (pc ProductsByCategory) PrintSubtotal4(category string, output chan float64) {
	subtotal := pc[category].Subtotal()
	fmt.Println(category+":", utils.ToCurrency(subtotal))
	output <- subtotal
}

func (pc ProductsByCategory) PrintTotal5() {
	var channel chan float64 = make(chan float64, 2)
	total := 0.00

	for category := range pc {
		go pc.PrintSubtotal5(category, channel)
	}

	for i := 0; i < len(pc); i++ {
		total += <-channel
	}

	fmt.Println("Total:", utils.ToCurrency(total))
}

func (pc ProductsByCategory) PrintSubtotal5(category string, output chan float64) {
	subtotal := pc[category].Subtotal()
	fmt.Println(category+":", utils.ToCurrency(subtotal))
	output <- subtotal
}

func (pc ProductsByCategory) PrintTotal6() {
	var channel chan float64 = make(chan float64, 2)
	total := 0.00

	for category := range pc {
		go pc.PrintSubtotal6(category, channel)
	}

	time.Sleep(time.Second * 1)

	for i := 0; i < len(pc); i++ {
		fmt.Println("Number of values stored:", len(channel), "Capacity:", cap(channel))
		total += <-channel
	}

	fmt.Println("Total:", utils.ToCurrency(total))
}

func (pc ProductsByCategory) PrintSubtotal6(category string, output chan float64) {
	subtotal := pc[category].Subtotal()
	fmt.Println(category+":", utils.ToCurrency(subtotal))
	output <- subtotal
}
