package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Chapter 14")

	fmt.Println()
	Goroutine1()

	fmt.Println()
	Goroutine2()

	fmt.Println()
	//Goroutine3()

	fmt.Println()
	Goroutine4()

	// Here it is used unbuffered channel so that function Coroutine1Async2() is blocked
	// and waits to value be received from channel in order to continue execution
	fmt.Println()
	task := make(chan int)
	go Goroutine1Async(task)
	//time.Sleep(time.Second * 5)
	fmt.Println("This goes first before any output from Coroutine1Async(task)")
	var _ = <-task
	//time.Sleep(time.Second * 1)

	fmt.Println()
	Goroutine5()

	// Here it is used buffered channel so that function Coroutine1Async2() not blocked
	// and finished before timeout expires
	fmt.Println()
	task2 := make(chan int, 1)
	go Goroutine1Async2(task2)
	time.Sleep(time.Second * 1)
	fmt.Println("This goes first before any output from Coroutine1Async2(task)")
	var _ = <-task2
	//time.Sleep(time.Second * 1)

	fmt.Println()
	Goroutine6()

	fmt.Println()
	Goroutine7()

	fmt.Println()
	Goroutine8()

	fmt.Println()
	Goroutine9()

	fmt.Println()
	Goroutine10()

	fmt.Println()
	Goroutine11()

	fmt.Println()
	Goroutine12()

	fmt.Println()
	Goroutine13()

	fmt.Println()
	Goroutine14()
}
