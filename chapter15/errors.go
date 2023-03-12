package main

import (
	"chapter15/models"
	"chapter15/seed"
	"chapter15/utils"
	"fmt"
)

func Errors1() {
	fmt.Println("Errors1()")

	categories := []string{"Food", "Clothes"}

	for _, category := range categories {
		subtotal, _ := seed.Products.Subtotal(category)
		fmt.Println("Category:", category+",", "Subtotal:", utils.ToCurrency(subtotal))
	}
}

func Errors2() {
	fmt.Println("Errors2()")

	categories := []string{"Food", "Sport"}

	for _, category := range categories {
		if subtotal, err := seed.Products.Subtotal(category); err == nil {
			fmt.Println("Category:", category+",", "Subtotal:", utils.ToCurrency(subtotal))
		} else {
			fmt.Println(err.Error())
		}
	}
}

func Errors3() {
	fmt.Println("Errors3()")

	categories := []string{"Food", "Sport"}

	input := make(chan models.ProductChannelMessage)
	go seed.Products.SubtotalAsync(categories, input)

	for message := range input {
		if message.CategoryNotExistError == nil {
			fmt.Println("Category:", message.Category+",", "Subtotal:", utils.ToCurrency(message.Subtotal))
		} else {
			fmt.Println(message.CategoryNotExistError.Error())
		}
	}
}

func Errors4() {
	fmt.Println("Errors4()")

	categories := []string{"Food", "Sport"}

	for _, category := range categories {
		if subtotal, err := seed.Products.Subtotal2(category); err == nil {
			fmt.Println("Category:", category+",", "Subtotal:", utils.ToCurrency(subtotal))
		} else {
			fmt.Println(err.Error())
		}
	}
}

func Errors5() {
	fmt.Println("Errors5()")

	categories := []string{"Food", "Sport"}

	input := make(chan models.ProductChannelMessage2)
	go seed.Products.SubtotalAsync2(categories, input)

	for message := range input {
		if message.Err == nil {
			fmt.Println("Category:", message.Category+",", "Subtotal:", utils.ToCurrency(message.Subtotal))
		} else {
			fmt.Println(message.Err.Error())
		}
	}
}

func Errors6() {
	fmt.Println("Errors6()")

	categories := []string{"Food", "Sport"}

	for _, category := range categories {
		if subtotal, err := seed.Products.Subtotal3(category); err == nil {
			fmt.Println("Category:", category+",", "Subtotal:", utils.ToCurrency(subtotal))
		} else {
			fmt.Println(err.Error())
		}
	}
}

func Errors7_1() {
	if arg := recover(); arg != nil {
		if err, ok := arg.(error); ok {
			fmt.Println("Error:", err)
		} else if str, ok := arg.(string); ok {
			fmt.Println("String:", str)
		} else {
			fmt.Println("Panic recovered")
		}
	}
}

func Errors7() {
	fmt.Println("Errors7()")

	defer Errors7_1()

	categories := []string{"Food", "Sport"}

	for _, category := range categories {
		if subtotal, err := seed.Products.Subtotal3(category); err == nil {
			fmt.Println("Category:", category+",", "Subtotal:", utils.ToCurrency(subtotal))
		} else {
			panic(err)
			//panic("Category not found")
		}
	}
}

func Errors8_1() {
	if arg := recover(); arg != nil {
		if err, ok := arg.(error); ok {
			fmt.Println("Error:", err)
			panic(err)
		} else if str, ok := arg.(string); ok {
			fmt.Println("String:", str)
		} else {
			fmt.Println("Panic recovered")
		}
	}
}

func Errors8() {
	fmt.Println("Errors8()")

	defer Errors8_1()

	categories := []string{"Food", "Sport"}

	for _, category := range categories {
		if subtotal, err := seed.Products.Subtotal3(category); err == nil {
			fmt.Println("Category:", category+",", "Subtotal:", utils.ToCurrency(subtotal))
		} else {
			panic(err)
			//panic("Category not found")
		}
	}
}

func Errors9_1(output chan<- models.CategorySubtotalMessage) {
	if arg := recover(); arg != nil {
		fmt.Println("Panic recovering from:", arg)
		output <- models.CategorySubtotalMessage{
			TerminalError: arg,
		}

		close(output)
	}
}

func Errors9(output chan<- models.CategorySubtotalMessage) {
	fmt.Println("Errors9()")

	defer Errors9_1(output)

	categories := []string{"Food", "Sport"}

	input := make(chan models.ProductChannelMessage2, 5)
	go seed.Products.SubtotalAsync2(categories, input)

	for message := range input {
		if message.Err == nil {
			output <- models.CategorySubtotalMessage{
				Category: message.Category,
				Subtotal: message.Subtotal,
			}
		} else {
			panic(message.Err)
		}
	}
	close(output)
}
