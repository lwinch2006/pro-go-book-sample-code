package main

import (
	"chapter15/models"
	"chapter15/utils"
	"fmt"
)

func main() {
	fmt.Println("Chapter 15")

	fmt.Println()
	Errors1()

	fmt.Println()
	Errors2()

	fmt.Println()
	Errors3()

	fmt.Println()
	Errors4()

	fmt.Println()
	Errors5()

	fmt.Println()
	Errors6()

	fmt.Println()
	Errors7()

	fmt.Println()
	//Errors8()

	fmt.Println()
	input := make(chan models.CategorySubtotalMessage)
	go Errors9(input)

	for message := range input {
		if message.TerminalError == nil {
			fmt.Println("Category:", message.Category, "Subtotal:", utils.ToCurrency(message.Subtotal))
		} else {
			switch arg := message.TerminalError.(type) {
			case error:
				fmt.Println("Terminal error with Error:", arg.Error())
			case string:
				fmt.Println("Terminal error with String:", arg)
			default:
				fmt.Println("Terminal error with Arg:", arg)
			}
		}
	}
}
