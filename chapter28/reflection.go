package main

import (
	"chapter28/data"
	"chapter28/models"
	"chapter28/utils"
	"fmt"
	"reflect"
	"strings"
)

func reflection1_1(args ...any) {
	for _, arg := range args {
		argType := reflect.TypeOf(arg)
		argValue := reflect.ValueOf(arg)

		if argType.Kind() == reflect.Pointer {
			argType = argType.Elem()
			argValue = argValue.Elem()
		}

		if argType.Kind() == reflect.Slice {
			return
		}

		utils.Printfln("%v: %v", argType.Kind(), argValue)

		if argType.Kind() == reflect.Struct {
			for i := 0; i < argType.NumField(); i++ {
				utils.Printfln("struct field %v: %v", argType.Field(i).Name, argValue.Field(i))
			}
		}
	}
}

func Reflection1() {
	utils.Printfln("Reflection1()")

	reflection1_1(data.Milk, data.JohnDoe, 10, true, 23.34)
}

func Reflection2() {
	utils.Printfln("Reflection2()")

	str := "Alice"

	type1 := reflect.TypeOf(str)
	type2 := reflect.PtrTo(reflect.TypeOf(str))
	type3 := type2.Elem()

	utils.Printfln("Type 1: %v", type1)
	utils.Printfln("Type 2: %v", type2)
	utils.Printfln("Type 3: %v", type3)
}

func reflection3_3(arg interface{}) {
	stringPointerType := reflect.TypeOf((*string)(nil))

	argValue := reflect.ValueOf(arg)
	if argValue.Type() != stringPointerType {
		return
	}

	if argValue = argValue.Elem(); argValue.CanSet() {
		argValue.SetString(strings.ToUpper(argValue.String()))
	}
}

func Reflection3() {
	utils.Printfln("Reflection3()")

	original := "Alice"

	str1 := original
	reflection3_3(str1)
	utils.Printfln("old value: %v - new value: %v", original, str1)

	str1 = original
	reflection3_3(&str1)
	utils.Printfln("old value: %v - new value: %v", original, str1)
}

func reflection4_4(element, array interface{}) bool {
	elementType := reflect.TypeOf(element)
	arrayType := reflect.TypeOf(array)

	return (arrayType.Kind() == reflect.Array || arrayType.Kind() == reflect.Slice) && arrayType.Elem() == elementType
}

func Reflection4() {
	utils.Printfln("Reflection4()")

	arr1 := [3]string{"Alice", "Bob", "John"}
	arr2 := []string{"Alice", "Bob", "John"}

	element1 := "Mary"
	element2 := "New York"
	element3 := 10

	utils.Printfln("Element (%v) fits to array (%v): %v", element1, arr1, reflection4_4(element1, arr1))
	utils.Printfln("Element (%v) fits to array (%v): %v", element2, arr2, reflection4_4(element2, arr2))
	utils.Printfln("Element (%v) fits to array (%v): %v", element3, arr2, reflection4_4(element3, arr2))
}

func reflection5_5(array interface{}, index int, replacement interface{}) {
	arrayType := reflect.TypeOf(array)
	arrayValue := reflect.ValueOf(array)
	replacementType := reflect.TypeOf(replacement)

	if arrayType.Kind() == reflect.Pointer {
		arrayType = arrayType.Elem()
		arrayValue = arrayValue.Elem()
	}

	if (arrayType.Kind() != reflect.Array &&
		arrayType.Kind() != reflect.Slice) ||
		arrayType.Elem() != replacementType ||
		index >= arrayValue.Len() {
		return
	}

	if !arrayValue.Index(index).CanSet() {
		return
	}

	replacementValue := reflect.ValueOf(replacement)
	arrayValue.Index(index).Set(replacementValue)
}

func Reflection5() {
	utils.Printfln("Reflection5()")

	element1 := "Alice"
	element2 := "Bob"
	element3 := "John"

	originalArr1 := [3]string{element1, element2, element3}
	originalArr2 := []string{element1, element2, element3}

	array1 := originalArr1
	reflection5_5(&array1, 1, "Test")
	utils.Printfln("Array after replacement: %v", array1)

	array2 := originalArr2
	reflection5_5(array2, 1, "Test")
	utils.Printfln("Array after replacement: %v", array2)
}

func reflection6_6(array interface{}, elemToFound interface{}) (result interface{}) {
	elemToFoundType := reflect.TypeOf(elemToFound)
	arrayType := reflect.TypeOf(array)
	arrayValue := reflect.ValueOf(array)

	if arrayType.Kind() != reflect.Slice ||
		arrayType.Elem() != elemToFoundType {
		return
	}

	for i := 0; i < arrayValue.Len(); i++ {
		if arrayValue.Index(i).Interface() == elemToFound {
			result = arrayValue.Slice(0, i+1)
		}
	}

	return
}

func Reflection6() {
	utils.Printfln("Reflection6()")

	arr1 := []string{"Alice", "Bob", "John"}
	arr2 := reflection6_6(arr1, "Bob")

	utils.Printfln("Original array: %v", arr1)
	utils.Printfln("New array: %v", arr2)
	utils.Printfln("New array: %v", reflection6_6(arr1, "Test"))

	fmt.Println()
	arr1[0] = "Mary"
	utils.Printfln("Original array: %v", arr1)
	utils.Printfln("New array: %v", arr2)
}

func reflection7_7(source interface{}, indexes ...int) (destination interface{}) {
	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source)

	if sourceType.Kind() != reflect.Slice {
		return
	}

	destinationValue := reflect.MakeSlice(sourceType, 0, len(indexes))

	for i := 0; i < len(indexes); i++ {
		if indexes[i] >= sourceValue.Len() {
			continue
		}

		destinationValue = reflect.Append(destinationValue, sourceValue.Index(indexes[i]))
	}

	destination = destinationValue.Interface()

	return
}

func Reflection7() {
	utils.Printfln("Reflection7()")

	arr1 := []string{"Alice", "Bob", "John"}
	arr2 := reflection7_7(arr1, 0, 2)

	utils.Printfln("Original array: %v", arr1)
	utils.Printfln("New array: %v", arr2)

	fmt.Println()
	arr1[0] = "Mary"
	utils.Printfln("Original array: %v", arr1)
	utils.Printfln("New array: %v", arr2)
}

func reflection8_8(arg interface{}) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() != reflect.Map {
		utils.Printfln("Not a map type")
		return
	}

	utils.Printfln("Map with key type (%v) and value type (%v)", argType.Key(), argType.Elem())
}

func Reflection8() {
	utils.Printfln("Reflection8()")

	map1 := make(map[string]string, 0)

	reflection8_8(10)
	reflection8_8(map1)
}

func reflection9_9(arg interface{}) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() != reflect.Map {
		utils.Printfln("Not a map type")
		return
	}

	argValue := reflect.ValueOf(arg)

	for _, keyValue := range argValue.MapKeys() {
		valueValue := argValue.MapIndex(keyValue)
		utils.Printfln("Map[%v]: %v", keyValue, valueValue)
	}
}

func Reflection9() {
	utils.Printfln("Reflection9()")

	map1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	reflection9_9(10)
	reflection9_9(map1)
}

func reflection10_10(arg interface{}) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() != reflect.Map {
		utils.Printfln("Not a map type")
		return
	}

	argValue := reflect.ValueOf(arg)
	argIter := argValue.MapRange()
	for argIter.Next() {
		utils.Printfln("Map[%v]: %v", argIter.Key(), argIter.Value())
	}
}

func Reflection10() {
	utils.Printfln("Reflection10()")

	map1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	reflection10_10(10)
	reflection10_10(map1)
}

func reflection11_11(arg interface{}, key, newValue interface{}) {
	argType := reflect.TypeOf(arg)
	keyType := reflect.TypeOf(key)
	newValueType := reflect.TypeOf(newValue)

	if argType.Kind() != reflect.Map ||
		argType.Key() != keyType ||
		argType.Elem() != newValueType {
		return
	}

	argValue := reflect.ValueOf(arg)
	keyValue := reflect.ValueOf(key)
	newValueValue := reflect.ValueOf(newValue)
	argValue.SetMapIndex(keyValue, newValueValue)
}

func reflection11_11_11(arg interface{}, key interface{}) {
	argType := reflect.TypeOf(arg)
	keyType := reflect.TypeOf(key)

	if argType.Kind() != reflect.Map ||
		argType.Key() != keyType {
		return
	}

	argValue := reflect.ValueOf(arg)
	keyValue := reflect.ValueOf(key)
	newValueValue := reflect.Value{}
	argValue.SetMapIndex(keyValue, newValueValue)
}

func Reflection11() {
	utils.Printfln("Reflection11()")

	map1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	utils.Printfln("Map: %v", map1)
	utils.Printfln("Map(len): %v", len(map1))

	delete(map1, "key1")
	utils.Printfln("Map: %v", map1)
	utils.Printfln("Map(len): %v", len(map1))

	reflection11_11(map1, "key2", "value22")
	utils.Printfln("Map: %v", map1)
	utils.Printfln("Map(len): %v", len(map1))

	reflection11_11(map1, "key3", "value3")
	utils.Printfln("Map: %v", map1)
	utils.Printfln("Map(len): %v", len(map1))

	reflection11_11_11(map1, "key4")
	utils.Printfln("Map: %v", map1)
	utils.Printfln("Map(len): %v", len(map1))

	reflection11_11_11(map1, "key3")
	utils.Printfln("Map: %v", map1)
	utils.Printfln("Map(len): %v", len(map1))
}

func reflection12_12(arg interface{}, valueFunction func(interface{}) interface{}) (result interface{}) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() != reflect.Slice {
		return
	}

	argValue := reflect.ValueOf(arg)
	resultValue := reflect.MakeMap(reflect.MapOf(argType.Elem(), argType.Elem()))
	for i := 0; i < argValue.Len(); i++ {
		resultValue.SetMapIndex(argValue.Index(i), reflect.ValueOf(valueFunction(argValue.Index(i).Interface())))
	}

	result = resultValue.Interface()
	return
}

func Reflection12() {
	utils.Printfln("Reflection11()")

	stringToUpperFunc := func(arg interface{}) (result interface{}) {
		if str, ok := arg.(string); ok {
			result = strings.ToUpper(str)
		}

		return
	}

	arr1 := []string{"Alice", "Bob"}
	utils.Printfln("Original slice: %v", arr1)

	map2 := reflection12_12(arr1, stringToUpperFunc)
	utils.Printfln("New map: %v", map2)
}

func isStructOrPointerToStruct(argType reflect.Type) (returnType reflect.Type, isStruct bool) {
	if argType.Kind() == reflect.Struct {
		returnType = argType
		isStruct = true
		return
	}

	if argType.Kind() == reflect.Pointer && argType.Elem().Kind() == reflect.Struct {
		returnType = argType.Elem()
		isStruct = true
	}

	return
}

func reflection13_13(baseIndex []int, arg interface{}) {
	if argType, isStruct := isStructOrPointerToStruct(reflect.TypeOf(arg)); isStruct {
		numFields := argType.NumField()
		if len(baseIndex) > 0 {
			if argType.FieldByIndex(baseIndex).Type.Kind() == reflect.Pointer {
				numFields = argType.FieldByIndex(baseIndex).Type.Elem().NumField()
			} else {
				numFields = argType.FieldByIndex(baseIndex).Type.NumField()
			}
		}

		for i := 0; i < numFields; i++ {
			currentIndex := append(baseIndex, i)
			fieldDescription := argType.FieldByIndex(currentIndex)
			utils.Printfln("Field %v with name (%v) of type (%v), exported(%v)", currentIndex, fieldDescription.Name, fieldDescription.Type, fieldDescription.PkgPath == "")

			if _, isStruct = isStructOrPointerToStruct(fieldDescription.Type); isStruct {
				reflection13_13(currentIndex, arg)
			}
		}
	}
}

func reflection13_13_13(baseIndex []int, argType reflect.Type) {
	for i := 0; i < argType.NumField(); i++ {
		fieldIndex := append(baseIndex, i)
		field := argType.Field(i)
		utils.Printfln("Field %v with name (%v) of type (%v), exported(%v)", fieldIndex, field.Name, field.Type, field.PkgPath == "")

		fieldType := field.Type
		if fieldType.Kind() == reflect.Pointer {
			fieldType = fieldType.Elem()
		}

		// This won't work for nested structs as index is global but fieldType, e.g. Customer, Product is calculated and is local
		if fieldType.Kind() == reflect.Struct {
			field := argType.FieldByIndex(fieldIndex)
			fieldType := field.Type
			if fieldType.Kind() == reflect.Pointer {
				fieldType = fieldType.Elem()
			}
			reflection13_13_13(fieldIndex, fieldType)
		}
	}

}

func Reflection13() {
	utils.Printfln("Reflection13()")

	reflection13_13([]int{}, models.Purchase{})

	fmt.Println()
	reflection13_13_13([]int{}, reflect.TypeOf(models.Purchase{}))
}

func reflection14_14(arg interface{}, name string) {
	if argType, isStruct := isStructOrPointerToStruct(reflect.TypeOf(arg)); isStruct {

		// Search for name is case-sensitive
		if field, found := argType.FieldByName(name); found {
			propertyPath := argType.Name()
			for i := 0; i < len(field.Index); i++ {
				index := field.Index[0 : i+1]
				propertyPath += "." + argType.FieldByIndex(index).Name
			}

			utils.Printfln("Property path: %v", propertyPath)
			utils.Printfln("Field %v with name (%v) of type (%v), exported(%v)", field.Index, field.Name, field.Type, field.PkgPath == "")
		} else {
			utils.Printfln("Field (%v) not found", name)
		}
	}
}

func Reflection14() {
	utils.Printfln("Reflection14()")

	// Search for name is case-sensitive
	reflection14_14(models.Purchase{}, "Price")
}

func reflection15_15(arg interface{}, tagName string) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)
		fieldTag1 := field.Tag.Get(tagName)
		fieldTag2, ok := field.Tag.Lookup(tagName)

		utils.Printfln("Field (%v) with tag (%v): %v", field.Name, tagName, fieldTag1)
		utils.Printfln("Field (%v) with tag (%v): %v, %v", field.Name, tagName, fieldTag2, ok)
	}
}

func Reflection15() {
	utils.Printfln("Reflection15()")

	reflection15_15(models.Person{}, "alias")
}

func reflection16_16(arg interface{}, tagName string) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)
		fieldTag1 := field.Tag.Get(tagName)
		fieldTag2, ok := field.Tag.Lookup(tagName)

		utils.Printfln("Field (%v) with tag (%v): %v", field.Name, tagName, fieldTag1)
		utils.Printfln("Field (%v) with tag (%v): %v, %v", field.Name, tagName, fieldTag2, ok)
	}
}

func Reflection16() {
	utils.Printfln("Reflection16()")

	stringType := reflect.TypeOf("string")
	structType := reflect.StructOf([]reflect.StructField{
		{Name: "Name", Type: stringType, Tag: `alias:"id"`},
		{Name: "City", Type: stringType, Tag: `alias:""`},
		{Name: "Country", Type: stringType},
	})

	reflection16_16(reflect.New(structType), "alias")
}

func reflection17_17(arg interface{}) {
	argType := reflect.TypeOf(arg)
	argValue := reflect.ValueOf(arg)

	if argType.Kind() == reflect.Pointer {
		argType = argType.Elem()
		argValue = argValue.Elem()
	}

	if argType.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < argValue.NumField(); i++ {
		fieldType := argType.Field(i)
		fieldValue := argValue.Field(i)
		utils.Printfln("Field (%v) of type (%v) with value (%v)", fieldType.Name, fieldType.Type, fieldValue)
	}
}

func Reflection17() {
	utils.Printfln("Reflection17()")
	reflection17_17(data.JohnDoe)
}

func reflection18_18(arg interface{}, newFieldValues map[string]interface{}) (result interface{}) {
	argType := reflect.TypeOf(arg)
	argValue := reflect.ValueOf(arg)

	if argType.Kind() == reflect.Pointer {
		argType = argType.Elem()
		argValue = argValue.Elem()
	}

	if argType.Kind() != reflect.Struct {
		return
	}

	for fieldName, newFieldValue := range newFieldValues {
		fieldValue := argValue.FieldByName(fieldName)
		if fieldValue.CanSet() {
			fieldValue.Set(reflect.ValueOf(newFieldValue))
		} else if fieldValue.CanAddr() {
			fieldValuePtr := fieldValue.Addr()
			if fieldValuePtr.CanSet() {
				fieldValue.Set(reflect.ValueOf(newFieldValue))
			}
		}
	}

	result = argValue.Interface()

	return
}

func Reflection18() {
	utils.Printfln("Reflection18()")

	newFieldValues := map[string]interface{}{
		"Name":  "Bob",
		"City":  "Los Angeles",
		"Total": 333.00,
	}

	purchase := models.Purchase{
		Customer: data.JohnDoe,
		Product:  data.Milk,
		Total:    100.0,
	}

	// Name won't be changed since property exists in Customer and Product and cannot be found unambiguously

	utils.Printfln("Sending struct as struct")
	utils.Printfln("Before: %v %v %v", *purchase.Customer, *purchase.Product, purchase.Total)
	result := reflection18_18(purchase, newFieldValues).(models.Purchase)
	utils.Printfln("After: %v %v %v", *result.Customer, *result.Product, result.Total)

	fmt.Println()
	utils.Printfln("Sending struct as pointer to struct")
	utils.Printfln("Before: %v %v %v", *purchase.Customer, *purchase.Product, purchase.Total)
	result = reflection18_18(&purchase, newFieldValues).(models.Purchase)
	utils.Printfln("After: %v %v %v", *result.Customer, *result.Product, result.Total)
}
