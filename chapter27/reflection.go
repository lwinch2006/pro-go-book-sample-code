package main

import (
	"chapter27/data"
	"chapter27/models"
	"chapter27/utils"
	"fmt"
	"reflect"
	"strings"
)

func printDetails1(args ...*models.Product) {
	for _, arg := range args {
		utils.Printfln("Product %v of category %v with price %.2f", arg.Name, arg.Category, arg.Price)
	}
}

func printDetails2(args ...interface{}) {
	for _, arg := range args {
		switch value := arg.(type) {
		case *models.Product:
			utils.Printfln("Product %v of category %v with price %.2f", value.Name, value.Category, value.Price)
		case *models.Customer:
			utils.Printfln("Customer %v from city %v", value.Name, value.City)
		}
	}
}

func printDetails3(args ...interface{}) {
	for _, arg := range args {
		argDetails := []string{}
		argType := reflect.TypeOf(arg)
		argValue := reflect.ValueOf(arg)

		if argType.Kind() == reflect.Pointer {
			argType = argType.Elem()
			argValue = argValue.Elem()
		}

		if argType.Kind() == reflect.Struct {
			for i := 0; i < argType.NumField(); i++ {
				argFieldName := argType.Field(i).Name
				argFieldValue := argValue.Field(i)
				argDetails = append(argDetails, fmt.Sprintf("%v: %v", argFieldName, argFieldValue))
			}

			utils.Printfln("%v: {%v}", argType.Name(), strings.Join(argDetails, ", "))
		} else {
			utils.Printfln("%v: %v", argType.Name(), argValue)
		}

	}
}

func printDetails4(args ...interface{}) {
	for _, arg := range args {
		argType := reflect.TypeOf(arg)
		argValue := reflect.ValueOf(arg)

		if argType.Kind() == reflect.Pointer {
			argType = argType.Elem()
			argValue = argValue.Elem()
		}

		pkgPath := argType.PkgPath()
		if pkgPath == "" {
			pkgPath = "(build-in)"
		}

		utils.Printfln("Name: %v, Kind: %v, Pkg path: %v", argType.Name(), argType.Kind(), pkgPath)
	}
}

func printDetails5(args ...interface{}) {
	floatPointerType := reflect.TypeOf((*float64)(nil))
	bytesArrayType := reflect.TypeOf((*[]byte)(nil))

	for _, arg := range args {
		argType := reflect.TypeOf(arg)
		argValue := reflect.ValueOf(arg)

		switch argType.Kind() {
		case reflect.Bool:
			utils.Printfln("Bool: %v", argValue.Bool())
		case reflect.Float32, reflect.Float64:
			utils.Printfln("Float: %v", argValue.Float())
		case reflect.Int:
			utils.Printfln("Int: %v", argValue.Int())
		case reflect.String:
			utils.Printfln("String: %v", argValue.String())
		case reflect.Pointer:
			if argType == floatPointerType {
				utils.Printfln("Pointer to float: %v", argValue.Elem().Float())
				break
			} else if argType == bytesArrayType {
				utils.Printfln("Pointer to bytes array: %v", argValue.Elem().Bytes())
				break
			}

			argType = argType.Elem()
			switch argType.Kind() {
			case reflect.Int:
				utils.Printfln("Pointer to integer: %v", argValue.Elem().Int())
			}
		default:
			utils.Printfln("Other: %v", argValue.String())
		}
	}
}

func printDetails6(index int, arg interface{}) (result interface{}) {
	argType := reflect.TypeOf(arg)
	argValue := reflect.ValueOf(arg)

	if argType.Kind() == reflect.Slice {
		result = argValue.Index(index).Interface()
	}

	return
}

func printDetails7(arg interface{}) {
	argValue := reflect.ValueOf(arg)

	if argValue.Kind() == reflect.Pointer {
		argValue = argValue.Elem()
	}

	if !argValue.CanSet() {
		utils.Printfln("Cannot set %v: %v", argValue.Kind(), argValue)
		return
	}

	switch argValue.Kind() {
	case reflect.Int:
		argValue.SetInt(argValue.Int() + 1)
	case reflect.String:
		argValue.SetString(strings.ToUpper(argValue.String()))
	}

	utils.Printfln("New value %v: %v", argValue.Kind(), argValue)
}

func printDetails8(source interface{}, target interface{}) {
	var oldTargetValue interface{}
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target)

	if sourceValue.Kind() == reflect.Pointer {
		sourceValue = sourceValue.Elem()
	}

	if targetValue.Kind() == reflect.Pointer {
		targetValue = targetValue.Elem()
	}

	if !targetValue.CanSet() || sourceValue.Kind() != targetValue.Kind() {
		utils.Printfln("Cannot change %v to %v since either target is not pointer or source and target have different types", sourceValue, targetValue)
		return
	}

	oldTargetValue = targetValue.Interface()

	targetValue.Set(sourceValue)
	utils.Printfln("Successfully changed %v to %v", oldTargetValue, targetValue)
}

func printDetails9(element interface{}, array interface{}) (found bool) {
	elementValue := reflect.ValueOf(element)
	arrayValue := reflect.ValueOf(array)

	if arrayValue.Kind() != reflect.Slice ||
		!arrayValue.Comparable() ||
		!elementValue.Comparable() {
		return
	}

	for i := 0; i < arrayValue.Len(); i++ {
		if arrayValue.Index(i).Interface() == element {
			found = true
			return
		}
	}

	return
}

func printDetails10(element interface{}, array interface{}) (found bool) {
	arrayValue := reflect.ValueOf(array)

	if arrayValue.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < arrayValue.Len(); i++ {
		if reflect.DeepEqual(arrayValue.Index(i).Interface(), element) {
			found = true
			return
		}
	}

	return
}

func printDetails11(source, target interface{}) (result interface{}, converted bool) {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target)

	if !sourceValue.Type().ConvertibleTo(targetValue.Type()) {
		result = sourceValue
		return
	}

	result = sourceValue.Convert(targetValue.Type()).Interface()
	converted = true
	return
}

func printDetails12(source, target interface{}) (result interface{}, converted bool) {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target)

	result = sourceValue

	if !sourceValue.Type().ConvertibleTo(targetValue.Type()) {
		return
	}

	switch sourceValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if targetValue.OverflowInt(sourceValue.Int()) {
			utils.Printfln("Overflow int")
			return
		}
	case reflect.Float32, reflect.Float64:
		if targetValue.OverflowFloat(sourceValue.Float()) {
			utils.Printfln("Overflow float")
			return
		}
	}

	result = sourceValue.Convert(targetValue.Type()).Interface()
	converted = true
	return
}

func printDetails13(source interface{}, target interface{}) {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target)

	if sourceValue.Kind() == reflect.Pointer {
		sourceValue = sourceValue.Elem()
	}

	if targetValue.Kind() == reflect.Pointer {
		targetValue = targetValue.Elem()
	}

	if !sourceValue.CanSet() || !targetValue.CanSet() || sourceValue.Kind() != targetValue.Kind() {
		utils.Printfln("Cannot change %v to %v since either target is not pointer or source and target have different types", sourceValue, targetValue)
		return
	}

	tempValue := reflect.New(targetValue.Type())
	tempValue = tempValue.Elem()
	tempValue.Set(targetValue)

	targetValue.Set(sourceValue)
	sourceValue.Set(tempValue)

	utils.Printfln("Successfully swapped values: new source %v and new target %v", sourceValue, targetValue)
}

func Reflection1() {
	utils.Printfln("Reflection1()")
	printDetails1(data.Milk)
}

func Reflection2() {
	utils.Printfln("Reflection2()")
	printDetails2(data.Milk, data.JohnDoe)
}

func Reflection3() {
	utils.Printfln("Reflection3()")
	printDetails3(data.Milk, data.JohnDoe)
}

func Reflection4() {
	utils.Printfln("Reflection4()")
	printDetails4(data.Milk, data.JohnDoe, 10, true)
}

func Reflection5() {
	utils.Printfln("Reflection5()")
	number := 10
	floatNumber := 23.34
	bytesArray := &[]byte{100, 101, 102}

	printDetails5(number, true, "Alice", &number, floatNumber, *data.Milk, &floatNumber, bytesArray)
}

func Reflection6() {
	utils.Printfln("Reflection6()")

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	utils.Printfln("Selected number: %v", printDetails6(2, numbers).(int))
}

func Reflection7() {
	utils.Printfln("Reflection7()")

	number := 100
	printDetails7(number)

	str := "Alice"
	printDetails7(str)

	printDetails7(&number)
	printDetails7(&str)
}

func Reflection8() {
	utils.Printfln("Reflection8()")

	number1 := 123
	number2 := 456

	printDetails8(number1, &number2)

	str1 := "Alice"
	str2 := "Bob"
	printDetails8(str1, &str2)

	str3 := "John"
	printDetails8(str2, str3)
	printDetails8(number1, &str2)
}

func Reflection9() {
	utils.Printfln("Reflection9()")

	element := "Alice"
	array := []string{"Bob", "Alice", "John"}

	utils.Printfln("Found \"%v\" in %v: %v", element, array, printDetails9(element, array))

	array2 := [][]string{
		array,
		{"First", "Second", "Third"},
	}

	utils.Printfln("Found \"%v\" in %v: %v", array, array2, printDetails9(array, array2))
}

func Reflection10() {
	utils.Printfln("Reflection10()")

	element := "Alice"
	array := []string{"Bob", "Alice", "John"}

	utils.Printfln("Found \"%v\" in %v: %v", element, array, printDetails10(element, array))

	array2 := [][]string{
		array,
		{"First", "Second", "Third"},
	}

	utils.Printfln("Found \"%v\" in %v: %v", array, array2, printDetails10(array, array2))
}

func Reflection11() {
	utils.Printfln("Reflection11()")

	number1 := 23.34
	number2 := 23

	result, converted := printDetails11(number1, number2)
	utils.Printfln("%v converted to type of %v: %v with result %v", number1, number2, converted, result)

	result, converted = printDetails11(number2, number1)
	utils.Printfln("%v converted to type of %v: %v with result %v", number2, number1, converted, result)

	str := "Alice"
	result, converted = printDetails11(number1, str)
	utils.Printfln("%v converted to type of %v: %v with result %v", number1, str, converted, result)
}

func Reflection12() {
	utils.Printfln("Reflection12()")

	number1 := int64(100000)
	number2 := int8(23)

	result, converted := printDetails12(number1, number2)
	utils.Printfln("%v converted to type of %v: %v with result %v", number1, number2, converted, result)

	result, converted = printDetails11(number2, number1)
	utils.Printfln("%v converted to type of %v: %v with result %v", number2, number1, converted, result)
}

func Reflection13() {
	utils.Printfln("Reflection13()")

	number1 := 123
	number2 := 456

	utils.Printfln("Trying swapping source %v with target %v", number1, number2)
	printDetails13(&number1, &number2)

	str1 := "Alice"
	str2 := "Bob"

	fmt.Println()
	utils.Printfln("Trying swapping source %v with target %v", str1, str2)
	printDetails13(&str1, &str2)
}
