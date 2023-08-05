package main

import (
	"chapter29/data"
	"chapter29/models"
	"chapter29/models/interfaces"
	"chapter29/utils"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func reflection1_1(arg interface{}) {
	argType := reflect.TypeOf(arg)
	argValue := reflect.ValueOf(arg)

	if argType.Kind() == reflect.Pointer {
		argType = argType.Elem()
		argValue = argValue.Elem()
	}

	if argType.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < argType.NumField(); i++ {
		fieldType := argType.Field(i)
		fieldValue := argValue.Field(i)
		utils.Printfln("Field (%v) of type (%v) with value (%v)", fieldType.Name, fieldType.Type, fieldValue)
	}
}

func Reflection1() {
	utils.Printfln("Reflection1()")

	purchase := models.NewPurchase(data.Milk, data.JohnDoe)
	reflection1_1(purchase)
}

func reflection2_2(arg interface{}) {
	argType := reflect.TypeOf(arg)
	//argValue := reflect.ValueOf(arg)

	if argType.Kind() != reflect.Func {
		return
	}

	utils.Printfln("Function (%v), Is Variadic (%v)", argType, argType.IsVariadic())
	utils.Printfln("Input parameters")
	for i := 0; i < argType.NumIn(); i++ {
		parameterInType := argType.In(i)
		utils.Printfln("Parameter (%v) of type (%v)", i, parameterInType)
	}

	utils.Printfln("Output parameters")
	for i := 0; i < argType.NumOut(); i++ {
		parameterOutType := argType.Out(i)
		utils.Printfln("Parameter (%v) of type (%v)", i, parameterOutType)
	}
}

func Reflection2() {
	utils.Printfln("Reflection2()")

	purchase := models.NewPurchase(data.Milk, data.JohnDoe)
	reflection2_2(purchase.GetTotal)

	fmt.Println()
	reflection2_2(models.NewPurchase)

	fmt.Println()
	reflection2_2(Find)
}

func reflection3_3(function interface{}, params ...interface{}) {
	functionValue := reflect.ValueOf(function)

	if functionValue.Type().Kind() != reflect.Func {
		return
	}

	paramsValues := make([]reflect.Value, 0, len(params))
	//paramsValues := []reflect.Value{}
	//var paramsValues []reflect.Value

	for _, param := range params {
		paramValue := reflect.ValueOf(param)
		paramsValues = append(paramsValues, paramValue)
	}

	resultValues := functionValue.Call(paramsValues)

	for _, resultValue := range resultValues {
		utils.Printfln("Result of type (%v) and value (%v)", resultValue.Type(), resultValue)
	}
}

func Reflection3() {
	utils.Printfln("Reflection3()")

	names := []string{"Alice", "Bob", "Jon"}
	searchName := "Bob"

	reflection3_3(Find, names, searchName)
}

func reflection4_4(function interface{}, params interface{}) {
	paramsValue := reflect.ValueOf(params)
	functionValue := reflect.ValueOf(function)
	functionType := reflect.TypeOf(function)

	if functionValue.Type().Kind() != reflect.Func ||
		functionType.NumIn() != 1 ||
		functionType.NumOut() != 1 ||
		paramsValue.Type().Kind() != reflect.Slice ||
		functionType.In(0) != paramsValue.Type().Elem() {
		return
	}

	for i := 0; i < paramsValue.Len(); i++ {
		paramValue := paramsValue.Index(i)
		utils.Printfln("Original value (%v)", paramValue)
		utils.Printfln("Processed value (%v)", functionValue.Call([]reflect.Value{paramValue})[0])
	}
}

func Reflection4() {
	utils.Printfln("Reflection4()")

	names := []string{"Alice", "Bob", "John"}
	reflection4_4(strings.ToUpper, names)
}

func reflection5_5(function interface{}, params interface{}) {
	paramsValue := reflect.ValueOf(params)

	if paramsValue.Type().Kind() != reflect.Slice {
		return
	}

	functionValue := reflect.ValueOf(function)
	expectedFunctionType := reflect.FuncOf([]reflect.Type{paramsValue.Type().Elem()}, []reflect.Type{paramsValue.Type().Elem()}, false)

	if functionValue.Type() != expectedFunctionType {
		return
	}

	for i := 0; i < paramsValue.Len(); i++ {
		paramValue := paramsValue.Index(i)
		utils.Printfln("Original value (%v)", paramValue)
		utils.Printfln("Processed value (%v)", functionValue.Call([]reflect.Value{paramValue})[0])
	}
}

func Reflection5() {
	utils.Printfln("Reflection5()")

	names := []string{"Alice", "Bob", "John"}
	reflection5_5(strings.ToUpper, names)
}

func reflection6_6(function interface{}) (newFunction interface{}) {
	functionType := reflect.TypeOf(function)

	if functionType.Kind() != reflect.Func {
		return
	}

	inParamsTypes := make([]reflect.Type, 0, functionType.NumIn())
	outParamsTypes := make([]reflect.Type, 0, functionType.NumOut())

	for i := 0; i < functionType.NumIn(); i++ {
		inParamsTypes = append(inParamsTypes, functionType.In(i))
	}

	for i := 0; i < functionType.NumOut(); i++ {
		outParamsTypes = append(outParamsTypes, functionType.Out(i))
	}

	newFunctionType := reflect.FuncOf(inParamsTypes, outParamsTypes, functionType.IsVariadic())

	newFunctionValue := reflect.MakeFunc(newFunctionType, func(params []reflect.Value) (results []reflect.Value) {
		functionValue := reflect.ValueOf(function)
		results = functionValue.Call(params)
		return
	})

	newFunction = newFunctionValue.Interface()
	return
}

func Reflection6() {
	utils.Printfln("Reflection6()")

	name := "Alice"
	newStringsUpperFunction := reflection6_6(strings.ToUpper).(func(string) string)

	utils.Printfln("Original value (%v)", name)
	utils.Printfln("Processed value (%v)", newStringsUpperFunction(name))

	numberAsString := "111"
	newAtoiFunction := reflection6_6(strconv.Atoi).(func(string) (int, error))

	fmt.Println()
	number, _ := newAtoiFunction(numberAsString)
	utils.Printfln("Original value (%v)", numberAsString)
	utils.Printfln("Processed value (%v)", number)
}

func reflection7_7(arg interface{}) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() != reflect.Struct && argType.Elem().Kind() != reflect.Struct {
		return
	}

	for i := 0; i < argType.NumMethod(); i++ {
		methodType := argType.Method(i)
		utils.Printfln("Method (%v) of type (%v), exported (%v), number of input parameters (%v)",
			methodType.Name,
			methodType.Type,
			methodType.IsExported(),
			methodType.Type.NumIn())
	}
}

func reflection7_7_7(arg interface{}, methodName string) (found bool) {
	argType := reflect.TypeOf(arg)
	_, found = argType.MethodByName(methodName)
	return
}

func Reflection7() {
	utils.Printfln("Reflection7()")
	reflection7_7(data.Milk)

	fmt.Println()
	methodName := "GetName"
	utils.Printfln("Method (%v) found (%v)", methodName, reflection7_7_7(&models.Purchase{}, methodName))
	utils.Printfln("GetName methods has not found since it exists in both Product and Customer properties of Purchase")

	fmt.Println()
	methodName = "GetAmount"
	utils.Printfln("Method (%v) found (%v)", methodName, reflection7_7_7(&models.Purchase{}, methodName))
	utils.Printfln("GetAmount method exists just in Product property of Purchase so it can be found")
}

func reflection8_8(arg interface{}) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() != reflect.Struct && argType.Elem().Kind() != reflect.Struct {
		return
	}

	argValue := reflect.ValueOf(arg)

	for i := 0; i < argType.NumMethod(); i++ {
		methodType := argType.Method(i)
		resultsValue := methodType.Func.Call([]reflect.Value{argValue})
		utils.Printfln("Calling method (%v) gave results", methodType.Name)
		for _, resultValue := range resultsValue {
			utils.Printfln("Result (%v) with value (%v)", resultValue.Type().Name(), resultValue.Interface())
		}
	}
}

func Reflection8() {
	utils.Printfln("Reflection8()")
	reflection8_8(data.Milk)
}

func reflection9_9(arg interface{}) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() != reflect.Struct && argType.Elem().Kind() != reflect.Struct {
		return
	}

	argValue := reflect.ValueOf(arg)

	for i := 0; i < argValue.NumMethod(); i++ {
		methodValue := argValue.Method(i)
		resultsValue := methodValue.Call([]reflect.Value{})
		utils.Printfln("Calling method (%v) gave results", methodValue.Type())
		for _, resultValue := range resultsValue {
			utils.Printfln("Result (%v) with value (%v)", resultValue.Type().Name(), resultValue.Interface())
		}
	}
}

func Reflection9() {
	utils.Printfln("Reflection9()")
	reflection9_9(data.Milk)
}

func reflection10_10(checkInterface interface{}, targets ...interface{}) {
	checkInterfaceType := reflect.TypeOf(checkInterface)

	if checkInterfaceType.Kind() == reflect.Pointer {
		checkInterfaceType = checkInterfaceType.Elem()
	}

	if checkInterfaceType.Kind() != reflect.Interface {
		return
	}

	for _, target := range targets {
		targetType := reflect.TypeOf(target)
		utils.Printfln("Type (%v) implements (%v): (%v)", targetType, checkInterfaceType.Name(), targetType.Implements(checkInterfaceType))
	}
}

func Reflection10() {
	utils.Printfln("Reflection10()")
	reflection10_10((*interfaces.CurrencyItem)(nil), &models.Product{}, &models.Purchase{}, models.Product{}, models.Purchase{})
}

func reflection11_11(arg interface{}) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() == reflect.Pointer {
		argType = argType.Elem()
	}

	argValue := reflect.ValueOf(arg)

	for i := 0; i < argValue.NumField(); i++ {
		fieldValue := argValue.Field(i)
		if fieldValue.Kind() != reflect.Interface {
			continue
		}

		utils.Printfln("Interface (%v) implemented by struct (%v)", fieldValue.Type(), fieldValue.Elem().Type())
	}
}

func Reflection11() {
	utils.Printfln("Reflection11()")

	test := models.WrappedNamedItem{
		NamedItem: data.Milk,
	}

	reflection11_11(test)
}

func reflection12_12(arg interface{}) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() == reflect.Pointer {
		argType = argType.Elem()
	}

	for i := 0; i < argType.NumMethod(); i++ {
		methodType := argType.Method(i)
		utils.Printfln("Method (%v), is exported (%v), pkgPath (%v)", methodType.Name, methodType.IsExported(), methodType.PkgPath)
	}
}

func Reflection12() {
	utils.Printfln("Reflection12()")

	reflection12_12((*interfaces.SampleInterface)(nil))
}

func reflection13_13(arg interface{}) {
	argType := reflect.TypeOf(arg)

	if argType.Kind() != reflect.Chan {
		return
	}

	utils.Printfln("Chan with direction (%v) and type (%v)", argType.ChanDir(), argType.Elem())
}

func Reflection13() {
	utils.Printfln("Reflection13()")
	channel := make(chan<- string, 0)
	reflection13_13(channel)
}

func reflection14_14(channel interface{}, args interface{}) {
	channelType := reflect.TypeOf(channel)
	argsType := reflect.TypeOf(args)

	if channelType.Kind() != reflect.Chan ||
		argsType.Kind() != reflect.Slice ||
		argsType.Elem() != channelType.Elem() {
		return
	}

	channelValue := reflect.ValueOf(channel)
	argsValue := reflect.ValueOf(args)

	for i := 0; i < argsValue.Len(); i++ {
		argValue := argsValue.Index(i)
		channelValue.Send(argValue)
	}

	channelValue.Close()
}

func Reflection14() {
	utils.Printfln("Reflection14()")

	names := []string{"Alice", "Bob"}
	channel := make(chan string, 0)
	go reflection14_14(channel, names)

	for {
		if value, open := <-channel; open {
			utils.Printfln("Received value from channel: %v", value)
		} else {
			utils.Printfln("No more values from channel or channel was closed")
			break
		}
	}
}

func reflection15_15(args interface{}) (channel interface{}) {
	argsType := reflect.TypeOf(args)

	if argsType.Kind() != reflect.Slice {
		return
	}

	channelType := reflect.ChanOf(reflect.BothDir, argsType.Elem())
	channelValue := reflect.MakeChan(channelType, 0)
	argsValue := reflect.ValueOf(args)

	go func() {
		for i := 0; i < argsValue.Len(); i++ {
			argValue := argsValue.Index(i)
			channelValue.Send(argValue)
		}

		channelValue.Close()
	}()

	channel = channelValue.Interface()
	return
}

func Reflection15() {
	utils.Printfln("Reflection15()")

	names := []string{"Alice", "Bob"}
	channel := reflection15_15(names).(chan string)

	for {
		if value, open := <-channel; open {
			utils.Printfln("Received value from channel: %v", value)
		} else {
			utils.Printfln("Channel has no values left or has been closed")
			break
		}
	}
}

func reflection16_16(channels ...interface{}) {
	channelsType := reflect.TypeOf(channels)

	if channelsType.Kind() != reflect.Slice {
		return
	}

	channelsValue := reflect.ValueOf(channels)
	selectCases := make([]reflect.SelectCase, 0, channelsValue.Len())

	for i := 0; i < channelsValue.Len(); i++ {
		channelValue := channelsValue.Index(i).Elem()
		selectCases = append(selectCases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: channelValue,
		})
	}

	for {
		if caseIndex, value, ok := reflect.Select(selectCases); ok {
			utils.Printfln("Received value: %v", value)
		} else {
			if len(selectCases) == 1 {
				utils.Printfln("All channels are closed")
				return
			}
			selectCases = append(selectCases[:caseIndex], selectCases[caseIndex+1:]...)
		}

	}

}

func reflection16_16_16(args interface{}) (channel interface{}) {
	argsType := reflect.TypeOf(args)

	if argsType.Kind() != reflect.Slice {
		return
	}

	channelType := reflect.ChanOf(reflect.BothDir, argsType.Elem())
	channelValue := reflect.MakeChan(channelType, 0)
	argsValue := reflect.ValueOf(args)

	go func() {
		for i := 0; i < argsValue.Len(); i++ {
			argValue := argsValue.Index(i)
			channelValue.Send(argValue)
		}

		channelValue.Close()
	}()

	channel = channelValue.Interface()
	return
}

func Reflection16() {
	utils.Printfln("Reflection16()")

	names := []string{"Alice", "Bob"}
	numbers := []int{111, 222, 333}

	channel1 := reflection16_16_16(names).(chan string)
	channel2 := reflection16_16_16(numbers).(chan int)

	reflection16_16(channel1, channel2)
}
