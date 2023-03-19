package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Chapter 5")

	fmt.Println()
	fmt.Println("Numbers overflow")
	overflow()

	fmt.Println()
	fmt.Println("Reminder")
	reminder()

	fmt.Println()
	fmt.Println("Increment")
	increment()

	fmt.Println()
	fmt.Println("Compare")
	compare()

	fmt.Println()
	fmt.Println("Compare pointers")
	comparePointers()

	fmt.Println()
	fmt.Println("Type conversions")
	typeConversion()

	fmt.Println()
	fmt.Println("Bool type conversions")
	boolTypeConversion()

	fmt.Println()
	fmt.Println("Float type conversions")
	floatTypeConversion()
}

func overflow() {
	var num1 = math.MaxInt64
	var num2 = math.MaxFloat64

	fmt.Println("Num1:", num1*2)
	fmt.Println("Num2:", num2*2)
	fmt.Println("IsInf:", math.IsInf(num2*2, 0))
}

func reminder() {
	var positiveRemainder = 3 % 2
	var negativeRemainder = -3 % 2
	var absValue = math.Abs(float64(negativeRemainder))

	fmt.Println("Positive reminder:", positiveRemainder)
	fmt.Println("Negative reminder:", negativeRemainder)
	fmt.Println("Absolute value of negative reminder:", absValue)
}

func increment() {
	var value = 10.2
	value++
	fmt.Println("Increment of 10.2:", value)
}

func compare() {
	num1 := 100
	const num2 = 200.00

	compResult1 := num1 == num2

	fmt.Println("Num1 (100) is equal to Num2 (200.00):", compResult1)
}

func comparePointers() {
	num1 := 100

	p1Num1 := &num1
	p2Num1 := &num1

	num2 := 100
	p3Num2 := &num2

	fmt.Println("p1Num1 is equal to p2Num1 (memory addresses):", p1Num1 == p2Num1)
	fmt.Println("p1Num1 is equal to p3Num2 (memory addresses):", p1Num1 == p3Num2)

	fmt.Println("p1Num1 is equal to p2Num1 (values):", *p1Num1 == *p2Num1)
	fmt.Println("p1Num1 is equal to p3Num2 (values):", *p1Num1 == *p3Num2)
}

func typeConversion() {
	num1 := 100
	num2 := 200.55

	fmt.Println("Sum of num1 (100) and num2 (200.55):", float64(num1)+num2)
}

func boolTypeConversion() {
	stringValue := "not true"

	if boolValue, convError := strconv.ParseBool(stringValue); convError == nil {
		fmt.Println("Successfully parsed value " + stringValue + " to " + strconv.FormatBool(boolValue))
	} else {
		fmt.Println("Failed to parse value (" + stringValue + "). Got error (" + convError.Error() + ")")
	}
}

func floatTypeConversion() {
	num := 49.95
	fmt.Println("Converted float value (49.95) to string (without precision loss, auto):", strconv.FormatFloat(num, 'f', -1, 64))
	fmt.Println("Converted float value (49.95) to string (with extra precision):", strconv.FormatFloat(num, 'f', 4, 64))
	fmt.Println("Converted float value (49.95) to string (with precision loss):", strconv.FormatFloat(num, 'f', 1, 64))
}
