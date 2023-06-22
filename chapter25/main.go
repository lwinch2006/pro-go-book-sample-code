package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Chapter 25")

	fmt.Println()
	go HttpServer1()
	go HttpServer2()
	go HttpServer3()
	go HttpServer4()
	time.Sleep(time.Second)

	fmt.Println()
	HttpClient1()
	fmt.Println()

	fmt.Println()
	HttpClient2()
	fmt.Println()

	fmt.Println()
	HttpClient3()
	fmt.Println()

	fmt.Println()
	HttpClient4()
	fmt.Println()

	fmt.Println()
	HttpClient5()
	fmt.Println()

	fmt.Println()
	HttpClient6()
	fmt.Println()

	fmt.Println()
	HttpClient7()
	fmt.Println()

	fmt.Println()
	HttpClient8()
	fmt.Println()

	fmt.Println()
	HttpClient9()
	fmt.Println()

	fmt.Println()
	HttpClient10()
	fmt.Println()

	fmt.Println()
	HttpClient11()
	fmt.Println()

	fmt.Println()
	HttpClient12()
	fmt.Println()

	fmt.Println()
	HttpClient13()
	fmt.Println()
}
