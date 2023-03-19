package main

import (
	"chapter14/seed"
	"chapter14/store"
	"fmt"
	"time"
)

func Goroutine1() {
	fmt.Println("Goroutine1()")

	for category, products := range seed.ProductsByCategory {
		fmt.Println("Category:", category)
		for _, product := range products {
			fmt.Println("Product:", product.Name)
		}
		fmt.Println()
	}
}

func Goroutine2() {
	fmt.Println("Goroutine2()")
	seed.ProductsByCategory.PrintTotal2()
}

func Goroutine3() {
	fmt.Println("Goroutine3()")
	seed.ProductsByCategory.PrintTotal3()
	time.Sleep(time.Second * 2)
}

func Goroutine4() {
	fmt.Println("Goroutine4()")
	seed.ProductsByCategory.PrintTotal4()
}

func Goroutine1Async(output chan int) {
	fmt.Println("Goroutine1Async()")
	Goroutine1()
	output <- 0
	fmt.Println("Goroutine1Async() finishes")
}

func Goroutine5() {
	fmt.Println("Goroutine5()")
	seed.ProductsByCategory.PrintTotal5()
}

func Goroutine1Async2(output chan int) {
	fmt.Println("Goroutine1Async2()")
	Goroutine1()
	output <- 0
	fmt.Println("Goroutine1Async2() finishes")
}

func Goroutine6() {
	fmt.Println("Goroutine6()")
	seed.ProductsByCategory.PrintTotal6()
}

func Goroutine7() {
	fmt.Println("Goroutine7()")
	dispatchChannel := make(chan store.DispatchNotification, 100)
	go seed.DispatchOrders(dispatchChannel)

	for {
		if orderDetails, channelOpen := <-dispatchChannel; channelOpen {
			fmt.Println("New order to customer:", orderDetails.Customer, "product:", orderDetails.Name+",", "quantity:", orderDetails.Quantity)
		} else {
			fmt.Println("Dispatch channel has been closed")
			break
		}
	}
}

func Goroutine8() {
	fmt.Println("Goroutine8()")
	dispatchChannel := make(chan store.DispatchNotification, 100)
	go seed.DispatchOrders(dispatchChannel)

	for orderDetails := range dispatchChannel {
		fmt.Println("New order to customer:", orderDetails.Customer, "product:", orderDetails.Name+",", "quantity:", orderDetails.Quantity)
	}

	fmt.Println("Dispatch channel has been closed")
}

func goroutine9_1(receiveChannel <-chan store.DispatchNotification) {
	for orderDetails := range receiveChannel {
		fmt.Println("New order to customer:", orderDetails.Customer, "product:", orderDetails.Name+",", "quantity:", orderDetails.Quantity)
	}

	fmt.Println("Dispatch channel has been closed")
}

func Goroutine9() {
	fmt.Println("Goroutine9()")
	dispatchChannel := make(chan store.DispatchNotification, 100)
	var sendChannel chan<- store.DispatchNotification = dispatchChannel
	var receiveChannel <-chan store.DispatchNotification = dispatchChannel

	go seed.DispatchOrders(sendChannel)
	goroutine9_1(receiveChannel)
}

func Goroutine10() {
	fmt.Println("Goroutine10()")
	dispatchChannel := make(chan store.DispatchNotification, 100)
	var sendChannel chan<- store.DispatchNotification = dispatchChannel
	var receiveChannel <-chan store.DispatchNotification = dispatchChannel

	go seed.DispatchOrders(chan<- store.DispatchNotification(sendChannel))
	goroutine9_1((<-chan store.DispatchNotification)(receiveChannel))
}

func Goroutine11() {
	fmt.Println("Goroutine11()")
	dispatchChannel := make(chan store.DispatchNotification, 100)
	go seed.DispatchOrders2(dispatchChannel)

	for {
		select {
		case orderDetails, channelOpen := <-dispatchChannel:
			if channelOpen {
				fmt.Println("New order to customer:", orderDetails.Customer, "product:", orderDetails.Name+",", "quantity:", orderDetails.Quantity)
			} else {
				fmt.Println("Dispatch channel has been closed")
				goto alldone
			}

		default:
			fmt.Println("No message received yet")
			time.Sleep(time.Millisecond * 500)
		}
	}

alldone:
	fmt.Println("All orders are received")
}

func booleanOr(values ...bool) (result bool) {
	result = values[0]

	for i := 1; i < len(values); i++ {
		result = result || values[i]
	}

	return
}

func Goroutine12() {
	fmt.Println("Goroutine12()")
	ordersChannel := make(chan store.DispatchNotification, 100)
	go seed.DispatchOrders2(ordersChannel)

	productsChannel := make(chan *store.Product, 100)
	go seed.EnumerateProducts(productsChannel)

	// States can be tracked by different methods: for example
	// by decreasing number of channels (be careful so that one channel does not decrease number multiple times)
	// by states as booleans, more safe if it will be triggered multiple times
	channelsCount := 2
	channelStates := []bool{true, true}

	for {
		select {
		case orderDetails, ordersChannelOpen := <-ordersChannel:
			if ordersChannelOpen {
				fmt.Println("New order to customer:", orderDetails.Customer, "product:", orderDetails.Name+",", "quantity:", orderDetails.Quantity)
			} else {
				fmt.Println("Orders channel has been closed")
				channelsCount--
				channelStates[0] = false
				ordersChannel = nil
			}

		case product, productChannelOpen := <-productsChannel:
			if productChannelOpen {
				fmt.Println("Product:", product.Name)
			} else {
				fmt.Println("Products channel has been closed")
				channelsCount--
				channelStates[1] = false
				productsChannel = nil
			}

		default:
			//if channelsCount == 0 {
			//	goto alldone
			//}

			if !booleanOr(channelStates...) {
				goto alldone
			}

			fmt.Println("No message received yet")
			time.Sleep(time.Millisecond * 500)
		}
	}

alldone:
	fmt.Println("All orders are received")
}

func Goroutine13() {
	fmt.Println("Goroutine13()")
	productsChannel := make(chan *store.Product, 5)
	go seed.EnumerateProducts2(productsChannel)

	time.Sleep(time.Second * 1)

	for receivedProduct := range productsChannel {
		fmt.Println("Received product:", receivedProduct.Name)
	}
}

func Goroutine14() {
	fmt.Println("Goroutine14()")
	productsChannel1 := make(chan *store.Product, 5)
	productsChannel2 := make(chan *store.Product, 5)
	go seed.EnumerateProducts3(productsChannel1, productsChannel2)

	time.Sleep(time.Second * 1)

	for receivedProduct := range productsChannel1 {
		fmt.Println("Received product:", receivedProduct.Name)
	}

	for receivedProduct := range productsChannel2 {
		fmt.Println("Received product:", receivedProduct.Name)
	}
}
