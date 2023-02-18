package main

import (
	"fmt"
	"strconv"
)

func main() {
	ifConditions()
	fmt.Println()
	ifWithInitStatement()
	fmt.Println()
	basicForLoop()
	fmt.Println()
	forLoopWithCondition()
	fmt.Println()
	forLoopWithInit()
	fmt.Println()
	forLoopWithRange()
	fmt.Println()
	switchCondition1()
	fmt.Println()
	switchCondition2()
	fmt.Println()
	loopWithGoTo()
}

func ifConditions() {
	num1 := 275 //275 or 27 or 632

	if num1 > 100 && num1 < 500 {
		scopedVar := 275
		fmt.Println("Num1 is bigger then 100 and less then 500: ", scopedVar)
	} else if num1 <= 100 {
		scopedVar := "Num1 is less or equal 100:"
		fmt.Println(scopedVar, num1)
	} else {
		scopedVar := false
		fmt.Println("Matched any condition: ", scopedVar)
	}
}

func ifWithInitStatement() {
	price := "275" //275 or 275f

	if value, err := strconv.Atoi(price); err == nil {
		fmt.Println("Parsed value:", value)
	} else {
		fmt.Println("Parse error:", err.Error())
	}
}

func basicForLoop() {
	counter := 1

	for {
		fmt.Println("Counter: ", counter)
		counter++

		if counter > 3 {
			break
		}
	}
}

func forLoopWithCondition() {
	counter := 1

	for counter <= 3 {
		fmt.Println("Counter: ", counter)
		counter++
	}
}

func forLoopWithInit() {
	for counter := 1; counter <= 3; counter++ {
		fmt.Println("Counter: ", counter)
	}
}

func forLoopWithRange() {
	arr := []int{1, 2, 3}

	for index, value := range arr {
		fmt.Println("Value at position (" + strconv.Itoa(index) + ") is (" + strconv.Itoa(value) + ")")
	}
}

func switchCondition1() {
	products := []string{"Milk", "Bread", "Beer", "Wine", "Cheese"}

	for _, value := range products {
		switch value {
		case "Milk":
			fallthrough
		case "Bread":
			fmt.Println("Common case for milk and bread")

		case "Beer", "Wine":
			fmt.Println("Common case for beer and wine")

		default:
			fmt.Println("Not recognized product:", value)
		}
	}
}

func switchCondition2() {
	products := []string{"Milk", "Bread", "Beer", "Wine", "Cheese"}

	for _, value := range products {
		switch {
		case value == "Milk":
			fallthrough
		case value == "Bread":
			fmt.Println("Common case for milk and bread")

		case value == "Beer", value == "Wine":
			fmt.Println("Common case for beer and wine")

		default:
			fmt.Println("Not recognized product:", value)
		}
	}
}

func loopWithGoTo() {
	counter := 1

target:
	fmt.Println("Counter: ", counter)
	counter++
	if counter <= 3 {
		goto target
	}
}
